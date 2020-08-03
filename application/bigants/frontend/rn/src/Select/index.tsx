import React from 'react';
import { View, Text, TouchableOpacity, ActivityIndicator, Modal, Dimensions, Alert, Animated, Easing, Image, StyleSheet } from 'react-native';
import { QueryRenderer, graphql } from 'react-relay';
import { relayEnvironment } from '../_lib/relay';
import { StockLatestHistoryPanel } from './StockLatestHistoryPanel';
import { SelectScreenQuery } from './__generated__/SelectScreenQuery.graphql';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParams } from '../route'
import { RouteProp } from '@react-navigation/native';
import { HolidaySetForm } from './HolidaySetForm';
import { theme } from '../theme';
import { predictionCreate } from '../_mutations/PredictionCreate';
import { datef } from '../_components/fmt';
import { Tx } from '../_components/async'


type State = {
  holidays: Date[],
  infoModalVisible: boolean,
};

const rotateImage = require('../assets/isPredicting.png');
const predictingIcon = require('../assets/isPredictingIcon.png');

class SelectScreen extends React.Component<{
  route: RouteProp<RootStackParams, 'Select'>
  navigation: StackNavigationProp<RootStackParams, 'Select'>
}, State> {
  state: State = { holidays: [], infoModalVisible: true };

  static screenOptions = {
    title: '데이터 조정',
    headerStyle: { borderBottomColor: 'transparent' },
    headerTitleAlign: "left",
    headerTitleStyle: { ...theme.lnText, marginLeft: 0 },
    headerBackTitle: '',
    headerBackTitleVisible: false,
    headerLeftContainerStyle: { marginLeft: theme.MU.H1 },
    headerBackground: () => <View style={{ backgroundColor: 'white', height: '100%' }} />,
  }

  private RotateValueHolder = new Animated.Value(0);

  componentDidMount() {
    this.StartImageRotate();
  }

  StartImageRotate() {
    this.RotateValueHolder.setValue(0);
    Animated.timing(this.RotateValueHolder, {
      toValue: 1,
      duration: 3000,
      easing: Easing.linear,
    }).start(() => this.StartImageRotate());
  }

  render() {
    const { navigation, route } = this.props;
    const RotateData = this.RotateValueHolder.interpolate({
      inputRange: [0, 1],
      outputRange: ['0deg', '360deg'],
    });

    return (
      <QueryRenderer<SelectScreenQuery>
        query={graphql`
        query SelectScreenQuery($stockId: ID) {
          stock: node(id: $stockId) {
            ...StockLatestHistoryPanel_stock
            ...HolidaySetForm_stock
          }
        }
      `}
        environment={relayEnvironment}
        variables={{ stockId: route.params.stockId }}
        render={({ error, retry, props }) => {
          if (error) {
            throw error;
          }
          if (!props) {
            return <ActivityIndicator style={{ flex: 1 }} />
          }
          if (!props.stock) {
            return (
              <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
                <Text>현재 해당 종목을 이용할 수 없습니다.</Text>
                <Text>다음에 다시 시도해주세요.</Text>
              </View>
            )
          }
          return (
            <Tx render={predictionCreateTx => (
              <View style={styles.root}>
                <Modal
                  animationType='none'
                  visible={this.state.infoModalVisible}
                  presentationStyle='overFullScreen'
                  transparent
                >
                  <TouchableOpacity style={styles.introductionModal} onPress={() => this.setState({ infoModalVisible: false })} activeOpacity={1}>
                    <View style={{ width: 300, backgroundColor: 'white', paddingHorizontal: theme.MU.H2, paddingBottom: theme.MU.V2, paddingTop: theme.MU.V3, borderRadius: 18, }}>
                      <Text style={theme.lnText}>주가예측 시 모델이 신경써야하는 주가의 날짜를 입력해주세요. 모델의 정확도가 올라갈 수 있습니다.</Text>
                      <Text style={[theme.lbblueText, { marginTop: theme.MU.V2, textAlign: 'center' }]}>확인</Text>
                    </View>
                  </TouchableOpacity>
                </Modal>

                <Modal
                  animated
                  visible={predictionCreateTx.inProgress}
                  presentationStyle='overFullScreen'
                >
                  <View style={styles.predictingModal}>
                    <Text style={[theme.xlbbText, { paddingBottom: theme.MU.V1 }]}>모델이 주가를 예측 중입니다</Text>
                    <Text style={theme.mnText}>네트워크 환경에 따라 일정 시간이 소요될 수 있습니다.</Text>
                    <Text style={theme.mnText}>잠시만 기다려주세요!</Text>
                    <View style={{ marginTop: theme.MU.V2 }}>
                      <Animated.Image
                        style={{ width: 130, height: 130, transform: [{ rotate: RotateData }] }}
                        source={rotateImage}
                      />
                      <View style={{ position: 'absolute', left: 15, top: 13, zIndex: 2, justifyContent: 'center', alignItems: 'center' }}>
                        <Image source={predictingIcon} style={{ width: 100, height: 100 }} />
                      </View>
                    </View>
                    <View style={{ position: 'absolute', right: 20, top: 32, }}>
                      <Text style={theme.mbblueText}>취소</Text>
                    </View>
                  </View>
                </Modal>
                <View>
                  <StockLatestHistoryPanel stock={props.stock} />
                </View>
                <HolidaySetForm stock={props.stock} holidays={this.state.holidays} onChange={console.log} />
                <View style={{
                  position: 'absolute',
                  bottom: 40,
                  left: Dimensions.get('window').width / 2 - 64,
                }}>
                  <TouchableOpacity
                    disabled={predictionCreateTx.inProgress}
                    onPress={predictionCreateTx.as(this.doPredictionCreate)}
                    style={{ width: 128, height: 44, borderRadius: 8, justifyContent: 'center', alignItems: 'center', backgroundColor: '#0084F4' }}>
                    <Text style={{ color: '#fff', fontSize: 16, fontWeight: 'bold' }}>예측하기</Text>
                  </TouchableOpacity>
                </View>
              </View>
            )} />
          );
        }}
      />
    )
  }

  private doPredictionCreate = async () => {
    const stockId = this.props.route.params.stockId;
    try {
      const res = await predictionCreate(relayEnvironment, {
        holidays: this.state.holidays.map(datef),
        stockId,
      })
      this.props.navigation.navigate('PredictionResult', {
        stockPredictionId: res.predictionCreate.result.node.id,
      })
    } catch (error) {
      Alert.alert(`죄송합니다. 종목(${stockId})의 분석 중에 오류가 발생했습니다. 일시적인 오류일 수 있습니다.`);
    }
  }
}

const styles = StyleSheet.create({
  root: {
    flex: 1,
    backgroundColor: 'white'
  },
  introductionModal: {
    flex: 1,
    backgroundColor: 'rgba(0, 0, 0, 0.58)',
    justifyContent: 'center',
    alignItems: 'center'
  },
  predictingModal: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  }
})

export default SelectScreen