import datetime
import json
import math
import pandas as pd
import threading

from .models import Stock, Price
from .utils import naver_stock
from django.http import HttpResponse, JsonResponse
from django.shortcuts import render, get_list_or_404
from django.views import generic
from fbprophet import Prophet

with open('config.json') as f:
    config = json.load(f)

class IndexView(generic.ListView):
    model = Stock
    template_name = 'predictor/index.html'

def get_stocks(request):
    if request.method != 'GET':
        return JsonResponse({'error': 'Invalid request method'})

    if 'q' in request.GET:
        stocks = Stock.objects.filter(name_ko__startswith=request.GET['q'])
    else:
        stocks = Stock.objects.all()

    content = {}
    content['data'] = [{'code': stock.code, 'name_ko': stock.name_ko}
                       for stock in stocks if not is_removed(stock.code)]

    return JsonResponse(content, json_dumps_params={'ensure_ascii': False})

def is_removed(stock_code):
    code = stock_code.split('.')[0]

    if code not in config:
        return False

    return 'remove' in config[code]

def get_history(request):
    if request.method != 'POST':
        return JsonResponse({'error': 'Invalid request method'})

    params = json.loads(request.body)
    if 'code' not in params:
        return JsonResponse({'error': 'Not included stock code'})

    stock = get_stock_or_none(params['code'])
    if not stock:
        return JsonResponse({'error': 'Invalid stock code'})

    start, end = get_dates_or_default(params)
    if start > end:
        return JsonResponse({'error': 'Invalid dates'})

    df = get_prices(stock, start, end)

    date_prices = {price.record_date: [price.high_value, price.low_value, price.close_value]
                   for _, price in df.iterrows()}

    content = {}
    content['data'] = [date_prices[date] if date in date_prices.keys() else None
                       for date in daterange(start, end)]

    return JsonResponse(content)

def get_stock_or_none(stock_code):
    code = stock_code.split('.')[0]
    if code in config and 'remove' in config[code]:
        return None

    return Stock.objects.filter(pk=stock_code).first()


def get_prices(stock, start, end):
    code = stock.code.split('.')[0]

    prices = Price.objects.filter(
        stock_code=stock
    ).filter(
        record_date__gte=start
    ).filter(
        record_date__lte=end
    ).exclude(
        high_value=0
    )

    if code in config and 'cut' in config[code]:
        prices = prices.filter(record_date__gte=config[code]['cut'])

    df = pd.DataFrame(list(prices.values()))

    if code in config and 'adjust' in config[code]:
        for cond in config[code]['adjust']:
            df[df['record_date'] < parse_date(cond['date']).date()][[
                'open_value', 'low_value', 'high_value', 'close_value'
            ]] /= cond['coef']

    return df

def get_predict(request):
    if request.method != 'POST':
        return JsonResponse({'error': 'Invalid request method'})

    params = json.loads(request.body)
    if 'code' not in params:
        return JsonResponse({'error': 'Not included stock code'})

    stock = get_stock_or_none(params['code'])
    if not stock:
        return JsonResponse({'error': 'Invalid stock code'})

    start, end = get_dates_or_default(params)
    if start > end:
        return JsonResponse({'error': 'Invalid dates'})

    type_ = params.get('type', 'CLOSE').lower()
    holidays = pd.DataFrame([{'holiday': f'holiday{i}', 'ds': holiday} for i, holiday in enumerate(params.get('holidays', []))])
    future = pd.date_range(start=datetime.datetime.now(), periods=14).to_frame(index=False, name='ds')

    df = get_prices(stock, start, end)[['record_date', f'{type_}_value']]
    df.columns = ['ds', 'y']

    forecast, mae, mape = predict(df, holidays, future)

    content = {}
    content['data'] = [row['yhat'] for _, row in forecast.iterrows()]
    content['mae'] = mae
    content['mape'] = mape

    return JsonResponse(content)

def get_dates_or_default(params):
    now = datetime.datetime.now()

    start = parse_date(params['start']) if 'start' in params else now - datetime.timedelta(365)
    end = parse_date(params['end']) if 'end' in params else now

    return start, end

def predict(df, holidays, future):
    model = Prophet(holidays=holidays if not holidays.empty else None)

    train_df, test_df = split(df, 0.2)

    model.fit(train_df)

    y = test_df['y'].reset_index(drop=True)
    yhat = model.predict(test_df)['yhat'].reset_index(drop=True)

    mae = (y - yhat).abs().mean()
    mape = ((y - yhat) / y).abs().mean() * 100

    return model.predict(future), mae, mape

def split(df, testset_ratio):
    idx = int(len(df) * (1 - testset_ratio))

    train_df = df.iloc[:idx]
    test_df = df.iloc[idx:]

    return train_df, test_df

# 'YYYY-MM-DD' to datetime
def parse_date(string):
    return datetime.datetime.strptime(string, '%Y-%m-%d')

def daterange(start, end):
    for i in range((end - start).days + 1):
        yield (start + datetime.timedelta(i)).date()

def result(request):
    df = pd.DataFrame(list(Price.objects.filter(stock_code=request.POST['stock_code']).values()))[['record_date', 'close_value']]
    df.columns = ['ds', 'y']

    holidays = pd.DataFrame(columns=['holiday', 'ds'])
    i = 0
    for date, name in zip(request.POST.getlist('dates[]'), request.POST.getlist('names[]')):
        if not date or not name:
            continue
        holidays.loc[i] = [name, pd.to_datetime(date)]
        i += 1

    model = Prophet(holidays=holidays)
    model.fit(df)
    future = model.make_future_dataframe(periods=365)
    forecast = model.predict(future)

    rmse = math.sqrt(((forecast.iloc[:-365]['yhat'] - df['y']) ** 2).sum() / len(df))

    start = datetime.datetime.now()
    end = start + datetime.timedelta(days=365)
    mask = (future['ds'] >= start) & (future['ds'] <= end)

    results = forecast[['ds', 'yhat', 'yhat_lower', 'yhat_upper']][mask].values
    results[:, 0] = [t.date() for t in results[:, 0]]
    for i in range(1, 4):
        results[:, i] = [f'{n:.4f}' for n in results[:, i]]

    return render(request, 'predictor/result.html', {
        'results': results,
        'rmse': rmse
   })

def update_prices(request):
    t = threading.Thread(target=update_prices_thread)
    t.start()

    return JsonResponse({'state': 'success?'})

def update_prices_thread():
    print('Thread starts!')
    now = datetime.datetime.now()
    end_time = now + datetime.timedelta(minutes=55)

    updated_stocks = [price.stock_code.code for price in Price.objects.filter(record_date=now.date())]

    for i, stock in enumerate(Stock.objects.all()):
        if stock.code in updated_stocks:
            continue

        latest = Price.objects.filter(stock_code=stock).latest('record_date').record_date
        df = naver_stock.get_prices(stock.code.split('.')[0], latest.strftime('%Y-%m-%d'), 0.25)
        df = df[df.index > pd.to_datetime(latest)]

        bulk = []
        for i, row in df.iterrows():
            cv, sv, hv, lv, vl = row
            bulk.append(Price(
                stock_code=stock,
                record_date=i,
                open_value=sv,
                high_value=hv,
                low_value=lv,
                close_value=cv,
                volume=vl
            ))

        Price.objects.bulk_create(bulk)

        print(i, stock.code, 'update!')

        if datetime.datetime.now() >= end_time:
            break

    print('Thread ends!')
