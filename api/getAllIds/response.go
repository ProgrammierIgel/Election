package getallids

import "github.com/programmierigel/voting/storage"

type ResponseBody struct {
	Votes []storage.AllVotes `json:"votes"`
}
