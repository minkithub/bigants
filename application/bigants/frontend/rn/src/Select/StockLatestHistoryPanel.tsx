import React from 'react';
import { createFragmentContainer, graphql, RelayProp } from 'react-relay';
import { View, Text, StyleSheet } from 'react-native';
import { StockLatestHistoryPanel_stock } from './__generated__/StockLatestHistoryPanel_stock.graphql';
import { numf } from '../_components/fmt';
import { theme } from '../theme';


export const StockLatestHistoryPanel = createFragmentContainer<{
  stock: StockLatestHistoryPanel_stock,
  relay: RelayProp,
}>(({ stock }) => {
  const { latestHistory } = stock;


  return (
    <View style={styles.root}>
      <View style={styles.left}>
        <Text style={theme.xlbbText}>{stock.nameKo}</Text>
        <Text style={[theme.xxlbText, { color: latestHistory.priceChange >= 0 ? '#FC0000' : '#2F80ED' }]}>{
          numf(latestHistory.close)
        }</Text>
        <Text style={theme.snText}>{latestHistory.date} 기준</Text>
        <Text style={[theme.snText, { paddingTop: 4 }]}>전일대비 <Text style={{ fontWeight: 'bold', color: latestHistory.priceChange >= 0 ? '#FC0000' : '#2F80ED' }}>{
          latestHistory.priceChange >= 0 ? '▲' : '▼'
        } {
            numf(Math.abs(latestHistory.priceChange))
          }</Text> | <Text style={{ fontWeight: 'bold', color: latestHistory.priceChange >= 0 ? '#FC0000' : '#2F80ED' }}>
            {latestHistory.priceChange >= 0 ? '+' : '-'}{
              numf(Math.abs(latestHistory.priceChangeRate - 1) * 100).slice(0, 4) // 양수일 때 반드시 +를 표시하기 위해서
            }%</Text></Text>
      </View>
      <View style={styles.right}>
        <Text style={theme.snText}>거래량 <Text style={theme.sbbText}>{numf(latestHistory.volume)}</Text></Text>
        <View style={styles.hlPriceInfo}>
          <Text style={theme.snText}>고가 <Text style={{ color: '#FC0000', fontWeight: 'bold', }}>{numf(latestHistory.high)}</Text></Text>
          <Text style={theme.snText}>(상한 {numf(latestHistory.highLimit)})</Text>
        </View>
        <View style={styles.hlPriceInfo}>
          <Text style={theme.snText}>저가 <Text style={{ color: '#2F80ED', fontWeight: 'bold', }}>{numf(latestHistory.low)}</Text></Text>
          <Text style={theme.snText}>(하한 {numf(latestHistory.lowLimit)})</Text>
        </View>
      </View>
    </View>
  );
}, {
  stock: graphql`
    fragment StockLatestHistoryPanel_stock on Stock {
      id
      nameKo
      latestHistory {
        date
        close
        volume
        high
        low
        highLimit
        lowLimit
        priceChange
        priceChangeRate
      }
    }
  `
})

const styles = StyleSheet.create({
  root: {
    flexDirection: 'row',
    paddingHorizontal: theme.MU.H2,
    paddingBottom: theme.MU.V2,
  },

  left: {
    flex: 3,
  },
  right: {
    flex: 2,
    alignItems: 'flex-end',
    paddingTop: theme.MU.V1,
  },

  dateInfo: {
    flexDirection: 'row',
    justifyContent: 'center',
  },

  //right
  hlPriceInfo: {
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    paddingTop: 4,
  }
})