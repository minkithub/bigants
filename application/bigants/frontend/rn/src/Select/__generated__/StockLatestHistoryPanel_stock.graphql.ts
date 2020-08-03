/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type StockLatestHistoryPanel_stock = {
    readonly id: string;
    readonly nameKo: string;
    readonly latestHistory: {
        readonly date: string;
        readonly close: number;
        readonly volume: number;
        readonly high: number;
        readonly low: number;
        readonly highLimit: number | null;
        readonly lowLimit: number | null;
        readonly priceChange: number | null;
        readonly priceChangeRate: number | null;
    } | null;
    readonly " $refType": "StockLatestHistoryPanel_stock";
};
export type StockLatestHistoryPanel_stock$data = StockLatestHistoryPanel_stock;
export type StockLatestHistoryPanel_stock$key = {
    readonly " $data"?: StockLatestHistoryPanel_stock$data;
    readonly " $fragmentRefs": FragmentRefs<"StockLatestHistoryPanel_stock">;
};



const node: ReaderFragment = {
  "kind": "Fragment",
  "name": "StockLatestHistoryPanel_stock",
  "type": "Stock",
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
    }
  ]
};
(node as any).hash = '238d5fcbd985a7e811fa6b408bc63a1f';
export default node;
