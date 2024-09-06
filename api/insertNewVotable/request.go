package insertnewvotable

type RequestBody struct {
	Password string `json:"password"`
	VoteID   string `json:"vote-id"`
}
