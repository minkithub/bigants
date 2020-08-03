import React from 'react';
import { View, Dimensions, Text, StyleSheet } from 'react-native';
import { AreaChart } from 'react-native-svg-charts'
import { createFragmentContainer, graphql } from 'react-relay';
import { HolidaySetForm_stock } from './__generated__/HolidaySetForm_stock.graphql';
import { ScrollView, TouchableOpacity, FlatList } from 'react-native-gesture-handler';

import { Line } from 'react-native-svg';
import { numf } from '../_components/fmt';
import { theme } from '../theme';

type HolidaySetFormProps = {
  holidays: Date[],
  stock: HolidaySetForm_stock,
  onChange: (state: HolidaySetFormState) => void,
}

type HolidaySetFormState = {
  data: number[],
  selectedIndex?: number,
  holidayIndexes: number[],
}

class HolidaySetFormInner extends React.Component<HolidaySetFormProps, HolidaySetFormState> {
  state: HolidaySetFormState = {
    data: [],
    holidayIndexes: [],
    selectedIndex: 0,
  };
  now = new Date();

  static getDerivedStateFromProps(props: HolidaySetFormProps) {
    return {
      data: props.stock.history.map(item => item.c),
    }
  }


  chartWidth = Dimensions.get('window').width - 40 // onLayout에 의해 계속 바뀌는 것 주의

  updateIndex(position: number) {
    const index = Math.round(this.props.stock.history.length * position);
    if (typeof this.state.selectedIndex !== 'number') {
      return;
    }
    const holidayIndexes = [
      ...this.state.holidayIndexes.slice(0, this.state.selectedIndex),
      index,
      ...this.state.holidayIndexes.slice(this.state.selectedIndex + 1),
    ]
    this.setState({ holidayIndexes });
  }

  dropHoliday(index: number) {
    const holidayIndexes = [
      ...this.state.holidayIndexes.slice(0, index),
      ...this.state.holidayIndexes.slice(index + 1),
    ]
    this.setState({ holidayIndexes, selectedIndex: undefined });
  }

  prependHoliday() {
    const holidayIndexes = [
      Math.round(this.state.data.length / 2),
      ...this.state.holidayIndexes,
    ]
    this.setState({ holidayIndexes, selectedIndex: 0 });
  }

  render() {
    return (
      <View style={{ flex: 1 }}>
        <AreaChart
          style={{ height: 140 }}
          data={this.state.data}
          svg={{
            stroke: '#aaa',
            fill: '#f2f2f2',
            onLayout: e => { this.chartWidth = e.nativeEvent.layout.width },
            onPress: e => this.updateIndex(e.nativeEvent.locationX / this.chartWidth),
            onResponderMove: e => this.updateIndex(e.nativeEvent.locationX / this.chartWidth)
          }}
          contentInset={{ top: 20, bottom: 20, right: -1 }}
        >
          {
            this.state.holidayIndexes.map((val, index) => (
              <Line
                key={index}
                x1={this.chartWidth * val / this.state.data.length}
                y1={0}
                x2={this.chartWidth * val / this.state.data.length}
                y2='100%'
                stroke={index === this.state.selectedIndex ? 'red' : 'gray'}
              />
            ))
          }
        </AreaChart>
        <View style={{ height: 20, flexDirection: 'row', justifyContent: 'space-between', backgroundColor: '#f2f2f2' }}>
          <Text style={theme.xsnText}>{this.props.stock.history[0].d}</Text>
          <Text style={theme.xsnText}>{this.props.stock.history[this.state.data.length - 1].d}</Text>
        </View>
        <FlatList
          style={{ flex: 1, backgroundColor: '#f2f2f2' }}
          data={this.state.holidayIndexes}
          stickyHeaderIndices={[0]}
          contentContainerStyle={{ paddingBottom: 120 }} // 이 플랫리스트가 SafeView 안에 쓰이지 않아서
          ListHeaderComponent={
            <View style={{
              paddingLeft: theme.MU.H2, paddingRight: theme.MU.H1, paddingVertical: theme.MU.V1, flexDirection: 'row', justifyContent: 'space-between',
              alignItems: 'center',
              borderBottomColor: '#ddd', borderBottomWidth: 1,
              backgroundColor: '#f2f2f2',
            }}>
              <Text style={theme.mnText}>날짜 (종가, 거래량)</Text>
              <TouchableOpacity onPress={() => this.prependHoliday()} style={styles.addHolidayButton}>
                <Text style={theme.sbText}>+ 호재일/악재일 추가</Text>
              </TouchableOpacity>
            </View>
          }
          renderItem={({ item: val, index }) => (
            <TouchableOpacity
              key={index}
              style={{
                flexDirection: 'row',
                justifyContent: 'space-between',
                padding: 16,
                borderBottomColor: 'gray',
                backgroundColor: this.state.selectedIndex === index ? '#ddd' : 'white',
              }}
              onPress={() => this.setState({ selectedIndex: index })}
            >
              <Text>{this.props.stock.history[val].d} ({numf(this.props.stock.history[val].c)}, {numf(this.props.stock.history[val].v)})</Text>
              <TouchableOpacity onPress={() => this.dropHoliday(index)}>
                <Text>지우기</Text>
              </TouchableOpacity>
            </TouchableOpacity>
          )}
        />
      </View>
    )
  }
}

export const HolidaySetForm = createFragmentContainer(HolidaySetFormInner, {
  stock: graphql`
    fragment HolidaySetForm_stock on Stock {
      history {
        d: date
        c: close
        v: volume
      }
    }
  `
})


const styles = StyleSheet.create({
  addHolidayButton: {
    flexDirection: 'column',
    justifyContent: 'center',
    backgroundColor: theme.colors.background.roundButton,
    borderRadius: 16,
    height: 36,
    paddingHorizontal: theme.MU.H1,
    borderWidth: 1,
    borderColor: theme.colors.primary.black,
  }
})