/* tslint:disable */
/* eslint-disable */
/* @relayHash 88b6e68d924513a2fdfe6b0ca773776a */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type SearchListQueryVariables = {
    count: number;
    cursor?: string | null;
    q?: string | null;
};
export type SearchListQueryResponse = {
    readonly " $fragmentRefs": FragmentRefs<"SearchList_query">;
};
export type SearchListQuery = {
    readonly response: SearchListQueryResponse;
    readonly variables: SearchListQueryVariables;
};



/*
query SearchListQuery(
  $count: Int!
  $cursor: String
  $q: String
) {
  ...SearchList_query_XhAmI
}

fragment SearchList_query_XhAmI on Query {
  stocks(first: $count, after: $cursor, q: $q) {
    edges {
      cursor
      node {
        id
        code
        nameKo
        __typename
      }
    }
    pageInfo {
      endCursor
      hasNextPage
    }
  }
}
*/

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "kind": "LocalArgument",
    "name": "count",
    "type": "Int!",
    "defaultValue": null
  },
  {
    "kind": "LocalArgument",
    "name": "cursor",
    "type": "String",
    "defaultValue": null
  },
  {
    "kind": "LocalArgument",
    "name": "q",
    "type": "String",
    "defaultValue": null
  }
],
v1 = {
  "kind": "Variable",
  "name": "q",
  "variableName": "q"
},
v2 = [
  {
    "kind": "Variable",
    "name": "after",
    "variableName": "cursor"
  },
  {
    "kind": "Variable",
    "name": "first",
    "variableName": "count"
  },
  (v1/*: any*/)
];
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "SearchListQuery",
    "type": "Query",
    "metadata": null,
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "FragmentSpread",
        "name": "SearchList_query",
        "args": [
          {
            "kind": "Variable",
            "name": "count",
            "variableName": "count"
          },
          {
            "kind": "Variable",
            "name": "cursor",
            "variableName": "cursor"
          },
          (v1/*: any*/)
        ]
      }
    ]
  },
  "operation": {
    "kind": "Operation",
    "name": "SearchListQuery",
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "stocks",
        "storageKey": null,
        "args": (v2/*: any*/),
        "concreteType": "StockConnection",
        "plural": false,
        "selections": [
          {
            "kind": "LinkedField",
            "alias": null,
            "name": "edges",
            "storageKey": null,
            "args": null,
            "concreteType": "StockEdge",
            "plural": true,
            "selections": [
              {
                "kind": "ScalarField",
                "alias": null,
                "name": "cursor",
                "args": null,
                "storageKey": null
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "node",
                "storageKey": null,
                "args": null,
                "concreteType": "Stock",
                "plural": false,
                "selections": [
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "id",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "code",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "nameKo",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "__typename",
                    "args": null,
                    "storageKey": null
                  }
                ]
              }
            ]
          },
          {
            "kind": "LinkedField",
            "alias": null,
            "name": "pageInfo",
            "storageKey": null,
            "args": null,
            "concreteType": "PageInfo",
            "plural": false,
            "selections": [
              {
                "kind": "ScalarField",
                "alias": null,
                "name": "endCursor",
                "args": null,
                "storageKey": null
              },
              {
                "kind": "ScalarField",
                "alias": null,
                "name": "hasNextPage",
                "args": null,
                "storageKey": null
              }
            ]
          }
        ]
      },
      {
        "kind": "LinkedHandle",
        "alias": null,
        "name": "stocks",
        "args": (v2/*: any*/),
        "handle": "connection",
        "key": "SearchList_stocks",
        "filters": [
          "q"
        ]
      }
    ]
  },
  "params": {
    "operationKind": "query",
    "name": "SearchListQuery",
    "id": null,
    "text": "query SearchListQuery(\n  $count: Int!\n  $cursor: String\n  $q: String\n) {\n  ...SearchList_query_XhAmI\n}\n\nfragment SearchList_query_XhAmI on Query {\n  stocks(first: $count, after: $cursor, q: $q) {\n    edges {\n      cursor\n      node {\n        id\n        code\n        nameKo\n        __typename\n      }\n    }\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n  }\n}\n",
    "metadata": {}
  }
};
})();
(node as any).hash = 'f726bf0d6d5e0ac46cba24842f676f73';
export default node;
