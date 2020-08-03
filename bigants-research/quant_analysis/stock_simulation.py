'''
stock simulation
해당 코드는 보조지표의 시그널을 스코어로 치환한 후 스코어의 모든 조합에 따라 주식 시뮬레이션 결과를 살펴볼 수 있는 코드입니다.
또한 combine_stock_data.main()을 통해 얻은 support_result를 통해 퀀트 단독의 시뮬레이션 결과도 살펴볼 수 있습니다.
latest Update : 2020-04-21
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
import time
start = time.time()  # 시작 시간 저장

def main():
    
    # Set Simulation options
    money = stock_value.money
    stock_count = stock_value.stock_count
    buy_ratio = stock_value.buy_ratio
    sell_ratio = stock_value.sell_ratio
    sell_tax = stock_value.sell_tax
    buy_tax = stock_value.buy_tax

    # import data
    combine_data, support_result, total_timing_binary, total_numeric_binary, total_binary_data, total_half_binary_data = combine_stock_data.main()
    val_col = ['code', 'date', 'close', 'start', 'high', 'low', 'volume', 'score']
    data = combine_data[val_col]

    # Get buy_score list, sell_score list
    max_score = max(data['score'])
    min_score = min(data['score'])
    num_trade = len(data[data['score']!=0])
    trade_ratio = num_trade/len(data)*100

    kinda_sell_score = set(data[data['score']<0]['score'])
    kinda_sell_score = list(kinda_sell_score)

    kinda_buy_score = set(data[data['score']>0]['score'])
    kinda_buy_score = list (kinda_buy_score)

    result = pd.DataFrame(columns = ['Code', 'Date', 'Close', 'Start', 'Volume', 'Support', 'Sell_score', 'Buy_score', 'SeedMoney',
    'TotalMoney', 'Return_Rate', 'Stock_Holdings', 'Money_WinRate', 'Transaction_WinRate', 'Transaction', 'MDD', 'CAGR'])

    for Sell_score in kinda_sell_score:
        for Buy_score in kinda_buy_score:
            print('Sell_score :', Sell_score, 'Buy_score :', Buy_score)
            data = GetRestMoney(data, money, stock_count, buy_tax, sell_tax, sell_ratio, buy_ratio, Sell_score, Buy_score)
            data = stock_trade_function.GetTotalMoney(data, sell_tax, money)
            data = stock_trade_function.GetEarnLoseRate(data)

            Code = data.iloc[-1]['code']
            Date = data.iloc[-1]['date']
            Close = data.iloc[-1]['close']
            Start = data.iloc[-1]['start']
            Volume = data.iloc[-1]['volume']
            Support = 'Null'

            Sell_score = Sell_score
            Buy_score = Buy_score
            SeedMoney = data.iloc[0]['TotalMoney']
            TotalMoney = data.iloc[-1]['TotalMoney']
            Return_Rate = data.iloc[-1]['Return_Rate']
            Stock_Holdings = data.iloc[-1]['stock_count']
            
            actual_transaction, buy_win, sell_win = Get_actionWinrate(data, Sell_score, Buy_score)
            Transaction_WinRate = (buy_win + sell_win) / actual_transaction * 100
            money_transaction, money_win = stock_trade_function.Get_Moneyrate(data, SeedMoney)
            Money_WinRate = money_win/money_transaction*100
            Transaction = actual_transaction

            MDD = min(data['TotalMoney'])/SeedMoney*100
            CAGR = (data.iloc[-1]['close']/data.iloc[0]['close'])**(1/(len(data)/249)) - 1

            result = result.append(pd.DataFrame([[Code, Date, Close, Start, Volume, Support, Sell_score, Buy_score, 
            SeedMoney, TotalMoney, Return_Rate, Stock_Holdings, Money_WinRate, Transaction_WinRate, Transaction, MDD, CAGR]],
            columns = ['Code', 'Date', 'Close', 'Start', 'Volume', 'Support', 'Sell_score', 'Buy_score', 'SeedMoney', 'TotalMoney', 'Return_Rate', 'Stock_Holdings', 
            'Money_WinRate', 'Transaction_WinRate', 'Transaction', 'MDD', 'CAGR']))

            result = result.sort_values(by=['Return_Rate'], axis=0, ascending=False)
    
    print("simulation time :", time.time() - start)  # 현재시각 - 시작시간 = 실행 시간
    return data, result

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

def GetRestMoney(data, money, stock_count, buy_tax, sell_tax, sell_ratio, buy_ratio, Sell_score, Buy_score):
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

    score_indexes = list(data[data['score'] != 0].index)

    for index in score_indexes:
        if index > 0:
            if data['score'].loc[index] == Sell_score:
                sell_price = data['start'].loc[index-1]
                stock_count, sell_cost, returnedMoney, restMoney = Sell(stock_count, sell_price, restMoney, sell_tax, sell_ratio)

                data['stock_count'].loc[index-1] = stock_count
                data['sell_price'].loc[index-1] = sell_cost
                data['returnedMoney'].loc[index-1] = returnedMoney
                data['restMoney'].loc[index-1] = restMoney
            
            if data['score'].loc[index] == Buy_score:
                buy_price = data['start'].loc[index-1]
                stock_count, buy_cost, purchasedCost, restMoney = Buy(stock_count, buy_price, restMoney, buy_tax, buy_ratio)

                data['stock_count'].loc[index-1] = stock_count
                data['buy_price'].loc[index-1] = buy_cost
                data['purchasedCost'].loc[index-1] = purchasedCost
                data['restMoney'].loc[index-1] = restMoney

            if data['restMoney'].loc[index-1] != data['restMoney'].loc[index]:
                data.loc[range(0, index-1), 'restMoney'] = data['restMoney'].loc[index-1]

            if data['stock_count'].loc[index-1] != data['stock_count'].loc[index]:
                data.loc[range(0, index-1), 'stock_count'] = data['stock_count'].loc[index-1]
    
    data['isTrading'] = np.where(data['stock_count'].shift(-1) < data['stock_count'], 1, np.where(data['stock_count'].shift(-1) > data['stock_count'], 2, 0))
    data['isTrading'] = data['isTrading'].shift(1)

    return data

def Get_actionWinrate(data, Sell_score, Buy_score):
    sell_transaction = len(data[data['score'] == Sell_score])
    buy_transaction = len(data[data['score'] == Buy_score])
    actual_transaction = sell_transaction + buy_transaction
    
    buy_win = np.where((data['score'] == Buy_score) & (data['sell_price'].shift(-1) < data['close'].shift(-1)), 1, 0).cumsum()[-1]
    sell_win = np.where((data['score'] == Sell_score) & (data['sell_price'].shift(-1) > data['close'].shift(-1)), 1, 0).cumsum()[-1]

    return actual_transaction, buy_win, sell_win

if __name__ == "__main__":
    main()