package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/allscape/bigants/grasshopper/app/gql"
)

func newGraphQLHandler(gqls *gql.Service) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		ctx := r.Context()
		ctx = gqls.WithResolver(ctx)

		response := gqls.Exec(ctx, params.Query, params.OperationName, params.Variables)
		if len(response.Errors) > 0 {
			formattedError := fmt.Sprintf(
				"GraphQLHandlerError: %s, %v\n", params.Query, params.Variables,
			)
			for _, er := range response.Errors {
				formattedError += er.Error() + "\n"
			}
			log.Println(formattedError)
		}
		responseJSON, err := json.Marshal(response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
	})
}
