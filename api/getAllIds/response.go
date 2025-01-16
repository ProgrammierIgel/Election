package getallids

import (
	"github.com/programmierigel/voting/voting"
)

type ResponseBody struct {
	Votes []voting.AllVotes `json:"votes"`
}
