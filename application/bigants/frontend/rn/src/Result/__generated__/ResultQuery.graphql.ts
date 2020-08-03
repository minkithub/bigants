/* tslint:disable */
/* eslint-disable */
/* @relayHash c727322a9a5a559588911038cfb050aa */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type ResultQueryVariables = {
    stockPredictionId?: string | null;
};
export type ResultQueryResponse = {
    readonly stockPrediction: {
        readonly " $fragmentRefs": FragmentRefs<"PredictionResultHeader_prediction">;
    } | null;
};
export type ResultQuery = {
    readonly response: ResultQueryResponse;
    readonly variables: ResultQueryVariables;
};



/*
query ResultQuery(
  $stockPredictionId: ID
) {
  stockPrediction: node(id: $stockPredictionId) {
    __typename
    ...PredictionResultHeader_prediction
    id
  }
}

fragment PredictionResultHeader_prediction on StockPrediction {
  id
  stock {
    nameKo
    latestHistory {
      close
    }
    history(count: 25) {
      d: date
      c: close
    }
    id
  }
  averageIncome
  accuracy
  dailyPredictions {
    date
    expectedPrice
    expectedIncome
  }
}
*/

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "kind": "LocalArgument",
    "name": "stockPredictionId",
    "type": "ID",
    "defaultValue": null
  }
],
v1 = [
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "stockPredictionId"
  }
],
v2 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "id",
  "args": null,
  "storageKey": null
};
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "ResultQuery",
    "type": "Query",
    "metadata": null,
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": "stockPrediction",
        "name": "node",
        "storageKey": null,
        "args": (v1/*: any*/),
        "concreteType": null,
        "plural": false,
        "selections": [
          {
            "kind": "FragmentSpread",
            "name": "PredictionResultHeader_prediction",
            "args": null
          }
        ]
      }
    ]
  },
  "operation": {
    "kind": "Operation",
    "name": "ResultQuery",
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": "stockPrediction",
        "name": "node",
        "storageKey": null,
        "args": (v1/*: any*/),
        "concreteType": null,
        "plural": false,
        "selections": [
          {
            "kind": "ScalarField",
            "alias": null,
            "name": "__typename",
            "args": null,
            "storageKey": null
          },
          (v2/*: any*/),
          {
            "kind": "InlineFragment",
            "type": "StockPrediction",
            "selections": [
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "stock",
                "storageKey": null,
                "args": null,
                "concreteType": "Stock",
                "plural": false,
                "selections": [
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "nameKo",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "latestHistory",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "StockHistory",
                    "plural": false,
                    "selections": [
                      {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "close",
                        "args": null,
                        "storageKey": null
                      }
                    ]
                  },
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "history",
                    "storageKey": "history(count:25)",
                    "args": [
                      {
                        "kind": "Literal",
                        "name": "count",
                        "value": 25
                      }
                    ],
                    "concreteType": "StockHistory",
                    "plural": true,
                    "selections": [
                      {
                        "kind": "ScalarField",
                        "alias": "d",
                        "name": "date",
                        "args": null,
                        "storageKey": null
                      },
                      {
                        "kind": "ScalarField",
                        "alias": "c",
                        "name": "close",
                        "args": null,
                        "storageKey": null
                      }
                    ]
                  },
                  (v2/*: any*/)
                ]
              },
              {
                "kind": "ScalarField",
                "alias": null,
                "name": "averageIncome",
                "args": null,
                "storageKey": null
              },
              {
                "kind": "ScalarField",
                "alias": null,
                "name": "accuracy",
                "args": null,
                "storageKey": null
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "dailyPredictions",
                "storageKey": null,
                "args": null,
                "concreteType": "StockDailyPrediction",
                "plural": true,
                "selections": [
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "date",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "expectedPrice",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "expectedIncome",
                    "args": null,
                    "storageKey": null
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  },
  "params": {
    "operationKind": "query",
    "name": "ResultQuery",
    "id": null,
    "text": "query ResultQuery(\n  $stockPredictionId: ID\n) {\n  stockPrediction: node(id: $stockPredictionId) {\n    __typename\n    ...PredictionResultHeader_prediction\n    id\n  }\n}\n\nfragment PredictionResultHeader_prediction on StockPrediction {\n  id\n  stock {\n    nameKo\n    latestHistory {\n      close\n    }\n    history(count: 25) {\n      d: date\n      c: close\n    }\n    id\n  }\n  averageIncome\n  accuracy\n  dailyPredictions {\n    date\n    expectedPrice\n    expectedIncome\n  }\n}\n",
    "metadata": {}
  }
};
})();
(node as any).hash = '5a775e1b07a31df6c8508bdf02bfc6df';
export default node;
