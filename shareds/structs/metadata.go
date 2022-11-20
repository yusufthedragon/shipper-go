package structs

type Metadata struct {
	Path           string `json:"path"`
	HttpStatusCode int    `json:"http_status_code"`
	HttpStatus     string `json:"http_status"`
	Timestamp      int    `json:"timestamp"`
	Errors         Error  `json:"errors"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
