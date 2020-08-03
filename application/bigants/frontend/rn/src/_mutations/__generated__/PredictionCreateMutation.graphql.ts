/* tslint:disable */
/* eslint-disable */
/* @relayHash d98ea4e7eac0db243c7f3cc302734525 */

import { ConcreteRequest } from "relay-runtime";
export type PredictionCreateInput = {
    clientMutationId?: string | null;
    stockId: string;
    holidays: Array<string>;
};
export type PredictionCreateMutationVariables = {
    input: PredictionCreateInput;
};
export type PredictionCreateMutationResponse = {
    readonly predictionCreate: {
        readonly clientMutationId: string | null;
        readonly result: {
            readonly cursor: string;
            readonly node: {
                readonly id: string;
                readonly requester: {
                    readonly id: string;
                } | null;
            };
        };
    } | null;
};
export type PredictionCreateMutation = {
    readonly response: PredictionCreateMutationResponse;
    readonly variables: PredictionCreateMutationVariables;
};



/*
mutation PredictionCreateMutation(
  $input: PredictionCreateInput!
) {
  predictionCreate(input: $input) {
    clientMutationId
    result {
      cursor
      node {
        id
        requester {
          id
        }
      }
    }
  }
}
*/

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "kind": "LocalArgument",
    "name": "input",
    "type": "PredictionCreateInput!",
    "defaultValue": null
  }
],
v1 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "id",
  "args": null,
  "storageKey": null
},
v2 = [
  {
    "kind": "LinkedField",
    "alias": null,
    "name": "predictionCreate",
    "storageKey": null,
    "args": [
      {
        "kind": "Variable",
        "name": "input",
        "variableName": "input"
      }
    ],
    "concreteType": "PredictionCreatePayload",
    "plural": false,
    "selections": [
      {
        "kind": "ScalarField",
        "alias": null,
        "name": "clientMutationId",
        "args": null,
        "storageKey": null
      },
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "result",
        "storageKey": null,
        "args": null,
        "concreteType": "StockPredictionEdge",
        "plural": false,
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
            "concreteType": "StockPrediction",
            "plural": false,
            "selections": [
              (v1/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "requester",
                "storageKey": null,
                "args": null,
                "concreteType": "User",
                "plural": false,
                "selections": [
                  (v1/*: any*/)
                ]
              }
            ]
          }
        ]
      }
    ]
  }
];
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "PredictionCreateMutation",
    "type": "Mutation",
    "metadata": null,
    "argumentDefinitions": (v0/*: any*/),
    "selections": (v2/*: any*/)
  },
  "operation": {
    "kind": "Operation",
    "name": "PredictionCreateMutation",
    "argumentDefinitions": (v0/*: any*/),
    "selections": (v2/*: any*/)
  },
  "params": {
    "operationKind": "mutation",
    "name": "PredictionCreateMutation",
    "id": null,
    "text": "mutation PredictionCreateMutation(\n  $input: PredictionCreateInput!\n) {\n  predictionCreate(input: $input) {\n    clientMutationId\n    result {\n      cursor\n      node {\n        id\n        requester {\n          id\n        }\n      }\n    }\n  }\n}\n",
    "metadata": {}
  }
};
})();
(node as any).hash = '6b9fc09c41574516b1437e96cb294f43';
export default node;
