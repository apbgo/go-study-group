package model

type Request struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}
