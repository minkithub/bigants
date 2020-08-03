import numpy as np
import pandas as pd
import datetime as dt
from fbprophet import Prophet
import warnings
warnings.filterwarnings("ignore")

import os

def _error(actual, predicted):
    return actual - predicted

EPSILON = 1e-10

def _percentage_error(actual, predicted):
    return _error(actual, predicted) / (actual + EPSILON)

def mae(actual, predicted):
    return np.mean(np.abs(_error(actual, predicted)))

def mape(actual, predicted):
    return np.mean(np.abs(_percentage_error(actual, predicted)))

# SMA index(Simple Moving Average)
# 주가의 평균 추세를 파악하기 위해 사용되는 지표

sma_list = ['sma5', 'sma10', 'sma20', 'sma30', 'sma60', 'sma120']

def make_sma(data):
    data = data.loc[::-1]
    df_close = data[['close']]
    for r in [5, 10, 20, 30, 60, 120]:
        df_close[f'sma{r}'] = data['close'].rolling(window = r).mean()
    df_close = df_close.loc[::-1]
    df_close['date'] = data[['date']]
    return df_close

# EMA index(Exponential Moving Average)
# 주로 주가의 단기적인 추세를 파악하기 위해 사용되는 지표

ema_list = ['ema5', 'ema10', 'ema20', 'ema30', 'ema60', 'ema120']

def make_ema(data):
    data = data.loc[::-1]
    df_close = data[['close']]
    for r in [5, 10, 20, 30, 60, 120]:
        df_close[f'ema{r}'] = df_close.close.ewm(span = r, adjust = False).mean()
    df_close = df_close.loc[::-1]
    df_close['date'] = data[['date']]
    return df_close

# DEMA index(Double Exponential Moving Average)
# EMA보다 더 빠르게 주가의 추세를 반영하기 위한 지표

dema_list = ['dema5', 'dema10', 'dema20', 'dema30', 'dema60', 'dema120']

def make_dema(data):
    data = data.loc[::-1]
    df_close = data[['close']]
    for r in [5, 10, 20, 30, 60, 120]:
        ema_r = df_close.close.ewm(span = r, adjust = False).mean()
        df_close[f'dema{r}'] = ema_r.ewm(span = r, adjust = False).mean()
    df_close = df_close.loc[::-1]
    df_close['date'] = data[['date']]
    return df_close

total_list = sma_list + ema_list + dema_list

# sma_index는 모두 string형태로 처리해야 한다.

def idx_cross(data, index1, index2):
    cross = pd.DataFrame()
    cross['date'] = data['date']
    cross['boolean'] = data[index1] - data[index2]
    cross['res'] = 'blank'
    
    for r in range(1, len(cross)+1):
        index = len(cross) - r

        if cross['boolean'][index] > 0 and cross['boolean'][index+1] < 0 :
            cross['res'][index] = 'golden cross'

        elif cross['boolean'][index] < 0 and cross['boolean'][index+1] > 0 :
            cross['res'][index] = 'death cross'

        else:
            cross['res'][index] = 'nothing'
        
    return cross

# sma_index는 모두 string형태로 처리해야 한다.

# 집 : "/Users/minki/pythonworkspace/bigants/dataset/data"
# 복지관 : "C:\\Users\\janga\\Desktop\\workspace\\data"

error_total = pd.DataFrame(columns = ['code', 'holidays', 'mae', 'mape', 'accuracy'])
idx_counter = 1
code_counter = 1

predict_periods = 60
prophet_periods = 100

for root, dirs, files in os.walk("/Users/minki/pythonworkspace/bigants/dataset/data"):
    for fname in files:
        full_fname = os.path.join(root, fname)
        data = pd.read_csv(full_fname)
        print('code : ', full_fname)
        
        # Read data
        df_sma = make_sma(data)
        df_ema = make_ema(data)
        df_dema = make_dema(data)
        
        del df_sma['date']
        del df_ema['close']
        del df_ema['date']
        del df_dema['close']
        
        df_close = pd.concat([df_sma, df_ema, df_dema], axis=1)
        
        # Make sma data
        for i in range(0, len(total_list)):
            total_id1 = total_list[i]
            for r in range(i+1, len(total_list)):
                total_id2 = total_list[r]
                print('total_id1', total_id1)
                print('total_id2', total_id2)
                df_cross = idx_cross(df_close, total_id1, total_id2)
                
                # Mkae holidays and holidays date
                cross_date = df_cross[df_cross['res'] != 'nothing']['date']
                names = total_id1 + '_' + total_id2
                cross_date = pd.DataFrame({
                    'holiday' : names,
                    'ds' : pd.to_datetime(list(cross_date)),
                    'lower_window' : 0,
                    'upper_window' : 1
                })
                
                print('number of holidays : ', len(cross_date))
                
                # Make prophet model data
                
                ch_scale = 0.05
                inter_scale = 0.95
                ph_df = data[['date', 'close']]
                ph_df.rename(columns={'close': 'y', 'date': 'ds'}, inplace=True)
                m = Prophet(changepoint_prior_scale = ch_scale,
                           interval_width = inter_scale,
                           holidays = cross_date)
                m.fit(ph_df.iloc[range(predict_periods, len(ph_df)), :])
                future_prices = m.make_future_dataframe(periods=prophet_periods)
                forecast = m.predict(future_prices)
                
                res = forecast[['ds', 'yhat', 'yhat_lower', 'yhat_upper']]
                res = res.rename(columns = {'ds' : 'date', 'yhat' : 'close_pre'})
                ## res = res.loc[::-1]
                ## res = res.reset_index(drop = True, inplace = False)
                res['index'] = res['date'].dt.strftime('%Y%m%d').astype(int)
                
                # data = original data
                data['date'] = pd.to_datetime(data['date'], errors='coerce')
                data['index'] = data['date'].dt.strftime('%Y%m%d').astype(int)
                
                # calculate of error
                cal_error = pd.merge(res, data, on = 'index')
                ## del cal_error['date_y']
                ## del cal_error['index']
                ## cal_error = cal_error.rename(columns = {'date_x' : 'date'})
                
                mae_res = mae(cal_error['close'], cal_error['close_pre'])
                mape_res = mape(cal_error['close'], cal_error['close_pre'])
                accuracy = 100 - mape_res*100
                code = fname[:6]
                
                error_total = error_total.append(pd.DataFrame([[code, names, mae_res, mape_res, accuracy]], 
                                                              columns = ['code', 'holidays', 'mae', 'mape', 'accuracy']))
                
                print('accuracy :', accuracy)
                
                idx_counter += 1
                print('idx_counter : ', idx_counter)
                
        code_counter += 1
        print('code_counter : ', code_counter)
                
        print("=================================================================================================")