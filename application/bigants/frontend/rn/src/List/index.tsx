import React from 'react';
import { View, Text, TouchableOpacity, SafeAreaView } from 'react-native';


function ResultScreen({ navigation }) {
  return (
    <SafeAreaView style={{ flex: 1, justifyContent: 'space-between', backgroundColor: '#e3e3e3' }}>
      <View style={{ flexGrow: 1 }}>
        <View style={{ marginHorizontal: 16, marginTop: 20, padding: 20, backgroundColor: 'white', borderRadius: 8, }}>
          <Text style={{ fontSize: 24, fontWeight: 'bold', }}>SK이노베이션</Text>
          <View style={{ flexDirection: 'row', paddingTop: 16 }}>
            <Text style={{ paddingTop: 8, fontSize: 16, color: '#353637' }}>예측평균수익률</Text>
            <Text style={{ paddingLeft: 20, paddingTop: 8, fontSize: 16, fontWeight: 'bold', color: '#FC0000' }}>31.26%</Text>
          </View>
          <View style={{ flexDirection: 'row' }}>
            <Text style={{ paddingTop: 8, fontSize: 16, color: '#353637' }}>모델 예측 정확도</Text>
            <Text style={{ paddingLeft: 20, paddingTop: 8, fontSize: 16, fontWeight: 'bold', color: '#FC0000' }}>94.6%</Text>
          </View>
        </View>

        <View style={{ marginHorizontal: 16, marginTop: 20, padding: 20, backgroundColor: 'white', borderRadius: 8, }}>
          <Text style={{ fontSize: 24, fontWeight: 'bold', }}>SK하이닉스</Text>
          <View style={{ flexDirection: 'row', paddingTop: 16 }}>
            <Text style={{ paddingTop: 8, fontSize: 16, color: '#353637' }}>예측평균수익률</Text>
            <Text style={{ paddingLeft: 20, paddingTop: 8, fontSize: 16, fontWeight: 'bold', color: '#FC0000' }}>21.26%</Text>
          </View>
          <View style={{ flexDirection: 'row' }}>
            <Text style={{ paddingTop: 8, fontSize: 16, color: '#353637' }}>모델 예측 정확도</Text>
            <Text style={{ paddingLeft: 20, paddingTop: 8, fontSize: 16, fontWeight: 'bold', color: '#FC0000' }}>95.6%</Text>
          </View>
        </View>
      </View>
      <TouchableOpacity onPress={() => navigation.navigate('User')} style={{ width: 375, height: 80, justifyContent: 'center', alignItems: 'center' }} activeOpacity={1}>
        <View style={{ paddingHorizontal: 24, height: 44, borderRadius: 8, justifyContent: 'center', alignItems: 'center', backgroundColor: '#0084F4' }}>
          <Text style={{ color: '#fff', fontSize: 16, fontWeight: 'bold' }}>내 분석</Text>
        </View>
      </TouchableOpacity>

    </SafeAreaView>
  );
}

export default ResultScreen