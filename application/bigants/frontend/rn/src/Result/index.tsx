import React from 'react';
import { View, Text, TouchableOpacity, SafeAreaView, ActivityIndicator, Image, Platform, Dimensions } from 'react-native';
import { QueryRenderer, graphql } from 'react-relay';
import { relayEnvironment } from '../_lib/relay';
import { ResultQuery } from './__generated__/ResultQuery.graphql';
import { RouteProp } from '@react-navigation/native';
import { RootStackParams } from '../route';
import { StackNavigationProp } from '@react-navigation/stack';
import { PredictionResultHeader } from './PredictionResultHeader';
import { theme } from '../theme';



const homeIcon = require('../assets/homeIcon.png');
const screenWidth = Dimensions.get('screen').width;

class ResultScreen extends React.Component<{
  route: RouteProp<RootStackParams, 'PredictionResult'>,
  navigation: StackNavigationProp<RootStackParams, 'PredictionResult'>,
}> {
  static screenOptions = {
    title: '주가 예측결과',
    headerStyle: { borderBottomColor: 'transparent' },
    headerTitleAlign: 'left',
    headerTitleStyle: { ...theme.lnText },
    headerBackTitle: '',
    headerBackTitleVisible: false,
    headerLeftContainerStyle: { marginLeft: theme.MU.H1 },
    headerBackground: () => <View style={{ backgroundColor: 'white', height: '100%' }} />,
    headerRight: ({ navigation }) => (
      <TouchableOpacity onPress={() => navigation.popToTop()} style={{ width: 26, height: 26, justifyContent: 'center', alignItems: 'center', marginRight: theme.MU.H2 }}>
        <Image source={homeIcon} style={{ width: 28, height: 28 }} />
      </TouchableOpacity>
    )
  }
  render() {
    const { navigation, route } = this.props;

    return (
      <QueryRenderer<ResultQuery>
        environment={relayEnvironment}
        query={graphql`
        query ResultQuery($stockPredictionId: ID) {
          stockPrediction: node(id: $stockPredictionId) {
            ...PredictionResultHeader_prediction
          }
        }
      `}
        variables={{
          stockPredictionId: route.params.stockPredictionId,
        }}
        render={({ error, props, retry }) => {
          if (error) {
            return <Text>{error.message}</Text>
          }
          if (!props) {
            return <ActivityIndicator style={{ flex: 1 }} />
          }
          return (
            <SafeAreaView style={{ flex: 1, backgroundColor: 'white' }}>
              <View style={{ flexGrow: 1, justifyContent: 'space-between' }}>
                {/* card */}
                <PredictionResultHeader prediction={props.stockPrediction} />
                <TouchableOpacity onPress={() => navigation.pop()} style={{ justifyContent: 'center', alignItems: 'center', position: 'absolute', zIndex: 2, bottom: 40, left: screenWidth / 2 }} activeOpacity={1}>
                  <View style={{ paddingHorizontal: 24, height: 44, borderRadius: 8, justifyContent: 'center', alignItems: 'center', backgroundColor: '#0084F4' }}>
                    <Text style={{ color: '#fff', fontSize: 16, fontWeight: 'bold' }}>다시하기</Text>
                  </View>
                </TouchableOpacity>

              </View>
            </SafeAreaView>
          )
        }}
      />
    );
  }
}

export default ResultScreen