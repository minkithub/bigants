'''
MACD Simulation Code
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
    Support = 'MACD'

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

            # Set Simulation options
            money = stock_value.money
            stock_count = stock_value.stock_count
            buy_ratio = stock_value.buy_ratio
            sell_ratio = stock_value.sell_ratio

            # Set Parameters
            short_val = stock_value.short_val
            long_val = stock_value.long_val
            t_val = stock_value.t_val
            args = [short_val, long_val, t_val]

            data = data.loc[::-1]
            SIM_data = Start_Simulation(data, money, Support, stock_count, args, sell_ratio, buy_ratio)

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

def Make_macd(data, args):
    short_val = args[0]
    long_val = args[1]
    t_val = args[2]

    data[f'MA_{short_val}'] = data['close'].ewm(span=short_val, min_periods = 1).mean()
    data[f'MA_{long_val}'] = data['close'].ewm(span=long_val, min_periods = 1).mean()
    data['Macd'] = data[f'MA_{short_val}'] - data[f'MA_{long_val}']
    data['Macd_Signal'] = data['Macd'].ewm(span=t_val, min_periods = 1).mean()
    data['Macd_Oscillator'] = data['Macd'] - data['Macd_Signal']
    
    data.fillna(0)
    return data


def Set_Cross(data, Support):
    # for idx, val in data.iterrows():
    #     while idx > 0:
    #         if data.at[idx, 'Macd'] > 0 and data.at[idx-1, 'Macd'] < 0:
    #             data.at[idx-1, f'{Support}_Cross'] = 'death_cross'
    #         elif data.at[idx, 'Macd'] < 0 and data.at[idx-1, 'Macd'] > 0:
    #             data.at[idx-1, f'{Support}_Cross'] = 'golden_cross'
    #         elif data.at[idx, 'Macd'] > data.at[idx, 'Macd_Signal'] and data.at[idx-1, 'Macd'] < data.at[idx-1, 'Macd_Signal']:
    #             data.at[idx-1, f'{Support}_Cross'] = 'death_cross'
    #         elif data.at[idx, 'Macd'] < data.at[idx, 'Macd_Signal'] and data.at[idx-1, 'Macd'] > data.at[idx-1, 'Macd_Signal']:
    #             data.at[idx-1, f'{Support}_Cross'] = 'golden_cross'
    #         else:
    #             data.at[idx-1, f'{Support}_Cross'] = 'nothing'
    #         idx -= 1
    #         break
    
    data[f'{Support}_Cross'] = np.where((data['Macd'] > 0 ) & (data['Macd'].shift(-1) < 0) | (data['Macd'] > data['Macd_Signal']) & (data['Macd'].shift(-1) < data['Macd_Signal'].shift(-1)), 'death_cross',
    np.where((data['Macd'] < 0 ) & (data['Macd'].shift(-1) > 0) | (data['Macd'] < data['Macd_Signal']) & (data['Macd'].shift(-1) > data['Macd_Signal'].shift(-1)), 'golden_cross', 'nothing'))
    data[f'{Support}_Cross'] = data[f'{Support}_Cross'].shift(1)

    return data

def Set_Timing(data, Support):
    # for idx, val in data.iterrows():
    #     if data.at[idx, f'{Support}_Cross'] == 'death_cross':
    #         data.at[idx, f'{Support}_Timing'] = 'Sell'
    #     elif data.at[idx, f'{Support}_Cross'] == 'golden_cross':
    #         data.at[idx, f'{Support}_Timing'] = 'Buy'
    #     else:
    #         data.at[idx, f'{Support}_Timing'] = 'Wait'

    data[f'{Support}_Timing'] = np.where(data[f'{Support}_Cross'] == 'death_cross', 'Sell', np.where(data[f'{Support}_Cross'] == 'golden_cross', 'Buy', 'Wait'))

    return data

def Start_Simulation(data, money, Support, stock_count, args, sell_ratio, buy_ratio):
    buy_tax = stock_value.buy_tax
    sell_tax = stock_value.sell_tax

    data = Make_macd(data, args)
    data = Set_Cross(data, Support)
    data = Set_Timing(data, Support)
    data = stock_trade_function.GetRestMoney(data, money, Support, stock_count, buy_tax, sell_tax, sell_ratio, buy_ratio)
    data = stock_trade_function.GetTotalMoney(data, sell_tax, money)
    data = stock_trade_function.GetEarnLoseRate(data)
    
    return data

if __name__ == "__main__":
    main()