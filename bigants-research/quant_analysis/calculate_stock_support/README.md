# stock_prophet

빅앤츠 주식 지표 계산 코드를 모아놓은 저장소입니다.
각 주식 지표별 퀀트 파라미터는 stock_value.py에 저장되어 있습니다.

또한 최종 파일은 stock_simulation.py를 실행함으로써 얻을 수 있습니다.
stock_simulation.py에서 출력되는 두개의 결과 파일은 다음과 같습니다.

1. combine_data = 각 지표별 Cross(death or golden), Timing(Sell or buy or wait)이 결합된 파일로 최종 스코어 까지 나와있습니다.
score 산출식 = buy(+1) + sell(-1) + wait(0)으로 계산하였습니다.

2. combine_result = 각 지표별 simulation options에 의해서 출력된 결과를 데이터프레임으로 합쳐서 보여주는 파일입니다. combine_result 결과를 통해 주식 지표별 해당 주식의 시뮬레이션 결과를 한눈에 볼 수 있습니다.
