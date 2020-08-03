import datetime
import pandas as pd
import requests
import time

from lxml import html

"""naver_stock.py

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
    
## 크롤링에서 페이지가 뒤로 갈수록 날짜가 더 과거임.
## 그래서 end는 가장 마지막 페이지, 즉 가장 먼 날짜를 의미함.

def get_prices(code, end, delay):
    page = 1
    total_df = pd.DataFrame()
    while True:
        # params는 requests 함수의 기본 변수.
        # 네이버 주식 url = https://finance.naver.com/item/frgn.nhn?code=155660&page=10
        # 위에서 url이 code와 page로 바뀌는데 그걸 requests 함수에서 지정해주는 것.
        response = requests.get(URL, params={ 'code': code, 'page': page }, timeout=5)

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
    # list형태의 table을 만들어줌. 
    table = []
    for r in range(start, end):
        row = []
        for c in range(1, 8):
            if c == 3:
                continue
            spans = tree.xpath(f'/html/body/table[1]/tr[{r}]/td[{c}]/span')
            if spans:
                row.append(spans[0].text.strip())

        if row:
            table.append(row)

    for row in table:
        row[0] = datetime.datetime.strptime(row[0], '%Y.%m.%d')
        row[1:] = map(lambda x: int(x.replace(',', '')), row[1:])

    return table

if __name__ == '__main__':
    main()