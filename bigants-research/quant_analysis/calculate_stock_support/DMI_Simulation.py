'''
DMI Simulation Code
latest Update : 2020-03-31
by Minki
'''

import numpy as np
import pandas as pd
import datetime as dt
import warnings
warnings.filterwarnings("ignore")
import os
import stock_value
import stock_trade_function

def main():
    Support = 'DMI'

    result = pd.DataFrame(columns = ['Code', 'Date', 'Close', 'Start', 'Volume', 'Support', f'{Support}_Cross', f'{Support}_Timing', 'SeedMoney',
    'TotalMoney', 'Return_Rate', 'Stock_Holdings', 'Money_WinRate', 'Transaction_WinRate', 'Transaction', 'MDD', 'CAGR'])

    for root, dirs, files in os.walk('/Users/minki/pythonworkspace/bigants/dataset/sample'):
        for fname in files:
            full_fname = os.path.join(root, fname)
            data = pd.read_csv(full_fname)
            Code = fname[:6]
            print('code :', fname[:6])

            # Delete Stop Date of Stock Transaction 
            data = data[data['high'] != 0]
            data = data.reset_index(drop = True)
            
            # Set Parameters
            n_Dmi = stock_value.n_Dmi
            n_Adx = stock_value.n_Adx
            short_term = 'PDI'
            long_term = 'MDI'

            # Set Simulation options
            money = stock_value.money
            stock_count = stock_value.stock_count
            buy_ratio = stock_value.buy_ratio
            sell_ratio = stock_value.sell_ratio

            data = data.loc[::-1]
            SIM_data = Start_Simulation(data, money, Support, stock_count, n_Dmi, n_Adx, short_term, long_term, sell_ratio, buy_ratio)
            
            Date = SIM_data.iloc[-1]['date']
            Close = SIM_data.iloc[-1]['close']
            Start = SIM_data.iloc[-1]['start']
            Volume = SIM_data.iloc[-1]['volume']

            Cross = SIM_data.iloc[-1][f'{Support}_Cross']
            Timing = SIM_data.iloc[-1][f'{Support}_Timing']
            SeedMoney = SIM_data.iloc[0]['TotalMoney']
            TotalMoney = SIM_data.iloc[-1]['TotalMoney']
            Return_Rate = SIM_data.iloc[-1]['Return_Rate']
            Stock_Holdings = SIM_data.iloc[-1]['stock_count']
            
            actual_transaction, buy_win, sell_win = stock_trade_function.Get_actionWinrate(SIM_data, Support)
            Transaction_WinRate = (buy_win + sell_win) / actual_transaction * 100
            money_transaction, money_win = stock_trade_function.Get_Moneyrate(SIM_data, SeedMoney)
            Money_WinRate = money_win/money_transaction*100
            Transaction = actual_transaction

            MDD = min(SIM_data['TotalMoney'])/SeedMoney*100
            CAGR = (SIM_data.iloc[-1]['close']/SIM_data.iloc[0]['close'])**(1/(len(data)/249)) - 1

            result = result.append(pd.DataFrame([[Code, Date, Close, Start, Volume, Support, Cross, Timing, SeedMoney, TotalMoney, 
            Return_Rate, Stock_Holdings, Money_WinRate, Transaction_WinRate, Transaction, MDD, CAGR]],
            columns = ['Code', 'Date', 'Close', 'Start', 'Volume', 'Support', f'{Support}_Cross', f'{Support}_Timing', 'SeedMoney', 
            'TotalMoney', 'Return_Rate', 'Stock_Holdings', 'Money_WinRate', 'Transaction_WinRate', 'Transaction', 'MDD', 'CAGR']))
      
    return result, SIM_data

