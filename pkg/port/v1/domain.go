package v1

type PostRequest struct {
	ID     string   `json:"id"`
	Values []string `json:"values"`
}

type PostResponse struct {
	ID string `json:"id"`
}
