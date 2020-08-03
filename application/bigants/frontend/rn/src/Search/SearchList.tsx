import React from 'react';
import { createPaginationContainer, graphql, RelayPaginationProp } from 'react-relay';

import { SearchList_query } from './__generated__/SearchList_query.graphql'
import { FlatList, View, Text, ActivityIndicator, TextInput, StyleSheet, SafeAreaView, TouchableOpacity } from 'react-native';
import { debounce } from 'lodash';
import { theme } from '../theme';


export const SearchList = createPaginationContainer<{
  query?: SearchList_query,
  relay: RelayPaginationProp,
  onPressItem: (code: string) => void,
}>(
  ({ query, relay, onPressItem }) => {
    const [reloading, setReloading] = React.useState(false);
    const [loadingMore, setLoadingMore] = React.useState(false);

    // TODO: 로컬 로딩상태 최적의 답을 찾을 것

    return (
      <FlatList
        keyboardDismissMode='on-drag'
        style={{ backgroundColor: 'white' }}
        contentContainerStyle={{ paddingBottom: theme.MU.V5 }}
        stickyHeaderIndices={[0]}
        data={query ? query.stocks.edges : []}
        onRefresh={() => {
          setReloading(true);
          relay.refetchConnection(20, () => setReloading(false))
        }}
        refreshing={reloading}
        onEndReached={() => {
          if (query && !relay.isLoading() && relay.hasMore()) {
            setLoadingMore(true)
            relay.loadMore(20, () => setLoadingMore(false))
          }
        }}
        keyboardShouldPersistTaps='handled'
        keyExtractor={item => item.cursor}
        ListHeaderComponent={(
          <View style={styles.header}>
            <View style={styles.searchBox}>
              <TextInput
                style={[theme.mnText, { paddingLeft: theme.MU.V2 }]}
                placeholder='주식명 또는 주식코드를 입력해주세요.'
                onChangeText={debounce(
                  q => relay.refetchConnection(20, null, { q }), 150
                )}
                autoFocus
              />
            </View>
          </View>
        )}
        renderItem={({ item }) => (
          <TouchableOpacity onPress={() => onPressItem(item.node.id)} style={styles.contentItem}>
            <Text style={theme.mnText}>{item.node.nameKo} ({item.node.code})</Text>
          </TouchableOpacity>
        )}
        ListFooterComponent={
          (!query || loadingMore) ? (
            <View style={styles.contentItem}>
              <ActivityIndicator />
            </View>
          ) : null
        }
      />
    );
  },
  {
    query: graphql`
      fragment SearchList_query on Query @argumentDefinitions(
        count: {type: "Int", defaultValue: 20 },
        cursor: {type: "String" },
        q: {type: "String" }
      ) {
        stocks(first: $count, after: $cursor, q: $q)  @connection(key: "SearchList_stocks") {
          edges {
            cursor
            node {
              id
              code
              nameKo
            }
          }
        }
      }
    `,
  },
  {
    getVariables: (props, info) => ({ ...info }),
    query: graphql`
      query SearchListQuery($count: Int! $cursor: String $q: String) {
        ...SearchList_query @arguments(count: $count, cursor: $cursor, q: $q)
      }
    `
  },
)

const styles = StyleSheet.create({
  header: {
    flexDirection: 'row',
    alignItems: 'center',
    height: 52,
    paddingHorizontal: theme.MU.H2,
    backgroundColor: theme.colors.primary.white,
    marginBottom: 4,
  },
  contentItem: {
    justifyContent: 'center',
    height: 56,
    marginHorizontal: theme.MU.H2,
    paddingLeft: theme.MU.H1,
    borderBottomColor: '#EFEFEF',
    borderBottomWidth: 1
  },

  //header
  searchBox: {
    flexGrow: 1,
    flexDirection: 'column',
    justifyContent: 'center',
    backgroundColor: '#EFEFEF',

    borderRadius: 12,
    height: 48
  },
  clearButton: {
    width: 28,
    height: 28,
  }
})