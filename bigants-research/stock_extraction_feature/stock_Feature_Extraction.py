'''
PCA Analytics Code
latest Update : 2020-04-02
by Minki

주식 지표의 값으로 PCA를 진행하기 위한 코드로
PCA를 위한 타겟 변수로는 종가의 변동이다.
주가의 변동은 Up, Down, NoChange로 구분되어 있으며
이는 각각 1, -1, 0으로 인코딩 되어있다.
'''

import numpy as np
import pandas as pd
import datetime as dt
import warnings
from sklearn.preprocessing import MinMaxScaler
from sklearn.feature_selection import SelectKBest, chi2, f_classif, f_regression, mutual_info_classif
warnings.filterwarnings("ignore")
import re
import os


def main():
    # Bollingerband parameters
    args_Bol = [20, 2]

    # CCI parameters
    args_CCI = [14]

    # DMI parameters
    args_DMI = [14, 6]

    # Envelope parameters
    args_Env = [20, 6, 5]

    # Macd parameters
    args_Macd = [12, 26, 9]

    # Rsi parameters
    args_Rsi = [14, 6, 'RSI', 'RSI_Signal']

    # Stochastic_Slow parameters
    args_Sto = [12, 3, 3]

    data = pd.read_csv('/Users/minki/pythonworkspace/bigants/dataset/sample/005380.csv')
    data = data.loc[::-1]
    data = Produce_data(data, args_Bol, args_CCI, args_DMI, args_Env, args_Macd, args_Rsi, args_Sto)

    data_only_support = data[['Bol_upper', 'Bol_lower', 'CCI', 'dema5', 'dema10', 'dema20', 'dema30', 'dema60', 'dema120',
    'PDI', 'MDI', 'ADX', 'ema5', 'ema10', 'ema20', 'ema30', 'ema60', 'ema120', 'Envelope_upper', 'Envelope_lower', 'Macd', 'Macd_Signal',
    'OBV', 'RSI', 'RSI_Signal', 'sma5', 'sma10', 'sma20', 'sma30', 'sma60', 'sma120', 'slowK', 'slowD', 'close']]

    # data normalization to feature extraction
    data_only_support_noraml = Min_Max(data_only_support)

    # Change close value to Up & Down & NoChange, Up = 1, Down = -1, NoChange = 0
    data_only_support_noraml['close_encoding'] = np.where(data_only_support_noraml['close'] > data_only_support_noraml['close'].shift(1), 1, 
    np.where(data_only_support_noraml['close'] < data_only_support_noraml['close'].shift(1), -1, 0))

    # Feature Extraction Using SelectBest
    data_SelectBest_chi = Extraction_SelectBest(data_only_support_noraml, chi2)
    data_SelectBest_f_classif = Extraction_SelectBest(data_only_support_noraml, f_classif)
    data_SelectBest_f_regression = Extraction_SelectBest(data_only_support_noraml, f_regression)
    data_SelectBest_mutual_info_classif = Extraction_SelectBest(data_only_support_noraml, mutual_info_classif)

    # adjust Min_Max sclaer to Extraction score
    data_SelectBest_chi['min_max_score_chi'] =  Min_Max(data_SelectBest_chi.loc[:, ['chi2_Score']])
    data_SelectBest_f_classif['min_max_score_f_classif'] =  Min_Max(data_SelectBest_f_classif.loc[:, ['f_classif_Score']])
    data_SelectBest_f_regression['min_max_score_f_regression'] =  Min_Max(data_SelectBest_f_regression.loc[:, ['f_regression_Score']])
    data_SelectBest_mutual_info_classif['min_max_score_mutual_info_classif'] =  Min_Max(data_SelectBest_mutual_info_classif.loc[:, ['mutual_info_classif_Score']])

    data_SelectBest_total = pd.merge(data_SelectBest_chi, data_SelectBest_f_classif, on = 'Specs')
    data_SelectBest_total = pd.merge(data_SelectBest_total, data_SelectBest_f_regression, on = 'Specs')
    data_SelectBest_total = pd.merge(data_SelectBest_total, data_SelectBest_mutual_info_classif, on = 'Specs')
    data_SelectBest_total['total_score'] = data_SelectBest_chi['min_max_score_chi'] + data_SelectBest_f_classif['min_max_score_f_classif'] + data_SelectBest_f_regression['min_max_score_f_regression'] + data_SelectBest_mutual_info_classif['min_max_score_mutual_info_classif']
    data_SelectBest_total = data_SelectBest_total.sort_values(by = 'total_score', ascending = False)

    data_SelectBest_total.to_excel('/Users/minki/pythonworkspace/bigants/data_semi_result/data_SelectBest_total.xlsx')
    print('done')

    return data_only_support, data_SelectBest_total

