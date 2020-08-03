import AsyncStorage from '@react-native-community/async-storage';

import { Buffer } from 'buffer';
import env from '../env';


const BearerKey = 'BEARER';
const BasicKey = 'BASIC';


export async function getToken() {
  const [basicToken, bearerToken] = await Promise.all([getBasicToken(), getBearerToken()]);
  return basicToken || bearerToken; // TODO: refresh 구현해야 의미 있을듯
}

function getBasicToken() { return AsyncStorage.getItem(BasicKey); }
function getBearerToken() { return AsyncStorage.getItem(BearerKey); }
function setBasicToken(token: string) { return AsyncStorage.setItem(BasicKey, token); }
function setBearerToken(token: string) { return AsyncStorage.setItem(BearerKey, token); }

export async function removeToken() {
  await Promise.all([ AsyncStorage.removeItem(BasicKey), AsyncStorage.removeItem(BearerKey) ]);
  return;
}

export function encodeBasicToken(username: string, password: string) {
  return `Basic ${Buffer.from(`${username}:${password}`).toString('base64')}`;
}

export async function login(username: string, password: string) {
  const basicToken = encodeBasicToken(username, password);
  // const bearerToken = await fetchBearerToken(basicToken);

  await Promise.all([ setBasicToken(basicToken) ]);
}

declare const global: any;
global.login = login;

export async function logout() {
  await removeToken();
}
global.logout = logout;

export async function signUp() {

  const res = await fetch(env.SIGNUP_ENDPOINT, {
    method: 'POST',
  });

  if (res.status !== 201) {
    console.log(env.SIGNUP_ENDPOINT, 'POST', await res.text());
    throw new Error(`회원가입에 실패했습니다. ${res.status}`);
  } else {
    console.log(env.SIGNUP_ENDPOINT, 'POST', res);
  }

  const basicToken = await res.text();

  await setBasicToken(basicToken);

}
