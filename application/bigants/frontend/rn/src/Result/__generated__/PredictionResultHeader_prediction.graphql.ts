/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type PredictionResultHeader_prediction = {
    readonly id: string;
    readonly stock: {
        readonly nameKo: string;
        readonly latestHistory: {
            readonly close: number;
        } | null;
        readonly history: ReadonlyArray<{
            readonly d: string;
            readonly c: number;
        }>;
    };
    readonly averageIncome: number;
    readonly accuracy: number;
    readonly dailyPredictions: ReadonlyArray<{
        readonly date: string;
        readonly expectedPrice: number;
        readonly expectedIncome: number;
    } | null>;
    readonly " $refType": "PredictionResultHeader_prediction";
};
export type PredictionResultHeader_prediction$data = PredictionResultHeader_prediction;
export type PredictionResultHeader_prediction$key = {
    readonly " $data"?: PredictionResultHeader_prediction$data;
    readonly " $fragmentRefs": FragmentRefs<"PredictionResultHeader_prediction">;
};



const node: ReaderFragment = {
  "kind": "Fragment",
  "name": "PredictionResultHeader_prediction",
  "type": "StockPrediction",
  "metadata": null,
  "argumentDefinitions": [],
  "selections": [
    {
      "kind": "ScalarField",
      "alias": null,
      "name": "id",
      "args": null,
      "storageKey": null
    },
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
        }
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
};
(node as any).hash = '2e507445009e57f319cd223253a10fd5';
export default node;