def Min_Max(data):
    scaler = MinMaxScaler()
    data[:] = scaler.fit_transform(data[:])

    return data

def Extraction_SelectBest(data, Ext_method):
    name = re.findall('function (\S\S+)', str(Ext_method))[0]
    data = data.fillna(0)
    col_num = len(data.columns)
    X = data.iloc[:, 0:(col_num-2)]
    Y = data.iloc[:, -1]

    bestfeatures = SelectKBest(score_func=Ext_method, k=10)
    fit = bestfeatures.fit(X,Y)
    dfscores = pd.DataFrame(fit.scores_)
    dfcolumns = pd.DataFrame(X.columns)

    featureScores = pd.concat([dfcolumns,dfscores],axis=1)
    featureScores.columns = ['Specs',f'{name}_Score']
    featureScores = featureScores.sort_values(by = f'{name}_Score', ascending = False)

    return featureScores

def Make_Bollinger(data, args_Bol):
    n_Bol = args_Bol[0]
    n_Std = args_Bol[1]

    data[f'Envelope_{n_Bol}'] = data['close'].rolling(window = n_Bol, min_periods = 1).mean()
    data['Bol_upper'] = data[f'Envelope_{n_Bol}'] + n_Std*data[f'Envelope_{n_Bol}'].rolling(window = n_Bol, min_periods = 1).std()
    data['Bol_lower'] = data[f'Envelope_{n_Bol}'] - n_Std*data[f'Envelope_{n_Bol}'].rolling(window = n_Bol, min_periods = 1).std()
    data['Bol_gap'] = data['Bol_upper'] - data['Bol_lower']

    data.fillna(0)
    return data

def Make_CCI(data, args_CCI):
    n_CCI = args_CCI[0]

    data['Mean_Price'] = (data['close'] + data['high'] + data['low'])/3
    data['MA_Price'] = data['Mean_Price'].rolling(window = n_CCI, min_periods = 1).mean()
    data['MA_abs'] = abs(data['Mean_Price'] - data['MA_Price']).rolling(window = n_CCI, min_periods = 1).mean()
    data['CCI'] = (data['Mean_Price'] - data['MA_Price'])/(data['MA_abs']*0.015)
    
    data.fillna(0)
    return data

def Make_DEMA(data):
    for r in [5, 10, 20, 30, 60, 120]:
        ema_r = data.close.ewm(span = r, adjust = False, min_periods = 1).mean()
        data[f'dema{r}'] = ema_r.ewm(span = r, adjust = False, min_periods = 1).mean()

    data.fillna(0)
    return data

def Make_DMI(data, args_DMI):
    n_Dmi = args_DMI[0]
    n_Adx = args_DMI[1]

    for idx, val in data.iterrows():
        while idx < len(data)-1:
            TR_1 = data.at[idx, 'high'] - data.at[idx, 'low']
            TR_2 = data.at[idx, 'high'] - data.at[idx+1, 'close']
            TR_3 = data.at[idx, 'low'] - data.at[idx+1, 'close']
            data.at[idx, 'TR'] = max(TR_1, TR_2, TR_3)

            up_move = data.at[idx, 'high'] - data.at[idx+1, 'high']
            down_move = data.at[idx+1, 'low'] - data.at[idx, 'low']
            if up_move > down_move and up_move>0:
                data.at[idx, 'PDI'] = up_move
            else:
                data.at[idx, 'PDI'] = 0
            if down_move > up_move and down_move>0:
                data.at[idx, 'MDI'] = down_move
            else:
                data.at[idx, 'MDI'] = 0
            
            idx += 1
            break

    data['ATR'] = data['TR'].ewm(span = n_Dmi, min_periods = 1).mean()
    data['PDI'] = (data['PDI'].ewm(span = n_Dmi, min_periods = 1).mean())/data['ATR']*100
    data['MDI'] = (data['MDI'].ewm(span = n_Dmi, min_periods = 1).mean())/data['ATR']*100
    data['DX'] = abs(data['PDI'] - data['MDI'])/(data['PDI'] + data['MDI'])*100
    data['ADX'] = data['DX'].ewm(span = n_Adx, min_periods = 1).mean()
    data['ADXR'] = data['ADX'].ewm(span = n_Adx, min_periods = 1).mean()

    return data

def MAKE_EMA(data):
    for r in [5, 10, 20, 30, 60, 120]:
        data[f'ema{r}'] = data.close.ewm(span = r, adjust = False, min_periods = 1).mean()
    
    data.fillna(0)
    return data

