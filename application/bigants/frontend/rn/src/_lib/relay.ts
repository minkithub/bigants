import { Environment, Network, RecordSource, Store, Observable, RequestParameters, Variables, SubscribeFunction, GraphQLResponse, QueryResponseCache, FetchFunction } from 'relay-runtime';

import env from '../env';

import { requestJSON } from './request';


export function initializeRelayEnvironment() {

  const source = new RecordSource();
  const store = new Store(source);


  /**
   * 30초간 데이터를 보관하는 응답캐시 (reset)
   */
  const cache = new QueryResponseCache({ size: 250, ttl: 30 * 1000 });

  const fetchQuery: FetchFunction = async (operation, variables, cacheConfig): Promise<GraphQLResponse> => {
    const queryID = operation.text;
    const isMutation = operation.operationKind === 'mutation';
    const isQuery = operation.operationKind === 'query';
    const forceFetch = cacheConfig && cacheConfig.force;

    // Try to get data from cache on queries
    const fromCache = queryID ? cache.get(queryID, variables) : null;
    if (
      isQuery &&
      fromCache !== null &&
      !forceFetch
    ) {
      return fromCache;
    }
    try {
      const json = await requestJSON(env.GRAPHQL_ENDPOINT, {
        method: 'POST',
        body: JSON.stringify({ query: operation.text, variables }),
      });

      if (isQuery && json && queryID) {
        cache.set(queryID, variables, json);
      }
      // Clear cache on mutations
      if (isMutation) {
        cache.clear();
      }
      return json;
    } catch (e) {
      return { errors: [new Error('네트워크 요청이 실패했습니다. ' + e.message)], data: undefined };
    }
  };
  const network = Network.create(fetchQuery);

  return new Environment({ store: store, network });
}

export let relayEnvironment: Environment = initializeRelayEnvironment();

export function resetRelayEnvironment() {
  relayEnvironment = initializeRelayEnvironment();
}

declare const global: any;
global.__RELAY_ENVIRONMENT__ = relayEnvironment;

export function mustGetRelayEnvironment() {
  if (relayEnvironment instanceof Environment) {
    return relayEnvironment;
  }
  throw new Error('Relay environment not initialized');
}
