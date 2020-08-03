'''
주식 상방, 하방 찾기
2020.03.10
by Minki Kim
'''

# read library
import numpy as np
import pandas as pd
import datetime as dt
import warnings
warnings.filterwarnings("ignore")
from sklearn.metrics import mean_squared_error
import os
from fbprophet import Prophet

# Set 2020 holidays list
holidays_2020 = ['2020-01-01', '2020-01-04', '2020-01-05', '2020-01-11', '2020-01-12', '2020-01-18', '2020-01-19', '2020-01-19', '2020-01-24', '2020-01-25', '2020-01-26', '2020-01-27',
'2020-02-01', '2020-02-02', '2020-02-08', '2020-02-09', '2020-02-15', '2020-02-16', '2020-02-22', '2020-02-23', '2020-02-29',
'2020-03-01', '2020-03-07', '2020-03-08', '2020-03-14', '2020-03-15', '2020-03-21', '2020-03-22', '2020-03-28', '2020-03-29',
'2020-04-04', '2020-04-05', '2020-04-11', '2020-04-12', '2020-04-18', '2020-04-19', '2020-04-25', '2020-04-26', '2020-04-30',
'2020-05-02', '2020-05-03', '2020-05-05', '2020-05-09', '2020-05-10', '2020-05-16', '2020-05-17', '2020-05-23', '2020-05-24', '2020-05-30', '2020-05-31',
'2020-06-06', '2020-06-07', '2020-06-13', '2020-06-14', '2020-06-20', '2020-06-21', '2020-06-27', '2020-06-28',
'2020-07-04', '2020-07-05', '2020-07-11', '2020-07-12', '2020-07-18', '2020-07-19', '2020-07-25', '2020-07-26',
'2020-08-01', '2020-08-02', '2020-08-08', '2020-08-09', '2020-08-15', '2020-08-16', '2020-08-22', '2020-08-23', '2020-08-29', '2020-08-30',
'2020-09-05', '2020-09-06', '2020-09-12', '2020-09-13', '2020-09-19', '2020-09-20', '2020-09-26', '2020-09-27', '2020-09-30',
'2020-10-01', '2020-10-02', '2020-10-03', '2020-10-04', '2020-10-09', '2020-10-10', '2020-10-11', '2020-10-17', '2020-10-18', '2020-10-24', '2020-10-25',
'2020-11-01', '2020-11-07', '2020-11-08', '2020-11-14', '2020-11-15', '2020-11-21', '2020-11-22', '2020-11-28', '2020-11-29',
'2020-12-05', '2020-12-06', '2020-12-12', '2020-12-13', '2020-12-19', '2020-12-20', '2020-12-25', '2020-12-26', '2020-12-27']

# define error

def _error(actual, predicted):
    return actual - predicted

EPSILON = 1e-10

def _percentage_error(actual, predicted):
    return _error(actual, predicted) / (actual + EPSILON)

def mae(actual, predicted):
    return np.mean(np.abs(_error(actual, predicted)))

def mape(actual, predicted):
    return np.mean(np.abs(_percentage_error(actual, predicted)))

# Define Various Moving Average

sma_list = ['sma5', 'sma10', 'sma20', 'sma30', 'sma60', 'sma120']

def make_sma(data):
    data = data.loc[::-1]
    df_close = data[['close']]
    for r in [5, 10, 20, 30, 60, 120]:
        df_close[f'sma{r}'] = data['close'].rolling(window = r).mean()
    df_close = df_close.loc[::-1]
    df_close['date'] = data[['date']]

    df_close['date'] = pd.to_datetime(df_close['date'], errors='coerce')
    df_close['index'] = df_close['date'].dt.strftime('%Y%m%d').astype(int)

    return df_close

ema_list = ['ema5', 'ema10', 'ema20', 'ema30', 'ema60', 'ema120']

