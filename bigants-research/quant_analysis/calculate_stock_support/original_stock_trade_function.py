'''
주식의 공통된 함수들이 정리되어 있는 파일입니다.
latest Update : 2020-03-31
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
import stock_value

def Buy(stock_count, price, restMoney, tax, buy_ratio):
    buy_cost = int(float(price) + price*tax)
    add_stock_count = (restMoney*buy_ratio)//buy_cost
    stock_count = stock_count + add_stock_count
    purchasedcost = add_stock_count*buy_cost
    restMoney = restMoney-purchasedcost

    return stock_count, buy_cost, purchasedcost, restMoney

def Sell(stock_count, price, restMoney, tax, sell_ratio):
    sell_cost = int(price - float(price)*tax)
    returnedMoney = (int(stock_count*sell_ratio))*sell_cost
    restMoney = restMoney + returnedMoney
    stock_count = stock_count - int(stock_count*sell_ratio)

    return stock_count, sell_cost, returnedMoney, restMoney

def GetRestMoney(data, money, Support, stock_count, buy_tax, sell_tax, sell_ratio, buy_ratio):
    # set conditions
    restMoney = money
    stock_count = stock_count

    # make placehold
    data['stock_count'] = stock_count
    data['sell_price'] = 0
    data['buy_price'] = 0
    data['returnedMoney'] = 0
    data['purchasedCost'] = 0
    data['restMoney'] = money
    data['isTrading'] = 0 # 0: wait, 1: sell, 2 : buy

    for index, val in data.iterrows():
        while index > 0:
            if val[f'{Support}_Timing'] == 'Sell':
                sell_price = data.at[index-1, 'start']
                stock_count, sell_cost, returnedMoney, restMoney = Sell(stock_count, sell_price, restMoney, sell_tax, sell_ratio)

                data.at[index-1,'stock_count'] = stock_count
                data.at[index-1,'sell_price'] = sell_cost
                data.at[index-1,'returnedMoney'] = returnedMoney
                data.at[index-1,'restMoney'] = restMoney

                if data.at[index-1,'stock_count'] < data.at[index,'stock_count']:
                    data.at[index-1,'isTrading'] = 1
            
            if val[f'{Support}_Timing'] == 'Buy':
                buy_price = data.at[index-1, 'start']
                stock_count, buy_cost, purchasedCost, restMoney = Buy(stock_count, buy_price, restMoney, buy_tax, buy_ratio)

                data.at[index-1,'stock_count'] = stock_count
                data.at[index-1,'buy_price'] = buy_cost
                data.at[index-1,'purchasedCost'] = purchasedCost
                data.at[index-1,'restMoney'] = restMoney

                if data.at[index-1,'stock_count'] > data.at[index,'stock_count']:
                    data.at[index-1,'isTrading'] = 2
                    
            if data.at[index-1,'restMoney'] != data.at[index,'restMoney']:
                data.loc[range(0, index-1), 'restMoney'] = data.at[index-1,'restMoney']

            if data.at[index-1,'stock_count'] != data.at[index,'stock_count']:
                data.loc[range(0, index-1), 'stock_count'] = data.at[index-1,'stock_count']
            index -= 1
            break

    return data

def GetTotalMoney(data, sell_tax, money):
    data['TotalMoney'] = (data['close'] - (data['close']*sell_tax).astype(int))*data['stock_count'] + data['restMoney']
    data['Return_Rate'] = round((data['TotalMoney']-data['TotalMoney'].iloc[0])/data['TotalMoney'].iloc[0]*100, 2)

    return data

def GetEarnLoseRate(data):
    # data['LoseNum'] = 0
    # data['EarnNum'] = 0
    # data['LoseNum'] = np.where(data['Return_Rate'] < 0, 1, 0)
    # data['EarnNum'] = np.where(data['Return_Rate'] > 0, 1, 0)

    # LoseNum = 0
    # EarnNum = 0

    # for idx, val in data.iterrows():
    #     LoseNum = LoseNum + data.at[idx, 'LoseNum']
    #     EarnNum = EarnNum + data.at[idx, 'EarnNum']

    #     data.at[idx, 'Accumulated_LoseNum'] = LoseNum
    #     data.at[idx, 'Accumulated_EarnNum'] = EarnNum

    data['Accumulated_LoseNum'] = np.where(data['Return_Rate'] < 0, 1, 0).cumsum()
    data['Accumulated_EarnNum'] = np.where(data['Return_Rate'] > 0, 1, 0).cumsum()

    return data

def Get_actionWinrate(data, Support):
    actual_transaction = len(data[data[f'{Support}_Timing']!= 'Wait'])
    
    # buy_win = 0
    # sell_win = 0

    # for idx, val in data.iterrows():
    #     while idx > 0:
    #         if data.at[idx, f'{Support}_Timing'] == 'Buy' and data.at[idx-1, 'sell_price'] < data.at[idx-1, 'close']:
    #             buy_win += 1
    #         elif data.at[idx, f'{Support}_Timing'] == 'Sell' and data.at[idx-1, 'sell_price'] > data.at[idx-1, 'close']:
    #             sell_win += 1
    #         idx -= 1
    #         break

    buy_win = np.where((data[f'{Support}_Timing'] == 'Buy') & (data['sell_price'].shift(-1) < data['close'].shift(-1)), 1, 0).cumsum()[-1]
    sell_win = np.where((data[f'{Support}_Timing'] == 'Sell') & (data['sell_price'].shift(-1) > data['close'].shift(-1)), 1, 0).cumsum()[-1]

    return actual_transaction, buy_win, sell_win

def Get_Moneyrate(data, money):
    money_transaction = len(data[data['isTrading']!=0])
    
    # money_win = 0
    # for idx, val in data.iterrows():
    #         if data.at[idx, 'isTrading']!=0 and data.at[idx, 'TotalMoney'] > money:
    #             money_win += 1

    money_win = np.where((data['isTrading'] != 0) & (data['TotalMoney'] > money), 1, 0).cumsum()[-1]

    return money_transaction, money_win