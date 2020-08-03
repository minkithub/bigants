"""
usage:
    docker-compose exec anthill python3 manage.py shell
    exec(open('utils/uploader.py').read())
"""
import glob
import os
import pandas as pd

from predictor.models import Stock, Price

BASE = 'initial-data'
MARKETS = 'kospi kosdaq konex'.split()

def stock_upload():
    bulk = []
    for market in MARKETS:
        df = pd.read_csv(f'{BASE}/{market}.csv')

        bulk += [Stock(code=encode_stock(row[1], market), name_ko=row[2]) for _, row in df.iterrows()]

    Stock.objects.bulk_create(bulk)

def price_upload():
    counter = 1

    for market in MARKETS:
        for path in glob.glob(f'{BASE}/{market}/*.csv'):
            code = os.path.splitext(os.path.basename(path))[0]
            df = pd.read_csv(path)

            bulk = []
            for i, row in df.iterrows():
                dt, cv, sv, hv, lv, vl = row
                bulk.append(Price(
                    stock_code_id=encode_stock(code, market),
                    record_date=dt,
                    open_value=sv,
                    high_value=hv,
                    low_value=lv,
                    close_value=cv,
                    volume=vl
                ))

            Price.objects.bulk_create(bulk)

            print(counter, code, 'is done!')

            counter += 1

def main():
    print('execute in django shell!')

def encode_stock(code, market):
    code = int(code)
    if market == 'kospi':
        return f'{code:06}.KS'
    elif market == 'kosdaq':
        return f'{code:06}.KQ'
    elif market == 'konex':
        return f'{code:06}.KN'
    else:
        return str(code)

if __name__ == '__main__':
    main()


# uncomment necessary commands

# Price.objects.all().delete()
# Stock.objects.all().delete()
# stock_upload()
# price_upload()