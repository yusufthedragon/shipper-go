package pickup

import "time"

// CreatedPickup struct contains response from API CreatePickup.
type CreatedPickup struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
	} `json:"data"`
}

// CreatePickupParams struct contains request parameter for API CreatePickup.
type CreatePickupParams struct {
	OrderIDs   []string `json:"orderIds" validate:"required"`
	DatePickup string   `json:"datePickup" validate:"required"`
}

// CancelledPickup struct contains response from API CancelPickup.
type CancelledPickup struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
	} `json:"data"`
}

// CancelPickupParams struct contains request parameter for API CancelPickup.
type CancelPickupParams struct {
	OrderIDs []string `json:"orderIds" validate:"required"`
}

// Agents struct contains response from API GetAgents.
type Agents struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
		Data       []struct {
			Agent struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"agent"`
			City struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"city"`
			Location struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"location"`
			Contact struct {
				Name  string `json:"name"`
				Phone string `json:"phone"`
			} `json:"contact"`
		} `json:"data"`
	} `json:"data"`
}

// TimeSlots struct contains response from API GetPickupTimeSlots.
type TimeSlots struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
		Timezones  string `json:"timezones"`
		Data       []struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"timeSlots"`
	} `json:"data"`
}
