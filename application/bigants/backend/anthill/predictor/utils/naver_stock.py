"""naver_stock.py

NOT USABLE!!!

Hyeonjin Kim
2020.02.11
"""
import datetime
import pandas as pd
import requests
import time

from lxml import html


URL = 'https://finance.naver.com/item/sise_day.nhn'
MARKETS = 'kospi kosdaq konex'.split()

def main():
    f = open('logs.txt', 'w')
    for market in MARKETS:
        stocks = pd.read_csv(f'{market}.csv', dtype='object')

        for i, row in stocks.iterrows():
            name = row[1]
            try:
                prices = get_prices(name, '2017-01-01', 0.15)
                prices.to_csv(f'{market}/{name}.csv')
                print(f'{market}: {i}/{len(stocks)} is done!')
            except Exception as e:
                f.write(f'{datetime.datetime.now()}: {name} ({e})')
                print(f'{market}: {i}/{len(stocks)} fails! Xp')

    f.close()

def get_prices(code, end, delay):
    page = 1
    total_df = pd.DataFrame()
    while True:
        response = requests.get(URL, params={ 'code': code, 'page': page }, timeout=10)
        df = parse_page(response.text)

        if not total_df.index.empty and df.index[-1] == total_df.index[-1]:
            break

        total_df = total_df.append(df)

        if len(df) != 10 or df.iloc[-1].name <= datetime.datetime.strptime(end, '%Y-%m-%d'):
            break

        time.sleep(delay)
        page += 1

    return total_df

def parse_page(html_str):
    tree = html.fromstring(html_str)

    table = parse_tree_by_row_range(tree, 3, 8) + (parse_tree_by_row_range(tree, 11, 16))

    df = pd.DataFrame(table, columns=['date', 'close', 'start', 'high', 'low', 'volume'])
    df = df.set_index('date')

    return df

def parse_tree_by_row_range(tree, start, end):
    table = []
    for r in range(start, end):
        row = []
        for c in range(1, 8):
            if c == 3:
                continue
            spans = tree.xpath(f'/html/body/table[1]/tr[{r}]/td[{c}]/span')
            if spans and spans[0].text:
                row.append(spans[0].text.strip())

        if row:
            table.append(row)

    for row in table:
        row[0] = datetime.datetime.strptime(row[0], '%Y.%m.%d')
        row[1:] = map(lambda x: int(x.replace(',', '')), row[1:])

    return table

if __name__ == '__main__':
    main()
