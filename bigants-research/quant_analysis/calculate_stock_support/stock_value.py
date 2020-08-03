'''
Set of stock value
latest Update : 2020-03-31
by Minki
'''

# Set Simulation options
money = 10000000
stock_count = 100
sell_ratio = 0.5
buy_ratio = 0.5
buy_tax = 0.00015
sell_tax = 0.00315

# Set bollinger options
n_Bol = 20
n_Bol_std = 2

# Set CCI options
n_CCI = 14

# Set DMI options
n_Dmi = 14
n_Adx = 6

# Set Env options
n_Env = 20
n_upper = 6
n_lower = 5

# Set MACD options
short_val = 12
long_val = 26
t_val = 9

# Set OBV options
n_OBV = 14

# Set RSI options
n_rsi = 14 # Moving Average of RSI
n_signal = 6 # Moving Average of RSI_Signal
overBuyThres = 70 # Upper Side of RSI
overSellThres = 30 # Lower Side of RSI

# Set StochasticSlow options
n_fast = 12 # Moving Average of StochasticSlow
n_slowk = 3 # Moving Average of StochasticSlow Fast
n_slowd = 3 # Moving Average of StochasticSlow Fast