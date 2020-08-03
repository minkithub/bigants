import React from 'react';
import { createFragmentContainer, graphql } from 'react-relay';
import { View, Text, FlatList, StyleSheet } from 'react-native';
import { AreaChart } from 'react-native-svg-charts'

import { PredictionResultHeader_prediction } from './__generated__/PredictionResultHeader_prediction.graphql';
import { Path, Defs, ClipPath, Rect } from 'react-native-svg';
import { numf, datef } from '../_components/fmt';
import { theme } from '../theme';

export const PredictionResultHeader = createFragmentContainer<{
  prediction: PredictionResultHeader_prediction,
}>(({ relay, prediction }) => {

  const preds = prediction.dailyPredictions.filter(pred => {
    const d = new Date(pred.date).getDay();
    return 0 < d && d < 6
  }); // 주말제거 (나중에 검토)

  const data = [
    ...prediction.stock.history.map(h => h.c),
    ...preds.map(p => p.expectedPrice),
  ];

  const indexToClipFrom = prediction.stock.history.length - 1;

  const Clips = ({ x, width }: any) => (
    <Defs key={'clips'}>
      <ClipPath id={'clip-path-1'} key={'0'}>
        <Rect x={0} y={'0'} width={x(indexToClipFrom)} height={'100%'} />
      </ClipPath>
      <ClipPath id='clip-path-2' key={'1'}>
        <Rect x={x(indexToClipFrom)} y={'0'} width={width - x(indexToClipFrom)} height={'100%'} />
      </ClipPath>
    </Defs>
  )
  const Line = ({ line }: any) => (
    <Path key={'line'} d={line} stroke={'green'} fill={'none'} clipPath={'url(#clip-path-1)'} />
  )
  const DashedLine = ({ line }: any) => (
    <Path
      key={'dashed-line'}
      stroke={'green'}
      d={line}
      fill={'none'}
      clipPath={'url(#clip-path-2)'}
      strokeDasharray={[4, 4]}
    />
  )


  return (
    <View style={styles.root}>
      <View style={styles.header}>
        <View style={{ flexDirection: 'row', alignItems: 'center', justifyContent: 'space-between' }}>
          <Text style={{ fontSize: 24, fontWeight: 'bold', }}>{prediction.stock.nameKo}</Text>
          <Text style={theme.mnText}>종가 {numf(prediction.stock.latestHistory.close)}</Text>
        </View>
        <View style={{ flexDirection: 'row', paddingTop: 16 }}>
          <Text style={{ paddingTop: 8, fontSize: 16, color: '#353637' }}>예측평균수익률</Text>
          <Text style={{ paddingLeft: 20, paddingTop: 8, fontSize: 16, fontWeight: 'bold', color: '#FC0000' }}>{
            numf(prediction.averageIncome).slice(0, 4)
          }%</Text>
        </View>
        <View style={{ flexDirection: 'row' }}>
          <Text style={{ paddingTop: 8, fontSize: 16, color: '#353637' }}>모델 예측 정확도</Text>
          <Text style={{ paddingLeft: 20, paddingTop: 8, fontSize: 16, fontWeight: 'bold', color: '#FC0000' }}>{
            numf(prediction.accuracy).slice(0, 4)
          }%</Text>
        </View>
      </View>
      <AreaChart style={{ height: 140, paddingVertical: theme.MU.V1 }} svg={{ fill: 'white' }} data={data} >
        <Clips />
        <Line />
        <DashedLine />
      </AreaChart>
      <FlatList
        style={{ flex: 1, flexShrink: 1 }}
        data={preds}
        ListHeaderComponent={
          <View style={styles.contentItem}>
            <Text style={theme.mnText}>날짜</Text>
            <View style={styles.result}>
              <Text style={theme.mnText}>예측 가격</Text>
              <Text style={[theme.mnText, { paddingLeft: theme.MU.H2 }]}>예상수익률</Text>
            </View>
          </View>
        }
        renderItem={({ item }) => (
          <View key={item.date} style={styles.contentItem}>
            <Text>{datef(new Date(item.date)).slice(5, 10)}</Text>
            <View style={styles.result}>
              <Text style={theme.mnText}>{numf(item.expectedPrice).slice(0, 7)}</Text>
              <Text style={[theme.mnText, { paddingLeft: theme.MU.H2 }]}>{numf(item.expectedIncome).slice(0, 3)}%</Text>
            </View>
          </View>
        )}

      />
    </View>
  )
}, {
  prediction: graphql`
    fragment PredictionResultHeader_prediction on StockPrediction {
      id
      stock {
        nameKo
        latestHistory {
          close
        }
        history(count: 25) {
          d: date
          c: close
        }
      }
      averageIncome
      accuracy
      dailyPredictions {
        date
        expectedPrice
        expectedIncome
      }
    }
  `
})

const styles = StyleSheet.create({
  root: {
    flex: 1,
  },
  header: {
    paddingHorizontal: theme.MU.H2,
    paddingTop: theme.MU.V1,
    paddingBottom: theme.MU.V2,
    backgroundColor: 'white',
  },
  contentItem: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    height: 56,
    marginHorizontal: theme.MU.H2,
    paddingLeft: theme.MU.H1,
    borderBottomColor: '#EFEFEF',
    borderBottomWidth: 1
  },
  result: {
    flexDirection: 'row',
    alignItems: 'center',
  }
})