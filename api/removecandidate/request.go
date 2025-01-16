package removecandidate

type RequestBody struct {
	Password  string `json:"password"`
	Candidate string `json:"candidate"`
}
