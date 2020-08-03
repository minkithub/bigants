'''
combine stock data
해당 코드를 통해 단독으로 얻은 퀀트 지표들의 시뮬레이션 결과와 시뮬레이션들을 취합합니다.
latest Update : 2020-04-21
by Minki
'''

import numpy as np
import pandas as pd
import datetime as dt
import warnings
warnings.filterwarnings("ignore")
import os
import re

from calculate_stock_support import stock_value
from calculate_stock_support import Bollinger_Simulation
from calculate_stock_support import CCI_Simulation
from calculate_stock_support import DEMA_Simulation
from calculate_stock_support import DMI_Simulation
from calculate_stock_support import EMA_Simulation
from calculate_stock_support import Envelope_Simulation
from calculate_stock_support import MACD_Simulation
from calculate_stock_support import OBV_Simulation
from calculate_stock_support import RSI_Simulation
from calculate_stock_support import SMA_Simulation
from calculate_stock_support import StochasticSlow_Simulation

def main():
    bol_result, bol_data = Bollinger_Simulation.main()
    # print(bol_data.columns)
    print('bollinger completed')
    cci_result, cci_data = CCI_Simulation.main()
    # print(cci_data.columns)
    print('cci completed')
    dema_result, dema_data = DEMA_Simulation.main()
    # print(dema_data.columns)
    print('dema completed')
    dmi_result, dmi_data = DMI_Simulation.main()
    # print(dmi_data.columns)
    print('dmi completed')
    ema_result, ema_data = EMA_Simulation.main()
    # print(ema_data.columns)
    print('ema completed')
    env_result, env_data = Envelope_Simulation.main()
    # print(env_data.columns)
    print('env completed')
    macd_result, macd_data = MACD_Simulation.main()
    # print(macd_data.columns)
    print('macd completed')
    obv_result, obv_data = OBV_Simulation.main()
    # print(obv_data.columns)
    print('obv completed')
    rsi_result, rsi_data = RSI_Simulation.main()
    # print(rsi_data.columns)
    print('rsi completed')
    sma_result, sma_data = SMA_Simulation.main()
    # print(sma_data.columns)
    print('sma completed')
    sto_result, sto_data = StochasticSlow_Simulation.main()
    # print(sto_data.columns)
    print('sto completed')

    # 1. 각 보조지표들의 시뮬레이션 결과를 취합한 데이터 생성
    # support_result : 보조지표 결과들을 취합한 데이터

    val_support = ['Code', 'Date', 'Close', 'Start', 'Volume', 'Support', 'SeedMoney', 'TotalMoney', 'Return_Rate','Stock_Holdings', 
    'Money_WinRate', 'Transaction_WinRate', 'Transaction', 'MDD', 'CAGR']

    support_result = pd.concat([bol_result, cci_result, dema_result, dmi_result, ema_result, env_result, 
    macd_result, obv_result, rsi_result, sma_result, sto_result])

    support_result = support_result[val_support]
    support_result = support_result.sort_values(by=['Return_Rate'], axis=0, ascending=False)

    # 2. 모든 수치 데이터와 Timing을 합친 데이터 생성
    # concat total data : 보조지표의 Cross와 Timing을 통합하여 score까지 뽑아낸 데이터
    # 실제 분석에서는 Cross와 Timing 둘 중 하나만 쓰임.

    basic_numeric = ['start', 'high', 'low', 'volume']

    val_bol_numeric = [f'Envelope_{stock_value.n_Bol}', 'Bol_upper', 'Bol_lower', 'Bol_gap']
    val_bol_timing = ['Bollinger_Timing']

    val_cci_numeric = ['Mean_Price', 'MA_Price', 'MA_abs', 'CCI']
    val_cci_timing = ['CCI_Timing']

    val_dema_numeric = ['dema5', 'dema10', 'dema20', 'dema30', 'dema60', 'dema120']
    val_dema_timing = ['dema5_dema30_Timing', 'dema5_dema60_Timing', 'dema5_dema120_Timing', 'dema10_dema30_Timing',
    'dema10_dema60_Timing', 'dema10_dema120_Timing', 'dema20_dema30_Timing', 'dema20_dema60_Timing', 'dema20_dema120_Timing']

    val_dmi_numeric = ['TR', 'PDI', 'MDI', 'ATR', 'DX', 'ADX', 'ADXR']
    val_dmi_timing = ['DMI_Timing']

    val_ema_numeric = ['ema5', 'ema10', 'ema20', 'ema30', 'ema60', 'ema120']
    val_ema_timing = ['ema5_ema30_Timing', 'ema5_ema60_Timing', 'ema5_ema120_Timing', 'ema10_ema30_Timing',
    'ema10_ema60_Timing', 'ema10_ema120_Timing', 'ema20_ema30_Timing', 'ema20_ema60_Timing', 'ema20_ema120_Timing']

    val_env_numeric = [f'Envelope_{stock_value.n_Env}', 'Envelope_upper', 'Envelope_lower']
    val_env_timing = ['Envelope_Timing']

    val_macd_numeric = [f'MA_{stock_value.short_val}', f'MA_{stock_value.long_val}', 'Macd', 'Macd_Signal', 'Macd_Oscillator']
    val_macd_timing = ['MACD_Timing']
    
    val_obv_numeric = ['OBV']
    val_obv_timing = ['OBV_Timing']

    val_rsi_numeric = [f'AU_{stock_value.n_rsi}', f'AD_{stock_value.n_rsi}', 'RSI', 'RSI_Signal']
    val_rsi_timing = ['RSI_Timing']

    val_sma_numeric = ['sma5', 'sma10', 'sma20', 'sma30', 'sma60', 'sma120']
    val_sma_timing = ['sma5_sma30_Timing', 'sma5_sma60_Timing', 'sma5_sma120_Timing', 'sma10_sma30_Timing', 
    'sma10_sma60_Timing', 'sma10_sma120_Timing', 'sma20_sma30_Timing', 'sma20_sma60_Timing', 'sma20_sma120_Timing']

    val_sto_numeric = ['fastK', 'slowK', 'slowD']
    val_sto_timing = ['Stochastic_Timing']

    concat_basic_data = bol_data[['date', 'start', 'high', 'low', 'volume']]
    concat_bol_data = bol_data[val_bol_numeric + val_bol_timing]
    concat_cci_data = cci_data[val_cci_numeric + val_cci_timing]
    concat_dema_data = dema_data[val_dema_numeric + val_dema_timing]
    concat_dmi_data = dmi_data[val_dmi_numeric + val_dmi_timing]
    concat_ema_data = ema_data[val_ema_numeric + val_ema_timing]
    concat_env_data = env_data[val_env_numeric + val_env_timing]
    concat_macd_data = macd_data[val_macd_numeric + val_macd_timing]
    concat_obv_data = obv_data[val_obv_numeric + val_obv_timing]
    concat_rsi_data = rsi_data[val_rsi_numeric + val_rsi_timing]
    concat_sma_data = sma_data[val_sma_numeric + val_sma_timing]
    concat_sto_data = sto_data[val_sto_numeric + val_sto_timing]

    total_data = pd.concat([concat_basic_data, concat_bol_data, concat_cci_data, concat_dema_data, concat_dmi_data, concat_ema_data, concat_env_data, 
    concat_macd_data, concat_obv_data, concat_rsi_data, concat_sma_data, concat_sto_data], axis = 1)

    total_numeric_data = total_data[['start', 'high', 'low', 'volume'] + val_bol_numeric + val_cci_numeric + val_dema_numeric + val_dmi_numeric + val_ema_numeric +
    val_env_numeric + val_macd_numeric + val_obv_numeric + val_rsi_numeric + val_sma_numeric + val_sto_numeric]

    concat_dema_data['score'] = 0
    for i in range(6, len(concat_ema_data.columns)):
        concat_dema_data['sub_score'] = np.where(concat_dema_data.iloc[:, i] == 'Buy', 1, np.where(concat_dema_data.iloc[:, i] == 'Sell', -1, 0))
        concat_dema_data['score'] = concat_dema_data['score'] + concat_dema_data['sub_score']

    concat_ema_data['score'] = 0
    for i in range(6, len(concat_ema_data.columns)):
        concat_ema_data['sub_score'] = np.where(concat_ema_data.iloc[:, i] == 'Buy', 1, np.where(concat_ema_data.iloc[:, i] == 'Sell', -1, 0))
        concat_ema_data['score'] = concat_ema_data['score'] + concat_ema_data['sub_score']

    concat_sma_data['score'] = 0
    for i in range(6, len(concat_sma_data.columns)):
        concat_sma_data['sub_score'] = np.where(concat_sma_data.iloc[:, i] == 'Buy', 1, np.where(concat_sma_data.iloc[:, i] == 'Sell', -1, 0))
        concat_sma_data['score'] = concat_sma_data['score'] + concat_sma_data['sub_score']

    concat_bol_data['score'] = np.where(concat_bol_data['Bollinger_Timing'] == 'Buy', 1, np.where(concat_bol_data['Bollinger_Timing'] == 'Sell', -1, 0))
    concat_cci_data['score'] = np.where(concat_cci_data['CCI_Timing'] == 'Buy', 1, np.where(concat_cci_data['CCI_Timing'] == 'Sell', -1, 0))
    concat_dmi_data['score'] = np.where(concat_dmi_data['DMI_Timing'] == 'Buy', 1, np.where(concat_dmi_data['DMI_Timing'] == 'Sell', -1, 0))
    concat_env_data['score'] = np.where(concat_env_data['Envelope_Timing'] == 'Buy', 1, np.where(concat_env_data['Envelope_Timing'] == 'Sell', -1, 0))
    concat_macd_data['score'] = np.where(concat_macd_data['MACD_Timing'] == 'Buy', 1, np.where(concat_macd_data['MACD_Timing'] == 'Sell', -1, 0))
    concat_obv_data['score'] = np.where(concat_obv_data['OBV_Timing'] == 'Buy', 1, np.where(concat_obv_data['OBV_Timing'] == 'Sell', -1, 0))
    concat_rsi_data['score'] = np.where(concat_rsi_data['RSI_Timing'] == 'Buy', 1, np.where(concat_rsi_data['RSI_Timing'] == 'Sell', -1, 0))
    concat_sto_data['score'] = np.where(concat_sto_data['Stochastic_Timing'] == 'Buy', 1, np.where(concat_sto_data['Stochastic_Timing'] == 'Sell', -1, 0))

    total_data['score'] = concat_dema_data['score'] + concat_ema_data['score'] + concat_sma_data['score'] + concat_bol_data['score'] + concat_cci_data['score'] + concat_dmi_data['score'] + concat_env_data['score'] + concat_macd_data['score'] + concat_obv_data['score'] + concat_rsi_data['score'] + concat_sto_data['score']
    total_data['close'] = bol_data['close']
    total_data['code'] = bol_result['Code'].loc[0]

    # 3. Binary 분석을 위해 Timing 신호(Sell, Wait, Buy)를 (-1. 0. 1)로 바꾼 데이터를 생성
    # 또한 추가적인 Binary 분석을 위하여 나머지 변수들도 (Down, Stay, Up)을 기준으로 Binary를 생성
    # 추가적으로 고민해야하는 부분은 Timing과 나머지 변수들 사이에 Binary 표기를 동일하게 해도 되는 것인가임.

    total_timing_binary = pd.concat([bol_data[val_bol_timing], cci_data[val_cci_timing], dema_data[val_dema_timing], dmi_data[val_dmi_timing], 
    ema_data[val_ema_timing], env_data[val_env_timing], macd_data[val_macd_timing], obv_data[val_obv_timing], rsi_data[val_rsi_timing], sma_data[val_sma_timing], sto_data[val_sto_timing]], axis = 1)

    total_numeric_binary = pd.concat([bol_data[basic_numeric], bol_data[val_bol_numeric], cci_data[val_cci_numeric], dema_data[val_dema_numeric], dmi_data[val_dmi_numeric], 
    ema_data[val_ema_numeric], env_data[val_env_numeric], macd_data[val_macd_numeric], obv_data[val_obv_numeric], rsi_data[val_rsi_numeric], sma_data[val_sma_numeric], sto_data[val_sto_numeric], total_data[['score']]], axis = 1)

    # 3-1 close 전처리 => (Down, Stay, Up) = (-1, 0, 1)
    total_data['binary_close'] = np.where(total_data['close'] > total_data['close'].shift(-1), -1, np.where(total_data['close'] < total_data['close'].shift(-1), 1, 0))
    total_close_binary = total_data['binary_close']
    
    # 3.2 total_timing_binary 전처리 => (Sell, Wait, Buy) = (-1, 0, 1)

    for i in range(0, len(total_timing_binary.columns)):
        total_timing_binary.iloc[:, i] = np.where(total_timing_binary.iloc[:, i] == 'Buy', 1, np.where(total_timing_binary.iloc[:, i] == 'Sell', -1, 0))

    total_timing_binary = pd.concat([total_timing_binary, total_close_binary], axis = 1) 

    # 3.3 total_numeric_binary 전처리 => (Down, Stay, Up) = (-1, 0, 1)

    for i in range(0, len(total_numeric_binary.columns)):
        total_numeric_binary.iloc[:, i] = np.where(total_numeric_binary.iloc[:, i] > total_numeric_binary.iloc[:, i].shift(1), 1, 
        np.where(total_numeric_binary.iloc[:, i] < total_numeric_binary.iloc[:, i].shift(1), -1, 0))

    total_numeric_binary = pd.concat([total_numeric_binary, total_close_binary], axis = 1)

    # 3.4 total_binary_data 생성
    total_binary_data = pd.concat([total_numeric_binary.drop(['binary_close'],axis=1), total_timing_binary], axis = 1)

    # 3.5 support 지표 original data + Timing 지표 더미변수 데이터.
    total_half_binary_data = pd.concat([total_numeric_data, total_data['score'], total_timing_binary, total_data['close']], axis = 1)

    total_data.to_excel('/Users/minki/pythonworkspace/bigants/data_semi_result/total_data2.xlsx')

    return total_data, support_result, total_timing_binary, total_numeric_binary, total_binary_data, total_half_binary_data

if __name__ == "__main__":
    main()
