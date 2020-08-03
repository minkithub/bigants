export function numf(val: number): string {
  return Number(val).toLocaleString();
}

/**
 * 2006-01-02 형식으로 날짜를 포매팅한다.
 * @param val 날짜
 */
export function datef(val: Date): string {
  return val.toISOString().slice(0, 10);
}