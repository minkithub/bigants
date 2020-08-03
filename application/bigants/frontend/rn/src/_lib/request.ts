// import { getCurrentPositionSync } from './geolocation';
import { getToken } from './auth';
// import { tryGetFCMToken, getIID } from './firebase';


let seq = 0;

export async function request(url: string, init: RequestInit) {

  const id = seq++;

  let { headers, ...rest } = init;

  const started = new Date().getTime();
  headers = { ...headers, ...(await getHeaders()) };

  console.log(`[Fetch:${id}] Started T+0`, url, headers, init.body);
  const res = await fetch(url, { ...rest, headers });

  console.log(`[Fetch:${id}] Done T+${new Date().getTime() - started}`, );
  if (200 <= res.status && res.status < 400) {
    return {res, id};
  }
  const errorMessage = `HTTP ${res.status}: ${await res.text()} (${res.statusText})`;
  console.log(`[Fetch:${id}] Result`, errorMessage);
  throw new Error(errorMessage);
}

export async function requestJSON(url: string, init: RequestInit) {
  const { headers, ...rest } = init;

  const {res, id} = await request(url, {
    ...rest,
    headers: { ...headers, 'content-type': 'application/json' },
  });
  const data = await res.text();
  console.log(`Fetch:${id}] Result`, data);
  return JSON.parse(data);
}

async function getHeaders(): Promise<any> {
  const headers: any = {};

  const [token, ] = await Promise.all([
    // getCurrentPositionSync(),
    getToken().catch(() => null),
    // getIID().catch(() => null),
    // tryGetFCMToken().catch(() => null),
  ]);

  if (token) {
    headers.Authorization = token;
  }

  return headers;
}