def Make_Envelope(data, args_Env):
    n_Env = args_Env[0]
    n_upper = args_Env[1]
    n_lower = args_Env[2] 
    data[f'Envelope_{n_Env}'] = data['close'].rolling(window = n_Env, min_periods = 1).mean()
    data['Envelope_upper'] = data[f'Envelope_{n_Env}']*(1 + n_upper/100)
    data['Envelope_lower'] = data[f'Envelope_{n_Env}']*(1 + n_lower/100)

    data.fillna(0)
    return data

def Make_macd(data, args_Macd):
    short_val = args_Macd[0]
    long_val = args_Macd[1]
    t_val = args_Macd[2]

    data['MA_12'] = data['close'].ewm(span=short_val, min_periods = 1).mean()
    data['MA_26'] = data['close'].ewm(span=long_val, min_periods = 1).mean()
    data['Macd'] = data['MA_12'] - data['MA_26']
    data['Macd_Signal'] = data['Macd'].ewm(span=t_val, min_periods = 1).mean()
    data['Macd_Oscillator'] = data['Macd'] - data['Macd_Signal']
    
    data.fillna(0)
    return data

def Make_OBV(data):
    data['OBV'] = np.where(data['close'] > data['close'].shift(1), data['volume'], 
    np.where(data['close'] < data['close'].shift(1), -data['volume'], 0)).cumsum()

    return data

def Make_RSI(data, args_Rsi):
    n_rsi = args_Rsi[0]
    n_signal = args_Rsi[1]
    short_term = args_Rsi[2]
    long_term = args_Rsi[3]

    data['price_change'] = data['close'].diff()
    data = data.fillna(0)

    data['price_up'] = data[data['price_change'] >= 0]['price_change']
    data['price_down'] = abs(data[data['price_change'] < 0]['price_change'])
    data = data.fillna(0)

    division_idx = (len(data) - n_rsi)

    for idx, val in data.iterrows():
        if idx > division_idx:
            data.at[idx, f'AU_{n_rsi}'] = 0
            data.at[idx, f'AD_{n_rsi}'] = 0
        elif idx == division_idx:
            data.at[idx, f'AU_{n_rsi}'] = sum(data.head(n_rsi)['price_up'])/n_rsi
            data.at[idx, f'AD_{n_rsi}'] = sum(data.head(n_rsi)['price_down'])/n_rsi
        elif idx < division_idx:
            data.at[idx, f'AU_{n_rsi}'] = (data.at[idx+1, f'AU_{n_rsi}']*(n_rsi-1) + val['price_up'])/n_rsi
            data.at[idx, f'AD_{n_rsi}'] = (data.at[idx+1, f'AD_{n_rsi}']*(n_rsi-1) + val['price_down'])/n_rsi
            
        data.at[idx, short_term] = (data.at[idx, f'AU_{n_rsi}'])*100/(data.at[idx, f'AU_{n_rsi}'] + data.at[idx, f'AD_{n_rsi}'])
    
    data = data.fillna(0)
    data[long_term] = data[short_term].rolling(window = n_signal).mean()
    data.fillna(0)

    return data

def Make_SMA(data):
    for r in [5, 10, 20, 30, 60, 120]:
        data[f'sma{r}'] = data['close'].rolling(window = r, min_periods = 1).mean()

    data.fillna(0)
    return data

def Make_StochasticSlow(data, args_Sto):
    n_fast = args_Sto[0]
    n_slowk = args_Sto[1]
    n_slowd = args_Sto[2]

    maxV = data['high'].rolling(n_fast, min_periods = 1).max()
    minV = data['low'].rolling(n_fast, min_periods = 1).min()
    maxV.fillna(0)
    minV.fillna(0)
    
    data['fastK'] = (data['close'] - minV) / (maxV - minV)
    data['slowK'] = data['fastK'].rolling(n_slowk, min_periods = 1).mean()
    data['slowD'] = data['slowK'].rolling(n_slowd, min_periods = 1).mean()

    return data

def Produce_data(data, args_Bol, args_CCI, args_DMI, args_Env, args_Macd, args_Rsi, args_Sto) :
    data = Make_Bollinger(data, args_Bol)
    data = Make_CCI(data, args_CCI)
    data = Make_DEMA(data)
    data = Make_DMI(data, args_DMI)
    data = MAKE_EMA(data)
    data = Make_Envelope(data, args_Env)
    data = Make_macd(data, args_Macd)
    data = Make_OBV(data)
    data = Make_RSI(data, args_Rsi)
    data = Make_SMA(data)
    data = Make_StochasticSlow(data, args_Sto)

    return data

if __name__ == "__main__":
    main()

