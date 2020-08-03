import React from 'react';
import 'react-native-gesture-handler';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import { NavigationProp } from '@react-navigation/native'

import Start from './Start';
import Search from './Search';
import Select from './Select';
import Result from './Result';
import Me from './Me';
import List from './List';
import { RootStackParams } from './route';


const RootStack = createStackNavigator<RootStackParams>();

function App() {
  return (
    <NavigationContainer>
      <RootStack.Navigator>
        <RootStack.Screen name="Home" component={Start} options={{ title: '홈', headerShown: false }} />
        <RootStack.Screen name="Select" component={Select} options={Select.screenOptions} />
        <RootStack.Screen name="PredictionResult" component={Result} options={Result.screenOptions} />
        <RootStack.Screen name="Search" component={Search} options={Search.screenOptions} />
        <RootStack.Screen name="List" component={List} options={{ title: '분석 목록' }} />
      </RootStack.Navigator>
    </NavigationContainer>
  );
}

export default App;