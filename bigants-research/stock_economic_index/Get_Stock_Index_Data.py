'''
Get Various Stock Index
2020.03.11
by Minki
'''

'''
해당 코드에서 사용된 지표 정리
1. USD_KRW : KRWUSD=X
2. USD_JPY : JPYUSD=X
3. USD_CNY : CNYUSD=X
4. USD_SGD : SGDUSD=X
5. USD_HKD : HKDUSD=X
6. AUD_USD : AUDUSD=X
7. GBP_USD : GBPUSD=X
8. EUR_USD : EURUSD=X
9. USD_CAD : CADUSD=X
10. USD : USD
11. GOLD : GC=F
12. OIL : OIL
13. COPPER : HG=F
14. KOR_TBOND
15. US_TBOND : ^IRX
16. US_HY : USHY
17. KOSPI : ^KOSPI
18. KOSDAQ : ^KOSDAQ
19. NIKKEI225 : ^N225
20. SHANGHAI : ^SSEC
21. SINGAPORE : ^STI
22. HANGSENG : ^HSI
23. ASX100 : ^HSI
24. DAX : ^GDAXI
25. TSX60 : XIU.TO
26. SPX500 : ^GSPC
27. NASDAQ : ^IXIC
28. FTSE100 : ^FTSE
29. VIX : ^VIX
30. UST5Y : ^FVX
31. UST10Y : ^TNX
32. UST30Y : ^TYX
33. UST_HQMYC_0_5
34. UST_HQMYC_10_0
35. HOUSE
36. PMI : PMI
37. INFLATION : CPI
38. LABOR
39. Dow Jones : DJI
40. 미국 최대의 ETF Fund : SPY
41. 미국 두번쨰 ETF Fund : IVV
42. Russell 2000 : ^RUT
43. Crude Oil : CL=F
44. Silver : SI=F
45. Bitcoin USD : BTC-USD
46. Bitcoin KRW : BTC-KRW
47. United States Brent Oil Fund : BNO
48. 
'''


from pandas_datareader import data as pdr
import pandas as pd
from datetime import datetime
import yfinance as yf

# Set index list
index_list = ['SPY', 'IVV', '^GSPC', 'DJI', '^IXIC', '^RUT', 
'CL=F', 'GC=F', 'SI=F', 'EURUSD=X', '^TNX', '^VIX', 'GBPUSD=X', 'JPYUSD=X', 'BTC-USD',
'^CMC200', '^FTSE', '^N225', '^KOSPI', '^KOSDAQ', 'AUDUSD=X', 'BNO', 'VTI', 'SHY',
'SOXX', 'OIL', 'CNYUSD=X']

# download dataframe using pandas_datareader
yf.pdr_override()
start_date = "2017-01-01"
end_date = datetime.today().strftime('%Y-%m-%d')
counter = 1


for i in index_list:
    print('i:', i)
    data = pdr.get_data_yahoo(i, start=start_date, end=end_date)
    data['index'] = i
    data['Date'] = data.index
    data.reset_index(drop = True, inplace = True)
    data = data[['Date', 'Open', 'High', 'Low', 'Close', 'Adj Close', 'Volume', 'index']]
    print(data.head(1))
    print(data.tail(1))
    data.to_excel(f'/Users/minki/pythonworkspace/bigants/stock_direction/dataset/{i}.xlsx')
    print(f'{counter}/{len(index_list)}', i, 'is done!')

    counter += 1







