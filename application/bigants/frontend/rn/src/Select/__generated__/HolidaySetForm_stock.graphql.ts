/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type HolidaySetForm_stock = {
    readonly history: ReadonlyArray<{
        readonly d: string;
        readonly c: number;
        readonly v: number;
    }>;
    readonly " $refType": "HolidaySetForm_stock";
};
export type HolidaySetForm_stock$data = HolidaySetForm_stock;
export type HolidaySetForm_stock$key = {
    readonly " $data"?: HolidaySetForm_stock$data;
    readonly " $fragmentRefs": FragmentRefs<"HolidaySetForm_stock">;
};



const node: ReaderFragment = {
  "kind": "Fragment",
  "name": "HolidaySetForm_stock",
  "type": "Stock",
  "metadata": null,
  "argumentDefinitions": [],
  "selections": [
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
};
(node as any).hash = '3c5750df671a4949230e972d183a9201';
export default node;
