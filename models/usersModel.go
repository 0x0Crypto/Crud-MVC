package models

// Request user is structure for request
type RequestUser struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// Response user is structure for response
type ResponseUser struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}
