package awb

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/yusufthedragon/shipper-go"
)

// The validator instance.
var validation = validator.New()

// Generate generates AWB Number from related logistic, in case AWB Number is not generated yet where order sent.
func Generate(params *GenerateParams) (GeneratedAWB, error) {
	return GenerateWithContext(context.Background(), params)
}

// GenerateWithContext generates AWB Number from related logistic, in case AWB Number is not generated yet where order sent with context.
func GenerateWithContext(ctx context.Context, params *GenerateParams) (GeneratedAWB, error) {
	var errValidation = validation.Struct(params)

	if errValidation != nil {
		log.Fatalln(errValidation.Error())
	}

	var endpoint = shipper.Conf.BaseURL + "'/awbs/generate"
	var responseStruct = GeneratedAWB{}
	var JSONParams, errEncode = json.Marshal(params)

	if errEncode != nil {
		log.Fatalln(errEncode.Error())
	}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:            ctx,
		HTTPMethod:     "GET",
		Endpoint:       endpoint,
		AdditionalBody: bytes.NewBuffer(JSONParams),
	}, &responseStruct)

	return responseStruct, err
}
