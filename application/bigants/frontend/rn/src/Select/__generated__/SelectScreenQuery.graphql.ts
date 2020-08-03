/* tslint:disable */
/* eslint-disable */
/* @relayHash abb773286d55184515e4b6d1ad22b16d */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type SelectScreenQueryVariables = {
    stockId?: string | null;
};
export type SelectScreenQueryResponse = {
    readonly stock: {
        readonly " $fragmentRefs": FragmentRefs<"StockLatestHistoryPanel_stock" | "HolidaySetForm_stock">;
    } | null;
};
export type SelectScreenQuery = {
    readonly response: SelectScreenQueryResponse;
    readonly variables: SelectScreenQueryVariables;
};



/*
query SelectScreenQuery(
  $stockId: ID
) {
  stock: node(id: $stockId) {
    __typename
    ...StockLatestHistoryPanel_stock
    ...HolidaySetForm_stock
    id
  }
}

fragment HolidaySetForm_stock on Stock {
  history {
    d: date
    c: close
    v: volume
  }
}

fragment StockLatestHistoryPanel_stock on Stock {
  id
  nameKo
  latestHistory {
    date
    close
    volume
    high
    low
    highLimit
    lowLimit
    priceChange
    priceChangeRate
  }
}
*/

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "kind": "LocalArgument",
    "name": "stockId",
    "type": "ID",
    "defaultValue": null
  }
],
v1 = [
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "stockId"
  }
];
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "SelectScreenQuery",
    "type": "Query",
    "metadata": null,
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": "stock",
        "name": "node",
        "storageKey": null,
        "args": (v1/*: any*/),
        "concreteType": null,
        "plural": false,
        "selections": [
          {
            "kind": "FragmentSpread",
            "name": "StockLatestHistoryPanel_stock",
            "args": null
          },
          {
            "kind": "FragmentSpread",
            "name": "HolidaySetForm_stock",
            "args": null
          }
        ]
      }
    ]
  },
  "operation": {
    "kind": "Operation",
    "name": "SelectScreenQuery",
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": "stock",
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
          {
            "kind": "ScalarField",
            "alias": null,
            "name": "id",
            "args": null,
            "storageKey": null
          },
          {
            "kind": "InlineFragment",
            "type": "Stock",
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
                    "name": "date",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "close",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "volume",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "high",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "low",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "highLimit",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "lowLimit",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "priceChange",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "priceChangeRate",
                    "args": null,
                    "storageKey": null
                  }
                ]
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "history",
                "storageKey": null,
                "args": null,
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
                  },
                  {
                    "kind": "ScalarField",
                    "alias": "v",
                    "name": "volume",
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
    "name": "SelectScreenQuery",
    "id": null,
    "text": "query SelectScreenQuery(\n  $stockId: ID\n) {\n  stock: node(id: $stockId) {\n    __typename\n    ...StockLatestHistoryPanel_stock\n    ...HolidaySetForm_stock\n    id\n  }\n}\n\nfragment HolidaySetForm_stock on Stock {\n  history {\n    d: date\n    c: close\n    v: volume\n  }\n}\n\nfragment StockLatestHistoryPanel_stock on Stock {\n  id\n  nameKo\n  latestHistory {\n    date\n    close\n    volume\n    high\n    low\n    highLimit\n    lowLimit\n    priceChange\n    priceChangeRate\n  }\n}\n",
    "metadata": {}
  }
};
})();
(node as any).hash = 'd2088d6fc5822dcafe1f525949d8fa0c';
export default node;
