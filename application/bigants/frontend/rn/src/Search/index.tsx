import React from 'react';
import { View } from 'react-native';
import { QueryRenderer, graphql } from 'react-relay';
import { relayEnvironment } from '../_lib/relay';
import { SearchQuery } from './__generated__/SearchQuery.graphql';
import { SearchList } from './SearchList';
import { RouteProp } from '@react-navigation/native'
import { RootStackParams } from '../route';
import { StackNavigationProp } from '@react-navigation/stack';
import { theme } from '../theme';


class SearchScreen extends React.Component<{
  route: RouteProp<RootStackParams, 'Search'>,
  navigation: StackNavigationProp<RootStackParams, 'Search'>,
}> {
  static screenOptions = {
    title: '종목 검색',
    headerStyle: { borderBottomColor: 'transparent' },
    headerTitleAlign: 'left',
    headerTitleStyle: { ...theme.lnText },
    headerBackTitle: '',
    headerBackTitleVisible: false,
    headerLeftContainerStyle: { marginLeft: theme.MU.H1 },
    headerBackground: () => <View style={{ backgroundColor: 'white', height: '100%' }} />,
  }
  render() {
    const { navigation } = this.props;
    return (
      <QueryRenderer<SearchQuery>
        environment={relayEnvironment}
        query={graphql`
          query SearchQuery {
            ...SearchList_query
          }
      `}
        variables={{}}
        render={({ error, props, retry }) => {
          if (error) {
            throw error;
          }
          return (
            <SearchList
              query={props}
              onPressItem={stockId => navigation.navigate('Select', { stockId })}
            />
          )
        }}
      />
    );
  }
}

export default SearchScreen