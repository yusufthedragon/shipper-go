package awb

// GenerateParams struct contains request parameter for API GenerateAWB.
type GenerateParams struct {
	EID string `json:"eid" validate:"required"`
	OID string `json:"oid" validate:"required"`
}

// GeneratedAWB struct contains response from API GenerateAWB.
type GeneratedAWB struct {
	Status string `json:"status"`
	Data   struct {
		Message    string `json:"message"`
		AWBNumber  string `json:"awbNumber"`
		StatusCode int    `json:"statusCode"`
	} `json:"data"`
}
