package tracking

import (
	"context"

	"github.com/yusufthedragon/shipper-go"
)

// GetAllStatus gets all Tracking status.
func GetAllStatus() (AllStatus, error) {
	return GetAllStatusWithContext(context.Background())
}

// GetAllStatusWithContext gets all Tracking status with context.
func GetAllStatusWithContext(ctx context.Context) (AllStatus, error) {
	var endpoint = shipper.Conf.BaseURL + "/logistics/status"
	var responseStruct = AllStatus{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
	}, &responseStruct)

	return responseStruct, err
}
