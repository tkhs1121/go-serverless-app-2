package main

type Payload struct {
	Body string `json:"body"`
}

type Request struct {
	URL string `json:"url"`
}
