package votingactive

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/programmierigel/voting/storage"
)

func Handle(store storage.Store) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
    response.Header().Set("Access-Control-Allow-Origin", "*")

		status := store.IsVotingActive()

		responseBody := ResponseBody{
			Status: status,
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(responseBytes)
	}
}
