"""prophet_example.py

$ docker-compose up
$ docker-compose exec anthill python3 manage.py
>>> exec.run(open('prophet_example.py').read())


Hyeonjin Kim
2020.01.23
"""
import numpy as np
import math
import json
import requests
import pandas as pd
from fbprophet import Prophet
from predictor.models import Stock, Price


def main():
    for stock in Stock.objects.all():
        response = requests.post('https://anthill-ljnq6oi2dq-an.a.run.app/api/v1/history', json={ "code": stock.code, "start": "2017-01-01", "end": "2020-02-15" }, timeout=10)
        data = response.json()

        data = pd.Series([float(i[2]) if i else -1 for i in data['data']])

        daterange = pd.date_range(start='2017-01-01', end='2020-02-14').to_series()

        df = pd.DataFrame({ 'ds': daterange.values, 'y': data })
        df = df[df['y'] != -1]

        trainset = df[:-100]
        testset = df[-1:]

        model = Prophet()
        model.fit(trainset)
        future = testset[['ds']]
        forecast = model.predict(future)['yhat']

        ave = sum(testset['y'].values) / len(testset)
        # rmse = math.sqrt(((forecast.values - testset['y'].values) ** 2).sum() / len(testset))
        # mae = abs(forecast.values - testset['y'].values).sum() / len(testset['y'].values)
        mape = mean_absolute_percentage_error(testset['y'].values, forecast.values)
        # mae = forecast.values - test

        print('name:', stock.name_ko, 'mape:', 100 -mape)

    # print(forecast[forecast['ds'].apply(lambda x: x in dates or x in dates2)])
    # print(forecast.tail())

def mean_absolute_percentage_error(y_true, y_pred):
    y_true, y_pred = np.array(y_true), np.array(y_pred)
    return np.mean(np.abs((y_true - y_pred) / y_true)) * 100


    # df = pd.read_csv('../sample.csv')
    # df = pd.read_csv('../000020.KS.csv')
    # df = df[['Date', 'Adj Close']]
    # df.columns = ['ds', 'y']

    # dates = pd.to_datetime([
    #     '2010-05-07',
    #     '2011-12-21',
    #     '2012-10-05',
    #     '2013-03-04',
    # ])

    # dates2 = pd.to_datetime([
    #     '2016-01-05',
    #     '2016-01-20',
    #     '2020-02-28',
    #     '2013-11-06'
    # ])

    # foo = pd.DataFrame({
    #     'holiday': 'foo',
    #     'ds': dates
    # })

    # bar = pd.DataFrame({
    #     'holiday': 'bar',
    #     'ds': dates2
    # })
    # holidays = pd.concat((foo, bar))

    # model = Prophet(holidays=holidays)
    # model.fit(df)
    # future = model.make_future_dataframe(periods=365)
    # forecast = model.predict(future)[['ds', 'yhat', 'yhat_lower', 'yhat_upper', 'foo', 'bar']]

    # print(forecast[forecast['ds'].apply(lambda x: x in dates or x in dates2)])
    # print(forecast.tail())


# if __name__ == '__main__':
    # main()

main()