import { commitMutation, Environment, graphql, ConnectionHandler } from 'relay-runtime';
import { PredictionCreateMutation, PredictionCreateMutationResponse, PredictionCreateInput } from './__generated__/PredictionCreateMutation.graphql';

export function predictionCreate(env: Environment, input: PredictionCreateInput) {
  return new Promise<PredictionCreateMutationResponse>(
    resolve => commitMutation<PredictionCreateMutation>(env, {
      mutation: graphql`
        mutation PredictionCreateMutation($input: PredictionCreateInput!) {
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
      `,
      variables: { input },
      onCompleted: resolve,
      updater: (store, data) => {
        // const newEdge = data.predictionCreate.result;
        // const requesterId = newEdge.node.requester?.id;

        // if (requesterId) {
        //   const conn = ConnectionHandler.getConnection(
        //     store.get(requesterId),
        //     ''
        //   )
        //   ConnectionHandler.insertEdgeAfter(conn, newEdge);
        // }
      }
    }),
  );
}