def Make_DMI(data, n_Dmi, n_Adx):
    # for idx, val in data.iterrows():
    #     while idx < len(data)-1:
    #         TR_1 = data.at[idx, 'high'] - data.at[idx, 'low']
    #         TR_2 = data.at[idx, 'high'] - data.at[idx+1, 'close']
    #         TR_3 = data.at[idx, 'low'] - data.at[idx+1, 'close']
    #         data.at[idx, 'TR'] = max(TR_1, TR_2, TR_3)

    #         up_move = data.at[idx, 'high'] - data.at[idx+1, 'high']
    #         down_move = data.at[idx+1, 'low'] - data.at[idx, 'low']
    #         if up_move > down_move and up_move>0:
    #             data.at[idx, 'PDI'] = up_move
    #         else:
    #             data.at[idx, 'PDI'] = 0
    #         if down_move > up_move and down_move>0:
    #             data.at[idx, 'MDI'] = down_move
    #         else:
    #             data.at[idx, 'MDI'] = 0
            
    #         idx += 1
    #         break

    data['TR_1'] = data['high'] - data['low']
    data['TR_2'] = data['high'] - data['close'].shift(1)
    data['TR_3'] = data['low'] - data['close'].shift(1)
    data['TR'] = data[['TR_1', 'TR_2', 'TR_3']].max(axis = 1)

    del data['TR_1']
    del data['TR_2']
    del data['TR_3']

    data['up_move'] = data['high'] - data['high'].shift(1)
    data['down_move'] = data['low'].shift(1) - data['low']

    data['PDI'] = np.where((data['up_move'] > data['down_move']) & (data['up_move'] > 0), data['up_move'], 0)
    data['MDI'] = np.where((data['down_move'] > data['up_move']) & (data['down_move'] > 0), data['down_move'], 0)

    del data['up_move']
    del data['down_move']

    data['ATR'] = data['TR'].ewm(span = n_Dmi, min_periods = 1).mean()
    data['PDI'] = (data['PDI'].ewm(span = n_Dmi, min_periods = 1).mean())/data['ATR']*100
    data['MDI'] = (data['MDI'].ewm(span = n_Dmi, min_periods = 1).mean())/data['ATR']*100
    data['DX'] = abs(data['PDI'] - data['MDI'])/(data['PDI'] + data['MDI'])*100
    data['ADX'] = data['DX'].ewm(span = n_Adx, min_periods = 1).mean()
    data['ADXR'] = data['ADX'].ewm(span = n_Adx, min_periods = 1).mean()

    return data

def Set_Cross(data, Support, short_term, long_term):
    # for idx, val in data.iterrows():
    #     while idx > 0:
    #         if data.at[idx, short_term] > data.at[idx, long_term] and data.at[idx-1, short_term] < data.at[idx-1, long_term]:
    #             data.at[idx-1, f'{Support}_Cross'] = 'death_cross'
    #         elif data.at[idx, short_term] < data.at[idx, long_term] and data.at[idx-1, short_term] > data.at[idx-1, long_term]:
    #             data.at[idx-1, f'{Support}_Cross'] = 'golden_cross'
    #         else:
    #             data.at[idx-1, f'{Support}_Cross'] = 'nothing'
    #         idx -= 1
    #         break

    data[f'{Support}_Cross'] = np.where((data[short_term] > data[long_term]) & (data[short_term].shift(-1) < data[long_term].shift(-1)), 'death_cross',
    np.where((data[short_term] < data[long_term]) & (data[short_term].shift(-1) > data[long_term].shift(-1)), 'golden_cross', 'nothing'))
    data[f'{Support}_Cross'] = data[f'{Support}_Cross'].shift(1)

    return data

def Set_Timing(data, Support):
    # for idx, val in data.iterrows():
    #     while idx < len(data)-1:
    #         if data.at[idx, f'{Support}_Cross'] == 'death_cross' and data.at[idx+1, 'ADX'] <= data.at[idx, 'ADX']:
    #             data.at[idx, f'{Support}_Timing'] = 'Sell'
    #         elif data.at[idx, f'{Support}_Cross'] == 'golden_cross' and data.at[idx+1, 'ADX'] <= data.at[idx, 'ADX']:
    #             data.at[idx, f'{Support}_Timing'] = 'Buy'
    #         else:
    #             data.at[idx, f'{Support}_Timing'] = 'Wait'
    #         break

    data[f'{Support}_Timing'] = np.where((data[f'{Support}_Cross'] == 'death_cross') & (data['ADX'].shift(1) <= data['ADX']), 'Sell', 
    np.where((data[f'{Support}_Cross'] == 'golden_cross') & (data['ADX'].shift(1) <= data['ADX']), 'Buy', 'Wait'))

    return data

def Start_Simulation(data, money, Support, stock_count, n_Dmi, n_Adx, short_term, long_term, sell_ratio, buy_ratio):
    buy_tax = stock_value.buy_tax
    sell_tax = stock_value.sell_tax

    data = Make_DMI(data, n_Dmi, n_Adx)
    data = Set_Cross(data, Support, short_term, long_term)
    data = Set_Timing(data, Support)
    data = stock_trade_function.GetRestMoney(data, money, Support, stock_count, buy_tax, sell_tax, sell_ratio, buy_ratio)
    data = stock_trade_function.GetTotalMoney(data, sell_tax, money)
    data = stock_trade_function.GetEarnLoseRate(data)

    return data

if __name__ == "__main__":
    main()