def make_ema(data):
    data = data.loc[::-1]
    df_close = data[['close']]
    for r in [5, 10, 20, 30, 60, 120]:
        df_close[f'ema{r}'] = df_close.close.ewm(span = r, adjust = False).mean()
    df_close = df_close.loc[::-1]
    df_close['date'] = data[['date']]

    df_close['date'] = pd.to_datetime(df_close['date'], errors='coerce')
    df_close['index'] = df_close['date'].dt.strftime('%Y%m%d').astype(int)

    return df_close

dema_list = ['dema5', 'dema10', 'dema20', 'dema30', 'dema60', 'dema120']

def make_dema(data):
    data = data.loc[::-1]
    df_close = data[['close']]
    for r in [5, 10, 20, 30, 60, 120]:
        ema_r = df_close.close.ewm(span = r, adjust = False).mean()
        df_close[f'dema{r}'] = ema_r.ewm(span = r, adjust = False).mean()
    df_close = df_close.loc[::-1]
    df_close['date'] = data[['date']]

    df_close['date'] = pd.to_datetime(df_close['date'], errors='coerce')
    df_close['index'] = df_close['date'].dt.strftime('%Y%m%d').astype(int)

    return df_close

total_list = sma_list + ema_list + dema_list
short_ma = ['sma5', 'ema5', 'dema5', 'sma10', 'ema10', 'dema10']
long_ma = ['sma20', 'ema20', 'dema20', 'sma30', 'ema30', 'dema30']

# Making Holidays Index Using Various Moving Average

def idx_cross(data, short_term, long_term, name):
    for idx, val in data.iterrows():
        while idx > 0:
            if data.at[idx, short_term] > data.at[idx, long_term] and data.at[idx-1, short_term] < data.at[idx-1, long_term]:
                data.at[idx, f'cross_{name}'] = 'death_cross'
            elif data.at[idx, short_term] < data.at[idx, long_term] and data.at[idx-1, short_term] > data.at[idx-1, long_term]:
                data.at[idx, f'cross_{name}'] = 'golden_cross'
            else:
                data.at[idx, f'cross_{name}'] = 'nothing'
            idx -= 1
            break

    return data

# Get Prediction and Accuracy

error_total = pd.DataFrame(columns = ['code', 'holidays', 'mae', 'mape', 'price_accuracy', 'direction_accuracy'])
idx_counter = 1
code_counter = 1

# Get 3month Accuracy

predict_periods = 20
prophet_periods = 100

