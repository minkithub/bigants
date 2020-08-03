'''
Rsi Simulation Code
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
    Support = 'RSI'

    result = pd.DataFrame(columns = ['Code', 'Date', 'Close', 'Start', 'Volume', 'Support', f'{Support}_Cross', f'{Support}_Timing', 
    'SeedMoney', 'TotalMoney', 'Return_Rate', 'Stock_Holdings', 'Money_WinRate', 'Transaction_WinRate', 'Transaction', 'MDD', 'CAGR'])

    for root, dirs, files in os.walk('/Users/minki/pythonworkspace/bigants/dataset/sample'):
        for fname in files:
            full_fname = os.path.join(root, fname)
            data = pd.read_csv(full_fname)
            print('code :', fname[:6])

            # Delete Stop Date of Stock Transaction 
            data = data[data['high'] != 0]
            data = data.reset_index(drop = True)

            # Set Parameters
            n_rsi = stock_value.n_rsi
            n_signal = stock_value.n_signal
            overBuyThres = stock_value.overBuyThres
            overSellThres = stock_value.overSellThres
            args = [n_rsi, n_signal, overBuyThres, overSellThres] # RSI parameters
            short_term = 'RSI'
            long_term = 'RSI_Signal'
            
            # Set Simulation options
            money = stock_value.money
            stock_count = stock_value.stock_count
            buy_ratio = stock_value.buy_ratio
            sell_ratio = stock_value.sell_ratio

            SIM_data = Start_Simulation(data, money, Support, stock_count, args, short_term, long_term, sell_ratio, buy_ratio)

            Code = fname[:6]
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
    
    # result.to_excel('/Users/minki/pythonworkspace/bigants/data_semi_result/RSI_result.xlsx')
    return result, SIM_data

def make_rsi(data, n_rsi, n_signal, overBuyThres, overSellThres, short_term, long_term):
    data = data.loc[::-1]
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
    data = data[['date', 'close', 'start', 'volume', 'price_change', 'price_up', 'price_down', f'AU_{n_rsi}', f'AD_{n_rsi}', short_term, long_term]]

    data.fillna(0)
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

def Set_Timing(data, Support, short_term, overBuyThres, overSellThres):
    # for idx, val in data.iterrows():
    #     while idx > 0:
    #         if data.at[idx, short_term] > overBuyThres and data.at[idx-1, f'{Support}_Cross'] == 'death_cross':
    #             data.at[idx-1, f'{Support}_Timing'] = 'Sell'
    #         elif data.at[idx, short_term] < overSellThres and data.at[idx-1, f'{Support}_Cross'] == 'golden_cross':
    #             data.at[idx-1, f'{Support}_Timing'] = 'Buy'
    #         else:
    #             data.at[idx-1, f'{Support}_Timing'] = 'Wait'
    #         idx -= 1
    #         break
    
    data[f'{Support}_Timing'] = np.where((data[short_term] > overBuyThres) & (data[f'{Support}_Cross'].shift(-1) == 'death_cross'), 'Sell', 
    np.where((data[short_term] < overSellThres) & (data[f'{Support}_Cross'].shift(-1) == 'golden_cross'), 'Buy', 'Wait'))
    data[f'{Support}_Timing'] = data[f'{Support}_Timing'].shift(1)

    return data

def Start_Simulation(data, money, Support, stock_count, args, short_term, long_term, sell_ratio, buy_ratio):
    buy_tax = stock_value.buy_tax
    sell_tax = stock_value.sell_tax

    n_rsi = args[0] # Moving Average of RSI
    n_signal = args[1] # Moving Average of RSI_Signal
    overBuyThres = args[2] # Upper Side of RSI
    overSellThres = args[3] # Lower Side of RSI

    data = make_rsi(data, n_rsi, n_signal, overBuyThres, overSellThres, short_term, long_term)
    data = Set_Cross(data, Support, short_term, long_term)
    data = Set_Timing(data, Support, short_term, overBuyThres, overSellThres)
    data = stock_trade_function.GetRestMoney(data, money, Support, stock_count, buy_tax, sell_tax, sell_ratio, buy_ratio)
    data = stock_trade_function.GetTotalMoney(data, sell_tax, money)
    data = stock_trade_function.GetEarnLoseRate(data)

    return data

if __name__ == "__main__":
    main()