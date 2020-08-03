"""
주식의 재무재표 정보를 scrapping하는 코드입니다.
Hyeonjin Kim
2020.02.11

<variable.>
분기 : quarter
매출액 : sales
영업이익 : operating_revenue
영업이익(발표기준) : operating_revenue_announce
세전계속사업이익 : gain_on_pretax_business
당기순이익 : net_income
당기순이익(지배) : net_income_controlling
당기순이익(비지배) : net_income_noncontrolling
자산총계 : total_assets
부채총계 : total_liabilites
자본총계 : total_equities
자본총계(지배) : total_equities_controlling
자본총계(비지배) : total_equities_noncontrolling
자본금 : cpaital_stock
영업활동현금흐름 : operating_activities
투자활동현금흐름 : investing_activities
재무활동현금흐름 : financing_activities
CAPEX : CAPEX
FCF : FCF
이자발생부채 : interest_accrued_liabilites
영업이익률 : operating_earning_rate
순이익률 : rate_of_return
ROE(%) : ROE
ROA(%) : ROA
부채비율 : debt_ratio
자본유보율 : capital_retention_rate
EPS(원) : EPS
PER(배) : PER
BPS(원) : BPS
PBR(배) : PBR
현금DPS(원) : money_DPS
현금배당수익률 : dividend_yeild_ratio
현금배당성향(%) : dividend_propensity
발행주식수(보통주) : number_of_stock
"""
import datetime
import pandas as pd
import requests
import time
import re
from lxml import html
from bs4 import BeautifulSoup


def get_html(url):
   _html = ""
   resp = requests.get(url)
   if resp.status_code == 200:
      _html = resp.text
      soup = BeautifulSoup(_html, 'html.parser')
   return _html, soup

code_list = open('/Users/minki/pythonworkspace/bigants/dataset/stock_codes.txt', mode = 'r')

# URL = 'https://navercomp.wisereport.co.kr/v2/company/c1010001.aspx?cmp_cd=006840'
URL = 'https://finance.naver.com/item/coinfo.nhn?code=005380'
html, soup = get_html(URL)

stock_values = soup.find_all('span', {'class': 'cBk'})
print(stock_values)
stock_last_values = soup.find_all('span', {'class': 'cUp'})
print(stock_last_values)
titles = soup.find_all('th', {'class': 'bg txt title'})
print(titles)
test = soup.find('div', id = 'aFVlanREZS').find_all("tr")
print(test)

inner_content = requests.get(soup.find("iframe")["src"])
print(inner_content)

# 행 인덱스가 될 분기 날짜를 뽑는 코드
# _list = []
# for i in range(1, 9):
#     if i != 8:
#         quarter_id = soup.find_all('th', {'class': f'r03c0{i} bg'})
#         print(quarter_id)
#         _value = quarter_id[0].text.strip()
#         _list.append(_value)
#     else:
#         quarter_id = soup.find_all('th', {'class': f'r03c0{i} bg endLine'})
#         _value = quarter_id[0].text.strip()
#         _list.append(_value)

# cBk 변수를 뽑는 코드
# _cBk = []
# for stock_value in stock_values:
#     val = stock_value.text
#     _cBk.append(val)

# # cUp 변수를 뽑는 코드
# _cUp = []
# for stock_last_value in stock_last_values:
#     val = stock_last_value.text
#     _cUp.append(val)

# # title 변수를 뽑는 코드(열이름으로 사용)
# col_name = []
# for title in titles:
#     val = titles.text
#     col_name.append(val)

# # print(_list)
# print('==============================================')
# print(_cBk)
# print('==============================================')
# print(_cUp)
# print('==============================================')
# print(col_name)