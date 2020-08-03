import React from 'react';
import { View, Text, TouchableOpacity, StyleSheet, SafeAreaView, Image } from 'react-native';
import { QueryRenderer, graphql } from 'react-relay';
import { relayEnvironment } from '../_lib/relay';
import { StartQuery } from './__generated__/StartQuery.graphql';
import { theme } from '../theme';


const searchIcon = require('../assets/searchIcon.png');

function HomeScreen({ navigation }) {
  return (
    <QueryRenderer<StartQuery>
      environment={relayEnvironment}
      query={graphql`
        query StartQuery {
          version
        }
      `}
      variables={{}}
      render={({ error, props, retry }) => {
        if (error) {
          <View style={{ flex: 1 }}>
            <Text>앱에 오류가 났습니다</Text>
            <Text>다시 시작해 주세요</Text>
          </View>
        }

        return (
          <View style={styles.root}>
            <View style={styles.top}>
              <View style={styles.header}>
                <View style={styles.headerTitleBox}>
                  <Text style={theme.xlwText}>쉽고 빠르게 원하는 종목의</Text>
                  <Text style={theme.xlwText}><Text style={{ fontWeight: 'bold' }}>주식가격을 예측</Text>해 보세요!</Text>
                </View>
              </View>
              <TouchableOpacity onPress={() => navigation.navigate('Search')} style={styles.searchButton} activeOpacity={1}>
                <Image source={searchIcon} style={styles.searchButtonIcon} />
                <Text style={theme.mnText}>주식명 또는 주식코드로 검색</Text>
              </TouchableOpacity>
            </View>
            <View style={styles.middle}>
              {/* <Text style={{ fontSize: 36, fontWeight: 'bold' }}> 1 2 0 0 개</Text>
              <Text style={{ paddingTop: 12, fontSize: 18 }}>다른 사용자의 분석이 있습니다</Text>
              <Text style={{ fontSize: 18, paddingTop: 24, fontWeight: 'bold', color: '#0084F4', }}>분석보기</Text> */}
            </View>
            <View style={styles.bottom}>
              {/* TODO : 이미지로 바꾸기*/}
              <Text style={theme.lbblueText}>Big</Text><Text style={theme.lbbText}>Ants. {props && props.version}</Text>
            </View>
          </View>
        )
      }}
    />
  );
}

const styles = StyleSheet.create({
  root: {
    flex: 1,
    flexDirection: 'column',
    backgroundColor: 'white',
  },
  top: {
    flexDirection: 'column',
    flexGrow: 2,
    paddingBottom: theme.MU.V5,
  },
  middle: {
    flexGrow: 3,
    paddingTop: 20,
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
  },
  bottom: {
    flexGrow: 1,
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center'
  },

  // top
  header: {
    flexGrow: 1,
    backgroundColor: theme.colors.primary.blue,
    justifyContent: 'flex-end',
    paddingBottom: theme.MU.V6,
    borderBottomLeftRadius: 44,
  },
  headerTitleBox: {
    paddingLeft: theme.MU.H2,
  },
  searchButton: {
    position: 'absolute',
    bottom: 20,
    flexDirection: 'row',
    alignItems: 'center',
    height: 50,
    marginHorizontal: theme.MU.H2,
    backgroundColor: theme.colors.primary.white,
    borderRadius: 8,
    paddingLeft: theme.MU.H1,
    ...theme.boxShadow,
    width: '90%'
  },
  searchButtonIcon: {
    width: 28,
    height: 28,
    marginRight: theme.MU.H1,
  },
})

export default HomeScreen