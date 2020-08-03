/* tslint:disable */
/* eslint-disable */
/* @relayHash ed220602f369ba9a2bfaf3c560448913 */

import { ConcreteRequest } from "relay-runtime";
export type StartQueryVariables = {};
export type StartQueryResponse = {
    readonly version: string;
};
export type StartQuery = {
    readonly response: StartQueryResponse;
    readonly variables: StartQueryVariables;
};



/*
query StartQuery {
  version
}
*/

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "version",
    "args": null,
    "storageKey": null
  }
];
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "StartQuery",
    "type": "Query",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": (v0/*: any*/)
  },
  "operation": {
    "kind": "Operation",
    "name": "StartQuery",
    "argumentDefinitions": [],
    "selections": (v0/*: any*/)
  },
  "params": {
    "operationKind": "query",
    "name": "StartQuery",
    "id": null,
    "text": "query StartQuery {\n  version\n}\n",
    "metadata": {}
  }
};
})();
(node as any).hash = '8372e9ddbe4c53893492f7313ab3a56c';
export default node;