for root, dirs, files in os.walk("/Users/minki/pythonworkspace/bigants/dataset/compare_kosho"):
    for fname in files:
        full_fname = os.path.join(root, fname)
        data = pd.read_csv(full_fname)

        data['date'] = pd.to_datetime(data['date'], errors='coerce')
        data['index'] = data['date'].dt.strftime('%Y%m%d').astype(int)

        # Read data
        df_sma = make_sma(data)
        df_ema = make_ema(data)
        df_dema = make_dema(data)

        # Delete double index
        del df_sma['date']
        del df_sma['close']
        del df_ema['date']
        del df_ema['close']
        del df_dema['date']
        del df_dema['close']
        
        df_total = pd.merge(data, df_sma, on = 'index')
        df_total = pd.merge(df_total, df_ema, on = 'index')
        df_total = pd.merge(df_total, df_dema, on = 'index')
        
        # Make sma data
        for i in range(0, len(total_list)):
            total_id1 = total_list[i]
            for r in range(i+1, len(total_list)):
                total_id2 = total_list[r]
                print('total_id1', total_id1)
                print('total_id2', total_id2)
                cross_list = idx_cross(df_total, total_id1, total_id2, 'holidays')
                
                # Mkae holidays and holidays date
                cross_date = df_total[cross_list['cross_holidays'] != 'nothing']['date']
                names = total_id1 + '_' + total_id2
                cross_date = pd.DataFrame({
                    'holiday' : names,
                    'ds' : pd.to_datetime(list(cross_date)),
                    'lower_window' : 0,
                    'upper_window' : 1
                })
                
                print('number of holidays : ', len(cross_date))
                
                # Make prophet model data
                # 수익률 계산을 위해 시초가(start)와 종가(close)를 예측해야 한다.
                
                ch_scale = 0.05
                inter_scale = 0.95
                ph_df_close = data[['date', 'close']]
                ph_df_start = data[['date', 'start']]
                ph_df_close.rename(columns={'close': 'y', 'date': 'ds'}, inplace=True)
                ph_df_start.rename(columns={'start': 'y', 'date': 'ds'}, inplace=True)

                # print('예측 데이터 셋입니다 ======================================================')
                # print(ph_df.iloc[range(predict_periods, len(ph_df)), :].head())

                # 종가(close) 예측 부분

                m_close = Prophet(changepoint_prior_scale = ch_scale,
                           interval_width = inter_scale,
                           holidays = cross_date)

                m_close.fit(ph_df_close.iloc[range(predict_periods, len(ph_df_close)), :])
                future_prices = m_close.make_future_dataframe(periods=prophet_periods)
                forecast_close = m_close.predict(future_prices)
                
                res_close = forecast_close[['ds', 'yhat', 'yhat_lower', 'yhat_upper']]
                res_close = res_close.rename(columns = {'ds' : 'date', 'yhat' : 'close_pre'})
                res_close['index'] = res_close['date'].dt.strftime('%Y%m%d').astype(int)

                del res_close['yhat_lower']
                del res_close['yhat_upper']

                # 시가(start) 예측 부분

                m_start = Prophet(changepoint_prior_scale = ch_scale,
                           interval_width = inter_scale,
                           holidays = cross_date)

                m_start.fit(ph_df_start.iloc[range(predict_periods, len(ph_df_start)), :])
                future_prices = m_start.make_future_dataframe(periods=prophet_periods)
                forecast_start = m_start.predict(future_prices)
                
                res_start = forecast_start[['ds', 'yhat', 'yhat_lower', 'yhat_upper']]
                res_start = res_start.rename(columns = {'ds' : 'date', 'yhat' : 'start_pre'})
                res_start['index'] = res_start['date'].dt.strftime('%Y%m%d').astype(int)

                del res_start['date']
                del res_start['yhat_lower']
                del res_start['yhat_upper']

                # res = res_close와 res_start를 합친 데이터
                res = pd.merge(res_close, res_start, on = 'index')

                cal_error = pd.merge(res, df_total, on = 'index', how = 'outer')
                del cal_error['date_y']
                cal_error = cal_error.rename(columns = {'date_x' : 'date'})

                print('cal_error =========================================================================================')
                print(cal_error.head())

                # make timing date
                # : 실제값의 끝날 이후로는 예측값이 붙는 것.
                # : 위의 holidays_2020에 있는 날짜들은 제거할 것.

                for idx, val in cal_error.iterrows():
                    cal_date = cal_error.at[idx, 'date']
                    cal_date = str(cal_date)
                    cal_date = cal_date[:10]
                    if cal_date in holidays_2020 :
                        cal_error = cal_error.drop(index = idx)

                # df_total = Original data
                # cal_error = add predict data to df_total
                # df_total의 가장 최근일을 가격을 통해 가격의 단위를 추출

                if df_total.iloc[0]['close']%100 == 0:
                    cal_error['close'] = np.where(pd.notnull(cal_error['close']) == True, cal_error['close'], round(cal_error['close_pre']/100)*100)
                else :
                    cal_error['close'] = np.where(pd.notnull(cal_error['close']) == True, cal_error['close'], round(cal_error['close_pre']))

                if df_total.iloc[0]['start']%100 == 0:
                    cal_error['start'] = np.where(pd.notnull(cal_error['start']) == True, cal_error['start'], round(cal_error['start_pre']/100)*100)
                else :
                    cal_error['start'] = np.where(pd.notnull(cal_error['start']) == True, cal_error['start'], round(cal_error['start_pre']))

                # RSI와 MA를 이용해 timing date를 뽑는 과정
                timing_df = cal_error[['date', 'close', 'start']]
                timing_df['index'] = timing_df['date'].dt.strftime('%Y%m%d').astype(int)
                timing_df = timing_df[::-1]

                timing_df_sma = make_sma(timing_df)
                timing_df_ema = make_ema(timing_df)
                timing_df_dema = make_dema(timing_df)

                # Delete double index
                del timing_df_sma['date']
                del timing_df_sma['close']
                del timing_df_ema['date']
                del timing_df_ema['close']
                del timing_df_dema['date']
                del timing_df_dema['close']
                
                timing_df = pd.merge(timing_df, timing_df_sma, on = 'index')
                timing_df = pd.merge(timing_df, timing_df_ema, on = 'index')
                timing_df = pd.merge(timing_df, timing_df_dema, on = 'index')

                # short_ma와 long_ma를 교차시켜 총 36개의 크로스 지표를 생성
                # golden cross일때는 +1, death cross일때는 -1을 매김
                # 양수 = golden cross 경향 우세 -> 매수
                # 음수 = death cross 경향 우세 -> 매도

                timing_df['cross_score'] = 0
                
                for i in short_ma :
                    for j in long_ma :
                        ma_name = i + j 
                        timing_df = idx_cross(timing_df, i, j, ma_name)
                        timing_df['each_score'] = np.where(timing_df[f'cross_{ma_name}'] == 'golden_cross', 1, np.where(timing_df[f'cross_{ma_name}'] == 'death_cross', -1, 0))
                        timing_df['cross_score'] = timing_df['cross_score'] + timing_df['each_score']

                del timing_df['each_score']

                cross_division = 5

                timing_df['timing_MA'] = np.where(timing_df['cross_score'] >= cross_division , 'Buy', np.where(timing_df['cross_score'] <= -cross_division , 'Sell', 'Stay'))

                result_timing_df = timing_df[['date', 'start', 'close', 'cross_score', 'timing_MA', 'index']]
                result_timing_df.to_excel('/Users/minki/pythonworkspace/bigants/data_semi_result/result_timing_df.xlsx', index = False)

                ## error
                # print('error입니다.===============================================================================================================')
                # mae_res = error(cal_error['close'], cal_error['close_pre'])

                # 수익률 시뮬레이션

                # calculate of price_accuracy
                mae_res = mae(cal_error['close'], cal_error['close_pre'])
                mape_res = mape(cal_error['close'], cal_error['close_pre'])
                price_accuracy = 100 - mape_res*100
                code = fname[:6]

                # calculate of direction_accuracy
                direction_error = cal_error[['close_pre', 'close']].tail(predict_periods+1)
                direction_error = direction_error.diff().iloc[1:len(direction_error), :]

                for i, row in direction_error.iterrows() :
                    if direction_error.at[i, 'close_pre'] * direction_error.at[i, 'close'] > 0:
                        direction_error.at[i, 'res'] = 'correct'
                    else:
                        direction_error.at[i, 'res'] = 'wrong'

                direction_accuracy = len(direction_error[direction_error['res'] == 'correct']['res'])/len(direction_error) * 100

                error_total = error_total.append(pd.DataFrame([[code, names, mae_res, mape_res, price_accuracy, direction_accuracy]], 
                                                              columns = ['code', 'holidays', 'mae', 'mape', 'price_accuracy', 'direction_accuracy']))
                
                print('direction_accuracy:', direction_accuracy)
                print('accuracy :', price_accuracy)
                idx_counter += 1
                print('idx_counter : ', idx_counter)
        code_counter += 1
        print('code_counter : ', code_counter)
        print("=================================================================================================")

error_total.to_excel('/Users/minki/pythonworkspace/bigants/stock_prophet/dataset/error_total.xlsx')