package makevote

type RequestBody struct {
	ID        string `json:"id"`
	Candidate string `json:"candidate"`
}
