"""
usage:
    $ docker-compose up &
    $ docker-compose exec anthill python3 manage.py shell
    >>> from predictor.scripts.db import *
    >>> delete_duplicate_prices()
"""
from predictor.models import Price, Stock

def delete_duplicate_prices(date):
    pids = set()

    queryset = Price.objects.filter(record_date__gte=date)
    count = queryset.count()
    for i, price in enumerate(queryset):
        pid = (price.stock_code.code, price.record_date)

        if pid in pids:
            print(i, price, 'is removed!')
            price.delete()
        else:
            pids.add(pid)

        if i % 1000 == 0:
            print(f'{i}/{count}', 'passed...')

def view_recent_prices(date):
    prices = [price for price in Price.objects.filter(record_date__gte=date)]

    len_stocks = {}
    for stock in Stock.objects.all():
        selected_prices = [price for price in prices if price.stock_code == stock]

        if len(selected_prices) in len_stocks:
            len_stocks[len(selected_prices)][stock.code] = selected_prices
        else:
            len_stocks[len(selected_prices)] = {}
            len_stocks[len(selected_prices)][stock.code] = selected_prices

    for key in len_stocks:
        print(key, '-', len(len_stocks[key]))

def main():
    pass

if __name__ == '__main__':
    main()
