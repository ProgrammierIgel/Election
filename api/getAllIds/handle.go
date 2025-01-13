package getallids

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/programmierigel/voting/storage"
)

func Handle(store storage.Store) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		requestBytes, err := io.ReadAll(io.LimitReader(request.Body, 4096))
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		getAllUndefinedVotes, err := store.GetAllUndefinedVotes()
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}

		responseBody := ResponseBody{
			Votes: getAllUndefinedVotes,
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(responseBytes)

		// fmt.Println(store.IsVotingActive())
	}
}
