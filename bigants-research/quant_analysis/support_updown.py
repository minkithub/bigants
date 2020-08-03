'''
주식 보조지표의 Signal(Sell, Buy, Wait)에 따른 주가 상승과 하락을 예측하는 코드입니다.
total_data : 보조지표의 Value와 Timing 모든것이 들어가있는 데이터.
support_result = 보조지표의 시뮬레이션 결과 데이터
latest Update : 2020-04-23
by Minki
'''

import numpy as np
import pandas as pd
import datetime as dt
import warnings
warnings.filterwarnings("ignore")
import os
import sys
sys.path.insert(0, '/Users/minki/pythonworkspace/bigants/bigants-research/quant_analysis/calculate_stock_support')
from calculate_stock_support import stock_value
from calculate_stock_support import stock_trade_function
import combine_stock_data
import stock_simulation
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import classification_report,confusion_matrix
from sklearn.tree import DecisionTreeClassifier
import time
start = time.time()  # 시작 시간 저장

total_data, support_result, total_timing_binary, total_numeric_binary, total_binary_data, total_half_binary_data = combine_stock_data.main()

# 기본 classification_LogisticRegression

# [1. Timing data만을 사용하여 예측 진행]

## 1-1. Make train and test data
## random_state는 seed를 일정하게 유지해주는 역할을 한다.
X_train_timing, X_test_timing, y_train_timing, y_test_timing = train_test_split(total_timing_binary.drop(['binary_close'],axis=1), 
total_timing_binary['binary_close'], test_size=0.15, random_state=101)

## 1-2. LogisticRegression
## train 과정
logmodel = LogisticRegression()
logmodel.fit(X_train_timing,y_train_timing)

## 학습 과정
predictions = logmodel.predict(X_test_timing)

## 결과 검증과정
print('Timing data만을 사용하여 얻은 logistic regression 결과데이터입니다.')
print(confusion_matrix(y_test_timing,predictions))
print(classification_report(y_test_timing,predictions))
print('=================================================')

# [2. Numeric data만을 사용하여 예측 진행]

## 2-1. Make train and test data
## random_state는 seed를 일정하게 유지해주는 역할을 한다.
X_train_numeric, X_test_numeric, y_train_numeric, y_test_numeric = train_test_split(total_numeric_binary.drop(['binary_close'],axis=1), 
total_numeric_binary['binary_close'], test_size=0.15, random_state=101)

## 2-2. LogisticRegression
## train 과정
logmodel = LogisticRegression()
logmodel.fit(X_train_numeric,y_train_numeric)

## 학습 과정
predictions = logmodel.predict(X_test_numeric)

## 결과 검증과정
print('Numeric data만을 사용하여 얻은 logistic regression 결과데이터입니다.')
print(confusion_matrix(y_test_numeric,predictions))
print(classification_report(y_test_numeric,predictions))
print('=================================================')

# [3. Numeric, Timing 모두 사용하여 예측 진행]

## 2-1. Make train and test data
## random_state는 seed를 일정하게 유지해주는 역할을 한다.
X_train_total, X_test_total, y_train_total, y_test_total = train_test_split(total_binary_data.drop(['binary_close'],axis=1), 
total_binary_data['binary_close'], test_size=0.15, random_state=101)

## 2-2. LogisticRegression
## train 과정
logmodel = LogisticRegression()
logmodel.fit(X_train_total,y_train_total)

## 학습 과정
predictions = logmodel.predict(X_test_total)

## 결과 검증과정
print('Numeric data와 Timing data 모두 사용하여 얻은 logistic regression 결과데이터입니다.')
print(confusion_matrix(y_test_total,predictions))
print(classification_report(y_test_total,predictions))
print('=================================================')

# [4. Numeric(original), Timing 모두 사용하여 예측 진행]

## 2-1. Make train and test data
## random_state는 seed를 일정하게 유지해주는 역할을 한다.
X_train_half, X_test_half, y_train_half, y_test_half = train_test_split(total_half_binary_data.drop(['close'],axis=1), 
total_binary_data['binary_close'], test_size=0.15, random_state=101)
## 2-2. LogisticRegression
## train 과정
logmodel = LogisticRegression()
logmodel.fit(X_train_half,y_train_half)

## 학습 과정
predictions = logmodel.predict(X_test_half)

## 결과 검증과정
print('Numeric data와 Timing data 모두 사용하여 얻은 logistic regression 결과데이터입니다.')
print(confusion_matrix(y_test_half,predictions))
print(classification_report(y_test_half,predictions))
print('=================================================')

# 기본 classification_DecisionTreeClassifier
# 데이터셋은 위해서 만든 데이터셋을 그대로 사용하였음

# [1. Timing 데이터만을 사용하여 예측 진행]

dt_model=DecisionTreeClassifier()
dt_model.fit(X_train_timing,y_train_timing)
dt_pred = dt_model.predict(X_test_timing)

print('Timing data만을 사용하여 얻은 decistion Tree 결과데이터입니다.')
print(confusion_matrix(y_test_timing,dt_pred))
print(classification_report(y_test_timing,dt_pred))
print('=================================================')

# [2. numeric 데이터만을 사용하여 예측 진행]

dt_model=DecisionTreeClassifier()
dt_model.fit(X_train_numeric,y_train_numeric)
dt_pred = dt_model.predict(X_test_numeric)

print('numeric data만을 사용하여 얻은 decistion Tree 결과데이터입니다.')
print(confusion_matrix(y_test_numeric,dt_pred))
print(classification_report(y_test_numeric,dt_pred))
print('=================================================')

# [3. Total 데이터를 사용하여 예측 진행]

dt_model=DecisionTreeClassifier()
dt_model.fit(X_train_total,y_train_total)
dt_pred = dt_model.predict(X_test_total)

print('numeric data만을 사용하여 얻은 decistion Tree 결과데이터입니다.')
print(confusion_matrix(y_test_total,dt_pred))
print(classification_report(y_test_total,dt_pred))
print('=================================================')

# [4. Total hal 데이터를 사용하여 예측 진행]

dt_model=DecisionTreeClassifier()
dt_model.fit(X_train_half,y_train_half)
dt_pred = dt_model.predict(X_test_half)

print('numeric data만을 사용하여 얻은 decistion Tree 결과데이터입니다.')
print(confusion_matrix(y_test_half,dt_pred))
print(classification_report(y_test_half,dt_pred))
print('=================================================')

print("support updown time :", time.time() - start)  # 현재시각 - 시작시간 = 실행 시간