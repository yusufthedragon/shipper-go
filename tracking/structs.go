package tracking

// AllStatus struct contains response from API GetAllStatus.
type AllStatus struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		StatusCode int    `json:"statusCode"`
		Rows       []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"rows"`
	} `json:"data"`
}
