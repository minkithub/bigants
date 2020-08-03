'''
get stock information
latest Update : 2020-04-28
by Minki
reference : https://shinminyong.tistory.com/15

<variable.>
분기 : quarter[QT]
매출액 : sales[SAL]
영업이익 : operating_revenue[OPR]
영업이익(발표기준) : operating_revenue_announce[OPRA]
세전계속사업이익 : gain_on_pretax_business[GPB]
당기순이익 : net_income[NI]
당기순이익(지배) : net_income_controlling[NIC]
당기순이익(비지배) : net_income_noncontrolling[NIN]
자산총계 : total_assets[TA]
부채총계 : total_liabilites[TL]
자본총계 : total_equities[TE]
자본총계(지배) : total_equities_controlling[TEC]
자본총계(비지배) : total_equities_noncontrolling[TEN]
자본금 : cpaital_stock[CS]
영업활동현금흐름 : operating_activities[OA]
투자활동현금흐름 : investing_activities[IA]
재무활동현금흐름 : financing_activities[FA]
CAPEX : CAPEX
FCF : FCF
이자발생부채 : interest_accrued_liabilites[IAL]
영업이익률 : operating_earning_rate[OER]
순이익률 : rate_of_return[ROR]
ROE(%) : ROE
ROA(%) : ROA
부채비율 : debt_ratio[DR]
자본유보율 : capital_retention_rate[CRR]
EPS(원) : EPS
PER(배) : PER
BPS(원) : BPS
PBR(배) : PBR
현금DPS(원) : money_DPS[DPS]
현금배당수익률 : dividend_yeild_ratio[DYR]
현금배당성향(%) : dividend_propensity[DP]
발행주식수(보통주) : number_of_stock[NOS]
'''

import requests
from bs4 import BeautifulSoup
from datetime import datetime
import pandas as pd
import time
import json
import re     
import datetime as dt
from selenium import webdriver
import os
import time
# from selenium.webdriver.chrome.options import Options
# from selenium.webdriver.common.keys import Keys
# import urllib.request
# from selenium.webdriver import Chrome

start = time.time()

def get_html(url):
   _html = ""
   resp = requests.get(url)
   if resp.status_code == 200:
      _html = resp.text
   return _html

URL = 'https://finance.naver.com/item/coinfo.nhn?code=005380'

browser = webdriver.PhantomJS('/Users/minki/crawling/phantomjs-2.1.1-macosx/bin/phantomjs')
browser.get(URL)
browser.switch_to_frame(browser.find_element_by_id('coinfo_cp'))

# 연간 재무제표
browser.find_elements_by_xpath('//*[@class="schtab"][1]/tbody/tr/td[3]')[0].click()

stock_html = browser.page_source
soup_html = BeautifulSoup(stock_html,'html.parser')

company_name = soup_html.find('head').find('title').text
company_name = company_name.split('-')[-1]

stock_info = soup_html.find('table',{'class':'gHead01 all-width','summary':'주요재무정보를 제공합니다.'})

# Make DataFrame
stock_df_year = pd.DataFrame(columns = ['QT', 'SAL', 'OPR', 'OPRA', 'GPB', 'NI', 'NCI', 'NIN', 'TA', 'TL', 'TE', 'TEC',
'TEN', 'CS', 'OA', 'IA', 'FA', 'CAPEX', 'FCF', 'IAL', 'OER', 'ROR', 'ROE', 'ROA', 'DR', 'CRR', 'EPS',
'PER', 'BPS', 'PBR', 'DPS', 'DYR', 'DP', 'NOS'])

# 분기 데이터 추출
date_thead = stock_info.find('thead')
date_tr = date_thead.find_all('tr')[1]
date_th = date_tr.find_all('th')

date = []
for i in date_th:
   val = i.text.strip()[:7]
   date.append(val)

stock_df_year['QT'] = date
stock_df_year['QT'] = stock_df_year['QT'] + '/31'

# 재무제표에서 데이터 추출
stock_tbody = stock_info.find('tbody')
stock_data = stock_tbody.find_all('tr')

# Get Stock Data
col_data = []
for i in range(0, len(stock_data)):
   total_data = stock_data[i].find_all('td')
   sub_val = []
   for j in range(len(total_data)):
      if total_data[j].text == '':
         sub_val.append(0)
      else:
         val = total_data[j].text.strip()
         sub_val.append(val)
   col_data.append(sub_val)
   stock_df_year.iloc[:, i+1] = col_data[i]

###############################################################################################################################################
###############################################################################################################################################

# 분기 재무제표
browser.find_elements_by_xpath('//*[@class="schtab"][1]/tbody/tr/td[4]')[0].click()

stock_html = browser.page_source
soup_html = BeautifulSoup(stock_html,'html.parser')

company_name = soup_html.find('head').find('title').text
company_name = company_name.split('-')[-1]

stock_info = soup_html.find('table',{'class':'gHead01 all-width','summary':'주요재무정보를 제공합니다.'})

# Make DataFrame
stock_df_quarter = pd.DataFrame(columns = ['QT', 'SAL', 'OPR', 'OPRA', 'GPB', 'NI', 'NCI', 'NIN', 'TA', 'TL', 'TE', 'TEC',
'TEN', 'CS', 'OA', 'IA', 'FA', 'CAPEX', 'FCF', 'IAL', 'OER', 'ROR', 'ROE', 'ROA', 'DR', 'CRR', 'EPS',
'PER', 'BPS', 'PBR', 'DPS', 'DYR', 'DP', 'NOS'])

# 분기 데이터 추출
date_thead = stock_info.find('thead')
date_tr = date_thead.find_all('tr')[1]
date_th = date_tr.find_all('th')

date = []
for i in date_th:
   val = i.text.strip()[:7]
   date.append(val)

stock_df_quarter['QT'] = date
stock_df_quarter['QT'] = stock_df_quarter['QT'] + '/15'

# 재무제표에서 데이터 추출
stock_tbody = stock_info.find('tbody')
stock_data = stock_tbody.find_all('tr')

# Get Stock Data
col_data = []
for i in range(0, len(stock_data)):
   total_data = stock_data[i].find_all('td')
   sub_val = []
   for j in range(len(total_data)):
      if total_data[j].text == '':
         sub_val.append(0)
      else:
         val = total_data[j].text.strip()
         sub_val.append(val)
   col_data.append(sub_val)
   stock_df_quarter.iloc[:, i+1] = col_data[i]

total_stock_df = pd.concat([stock_df_year, stock_df_quarter], axis = 0)
total_stock_df['QT'] = pd.to_datetime(total_stock_df['QT'])
total_stock_df = total_stock_df.sort_values(by='QT')

print(total_stock_df)
print("simulation time :", time.time() - start)  # 현재시각 - 시작시간 = 실행 시간