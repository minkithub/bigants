import React from 'react';


export type TxProps<T> = {
  render(tx: {
    /**
     * 현재 트랜잭션이 실행중인 지 여부를 반환
     */
    inProgress: boolean,
    /**
     * 주어진 함수를 트랜잭션으로 감싸기
     */
    as: (func: T) => T,

    /**
     * 트랜잭션 취소
     */
  }): React.ReactNode;
}

type State = {
  inProgress: boolean;
}
/**
 * `Tx` 컴포넌트는 `render` 함수를 `prop`으로 받는다.
 *
 * `render` 함수는 현재 트랜잭션의 진행 여부(`inProgress`)와 특정 함수를 트랜잭션 래퍼로 감싸는 `as`를 인자로 주입하여 호출된다.
 */
export class Tx<T extends (...args: any[]) => Promise<any>> extends React.Component<TxProps<T>, State> {
  state: State = { inProgress: false };

  get inProgress() { return this.state.inProgress; }

  as = (func: T) => {
    return (async (...args: any[]) => {
      if (this.state.inProgress) { return; } // noop
      await setStateAsync(this, { inProgress: true });
      try {
        return await func(...args);
      } catch (err) {
        throw err;
      }
      finally {
        await setStateAsync(this, { inProgress: false }); // ! Finally 문법 주의
      }
    }) as T
  }

  render() {
    return this.props.render(this);
  }
}

export function setStateAsync<P, S, K extends keyof S>(
  c: React.Component<P, S>,
  state: ((prevState: Readonly<S>, props: Readonly<P>) => (Pick<S, K> | S | null)) | (Pick<S, K> | S | null),
) {
  return new Promise(resolve => c.setState(state, resolve));
}
