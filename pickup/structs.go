package pickup

import "time"

// CreatedPickup struct contains response from API CreatePickup.
type CreatedPickup struct {
	Status string           `json:"status"`
	Data   DataCreatePickup `json:"data"`
}
type DataCreatePickup struct {
	Title      string `json:"title"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type CreatePickupV3 struct {
	Metadata struct {
		Path           string `json:"path"`
		HttpStatusCode int    `json:"http_status_code"`
		HttpStatus     string `json:"http_status"`
		Timestamp      int    `json:"timestamp"`
	} `json:"metadata"`
	Data struct {
		OrderActivations []OrderActivations `json:"order_activations"`
	} `json:"data"`
}

type OrderActivations struct {
	OrderId    string    `json:"order_id"`
	PickupCode string    `json:"pickup_code"`
	IsActivate bool      `json:"is_activate"`
	PickupTime time.Time `json:"pickup_time"`
}

// CreatePickupParams struct contains request parameter for API CreatePickup.
type CreatePickupParams struct {
	OrderIDs   []string `json:"orderIds" validate:"required"`
	DatePickup string   `json:"datePickup" validate:"required"`
}

type CreatePickupParamsV3 struct {
	Data DataRequestPickup `json:"data"`
}

type DataRequestPickup struct {
	OrderActivation OrderActivation `json:"order_activation"`
}

type OrderActivation struct {
	OrderId    []string `json:"order_id"`
	PickupTime string   `json:"pickup_time"`
}

func (r CreatePickupParams) ToCreatePickupParamsV3() CreatePickupParamsV3 {
	return CreatePickupParamsV3{
		Data: DataRequestPickup{
			OrderActivation: OrderActivation{
				OrderId:    r.OrderIDs,
				PickupTime: r.DatePickup,
			},
		},
	}
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
	PickupCode string `json:"pickupCode" validate:"required"`
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
		Title      string            `json:"title"`
		Message    string            `json:"message"`
		StatusCode int               `json:"statusCode"`
		Timezones  string            `json:"timezones"`
		Data       []DataTimeSlotsV3 `json:"timeSlots"`
	} `json:"data"`
}

type TimeSlotsV3 struct {
	Metadata struct {
		Path           string `json:"path"`
		HttpStatusCode int    `json:"http_status_code"`
		HttpStatus     string `json:"http_status"`
		Timestamp      int    `json:"timestamp"`
	} `json:"metadata"`
	Data struct {
		TimeZone  string `json:"time_zone"`
		TimeSlots []struct {
			StartTime string `json:"start_time"`
			EndTime   string `json:"end_time"`
		} `json:"time_slots"`
	} `json:"data"`
}

type DataTimeSlotsV3 struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func (r TimeSlotsV3) ToTimeSlot() TimeSlots {
	var data []DataTimeSlotsV3
	for _, val := range r.Data.TimeSlots {
		data = append(data, DataTimeSlotsV3{
			StartTime: val.StartTime,
			EndTime:   val.EndTime,
		})
	}
	return TimeSlots{
		Status: r.Metadata.HttpStatus,
		Data: struct {
			Title      string            `json:"title"`
			Message    string            `json:"message"`
			StatusCode int               `json:"statusCode"`
			Timezones  string            `json:"timezones"`
			Data       []DataTimeSlotsV3 `json:"timeSlots"`
		}{
			Title:      "timeslots",
			Message:    "success get timeslots",
			StatusCode: r.Metadata.HttpStatusCode,
			Timezones:  r.Data.TimeZone,
			Data:       data,
		},
	}
}
