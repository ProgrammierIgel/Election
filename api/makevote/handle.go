package makevote

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/programmierigel/voting/storage"
	"github.com/programmierigel/voting/voting"
)

func Handle(store storage.Store) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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

		err = store.InsertVote(voting.Vote{
			ID:        requestBody.ID,
			Candidate: requestBody.Candidate,
		})

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return

		}

		response.Header().Set("Content-Type", "application/json")
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte(http.StatusText(http.StatusOK)))
	}
}
