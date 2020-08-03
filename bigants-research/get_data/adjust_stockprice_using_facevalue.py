# import numpy as np
import pandas as pd
# import datetime as dt
# import warnings
# warnings.filterwarnings("ignore")
# from sklearn.metrics import mean_squared_error
# import os
# from fbprophet import Prophet

# 1. 프로펫 예측값 나옴
# 2. 액면가에 맞춰서 비율 조정
# 3. 원래 close와 합침
# 4. MA 계산 후 cross 파악 후 매도/매수 시점 파악

# 매수평단 : 같은 주식을 다른 가격에 분할해서 샀을 때 전체 평균 가격
# 매수평단 관점에서는 주가가 1원단위까지 나올 수 있다.
# 그러나 매수평단은 종가랑은 또 다른 개념이므로 최종 종가는 100원 or 1원으로 하는게 가장 합리적으로 보인다.

facevalue = pd.read_excel('/Users/minki/pythonworkspace/bigants/dataset/facevalue.xlsx')
data = pd.read_excel('/Users/minki/pythonworkspace/bigants/data_semi_result/cal_error2.xlsx')

facevalue['code'] = facevalue['code'].apply(lambda x: "{:0>6}".format(x))
data['code'] = data['code'].apply(lambda x: "{:0>6}".format(x))

data = data[['date', 'close', 'code']]
print(data.tail())
print(data.iloc[-1]['close'])
