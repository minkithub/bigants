import React from 'react';
import { View, Text, TouchableOpacity, Alert } from 'react-native';


function ResultScreen({ navigation }) {
  return (
    <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center', backgroundColor: 'white' }}>
      <Text style={{ fontSize: 16, color: '#535358', textAlign: 'center' }}>내 목록을 만드려면 로그인이 필요합니다..</Text>

      <TouchableOpacity onPress={() => Alert.alert('카카오로 로그인')} style={{ width: 375, height: 80, justifyContent: 'center', alignItems: 'center' }} activeOpacity={1}>
        <View style={{ paddingHorizontal: 24, height: 44, borderRadius: 8, justifyContent: 'center', alignItems: 'center', backgroundColor: '#FFCD00' }}>
          <Text style={{ color: '#7C4B1D', fontSize: 16, fontWeight: 'bold' }}>카카오톡 로그인</Text>
        </View>
      </TouchableOpacity>


    </View>
  );
}

export default ResultScreen