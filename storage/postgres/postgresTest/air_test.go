package postgres_test

import (
	"context"
	"essy_travel/models"
	"math/rand"
	"strconv"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func TestCreateAirport(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateAirport
		Output  *models.Airport
		WantErr bool
	}{{
		Name: "Case 1",
		Input: &models.CreateAirport{
			Title:        faker.Sentence(),
			Guid:         uuid.NewString(),
			CountryId:    uuid.NewString(),
			CityId:       uuid.NewString(),
			Latitude:     faker.Latitude(),
			Longitude:    faker.Longitude(),
			Radius:       strconv.Itoa(rand.Intn(10)),
			Image:        faker.URL(),
			Address:      faker.Century(),
			TimezoneId:   faker.UUIDHyphenated(),
			Country:      faker.Century(),
			City:         faker.Century(),
			SearchText:   faker.DomainName(),
			Code:         faker.Word(),
			ProductCount: rand.Intn(10),
			Gmt:          faker.CCNumber(),
		},
	}}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()
			data, err := airportestRepo.Create(ctx, *test.Input)
			if test.WantErr || err != nil {
				t.Errorf("%s: got: %+v", test.Name, test.Input)
				return
			}

			if data.Title != test.Input.Title &&
				data.Guid != test.Input.Guid &&
				data.CountryId != test.Input.CountryId &&
				data.CityId != test.Input.CityId &&
				data.Latitude != test.Input.Latitude &&
				data.Longitude != test.Input.Longitude &&
				data.Radius != test.Input.Radius &&
				data.Image != test.Input.Image &&
				data.Address != test.Input.Address &&
				data.TimezoneId != test.Input.TimezoneId &&
				data.Country != test.Input.Country &&
				data.City != test.Input.City &&
				data.SearchText != test.Input.SearchText &&
				data.Code != test.Input.Code &&
				data.ProductCount != test.Input.ProductCount && data.Gmt != test.Input.Gmt {
				t.Errorf("%s: no mistmach data got %+v", test.Name, data)
			}

		})
	}
}

func TestGetByIdAirport(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.AirportPrimaryKey
		Output  *models.Airport
		WantErr bool
	}{
		{
			Name:  "Case 1",
			Input: &models.AirportPrimaryKey{Id: "0e952ad9-fc66-4a28-874b-9b1eccc5e6ec"},
		}, {
			Name:  "Case 2",
			Input: &models.AirportPrimaryKey{Id: "5facc0ae-763b-4946-b7c5-a61c8f523ce1"},
		}, {
			Name:  "Case 3",
			Input: &models.AirportPrimaryKey{Id: "414a60e3-bc40-433a-aa79-f3dd14c8227c"},
		}, {
			Name:  "Case 4",
			Input: &models.AirportPrimaryKey{Id: "05fa5389-6815-47d4-9a2c-4631e76ed9d3"},
		}, {
			Name:  "Case 5",
			Input: &models.AirportPrimaryKey{Id: "ba470591-5080-4817-b551-ba12defe7e6c"},
		}, {
			Name:  "Case 6",
			Input: &models.AirportPrimaryKey{Id: "215067da-8dc4-4595-815c-839475b3aa03"},
		}, {
			Name:  "Case 7",
			Input: &models.AirportPrimaryKey{Id: "087f682f-07a6-4c2e-bdca-60f462a7cc62"},
		}, {
			Name:  "Case 8",
			Input: &models.AirportPrimaryKey{Id: "18f7ab7d-3a74-45b0-8b84-654cc74e29a6"},
		}, {
			Name:  "Case 9",
			Input: &models.AirportPrimaryKey{Id: "06630d04-be7e-4b5b-b65d-1ecbc8dd679a"},
		}, {
			Name:  "Case 10",
			Input: &models.AirportPrimaryKey{Id: "289f1a47-f445-4ddb-8ca3-9e977524dab2"},
		}, {
			Name:  "Case 11",
			Input: &models.AirportPrimaryKey{Id: "b47bcba2-66f0-4f99-8382-b51f672d3499"},
		}, {
			Name:  "Case 12",
			Input: &models.AirportPrimaryKey{Id: "20dd8312-8b59-4adf-8be6-5f344d878efd"},
		}, {
			Name:  "Case 13",
			Input: &models.AirportPrimaryKey{Id: "e500b92c-d347-4399-acab-6943a2d43cbd"},
		}, {
			Name:  "Case 14",
			Input: &models.AirportPrimaryKey{Id: "d336e5dc-60f2-4fcc-ab40-0ac8de8130fd"},
		}, {
			Name:  "Case 15",
			Input: &models.AirportPrimaryKey{Id: "1d202e50-5baf-4b2e-bb90-0f8a0910204c"},
		},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()

			data, err := airportestRepo.GetById(ctx, *test.Input)

			if test.WantErr || err != nil {
				t.Errorf("%s: got: %+v", test.Name, test.Input)
				return
			}

			if data.Id != test.Input.Id {
				t.Errorf("%s: no mistmach data got %+v", test.Name, data)
			}

		})
	}
}

func TestUpdateAirport(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.UpdateAirport
		Output  *models.Airport
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateAirport{
				Id:           "49965e9c-d369-41ef-84ba-4362029e38b3",
				Title:        faker.Sentence(),
				Guid:         uuid.NewString(),
				CountryId:    uuid.NewString(),
				CityId:       uuid.NewString(),
				Latitude:     faker.Latitude(),
				Longitude:    faker.Longitude(),
				Radius:       strconv.Itoa(rand.Intn(10)),
				Image:        faker.URL(),
				Address:      faker.Century(),
				TimezoneId:   faker.UUIDHyphenated(),
				Country:      faker.Century(),
				City:         faker.Century(),
				SearchText:   faker.DomainName(),
				Code:         faker.Word(),
				ProductCount: rand.Intn(10),
				Gmt:          faker.CCNumber(),
			},
		},
		{
			Name: "Case 2",
			Input: &models.UpdateAirport{
				Id:           "b4f9b63e-e9e3-441a-8491-9e98f9ee4558",
				Title:        faker.Sentence(),
				Guid:         uuid.NewString(),
				CountryId:    uuid.NewString(),
				CityId:       uuid.NewString(),
				Latitude:     faker.Latitude(),
				Longitude:    faker.Longitude(),
				Radius:       strconv.Itoa(rand.Intn(10)),
				Image:        faker.URL(),
				Address:      faker.Century(),
				TimezoneId:   faker.UUIDHyphenated(),
				Country:      faker.Century(),
				City:         faker.Century(),
				SearchText:   faker.DomainName(),
				Code:         faker.Word(),
				ProductCount: rand.Intn(10),
				Gmt:          faker.CCNumber(),
			},
		},
		{
			Name: "Case 3",
			Input: &models.UpdateAirport{
				Id:           "cb7d9264-1e92-4a3a-8fa4-e9f331938d53",
				Title:        faker.Sentence(),
				Guid:         uuid.NewString(),
				CountryId:    uuid.NewString(),
				CityId:       uuid.NewString(),
				Latitude:     faker.Latitude(),
				Longitude:    faker.Longitude(),
				Radius:       strconv.Itoa(rand.Intn(10)),
				Image:        faker.URL(),
				Address:      faker.Century(),
				TimezoneId:   faker.UUIDHyphenated(),
				Country:      faker.Century(),
				City:         faker.Century(),
				SearchText:   faker.DomainName(),
				Code:         faker.Word(),
				ProductCount: rand.Intn(10),
				Gmt:          faker.CCNumber(),
			},
		},
		{
			Name: "Case 4",
			Input: &models.UpdateAirport{
				Id:           "01ee894d-efee-4625-9795-bcfc12a7a012",
				Title:        faker.Sentence(),
				Guid:         uuid.NewString(),
				CountryId:    uuid.NewString(),
				CityId:       uuid.NewString(),
				Latitude:     faker.Latitude(),
				Longitude:    faker.Longitude(),
				Radius:       strconv.Itoa(rand.Intn(10)),
				Image:        faker.URL(),
				Address:      faker.Century(),
				TimezoneId:   faker.UUIDHyphenated(),
				Country:      faker.Century(),
				City:         faker.Century(),
				SearchText:   faker.DomainName(),
				Code:         faker.Word(),
				ProductCount: rand.Intn(10),
				Gmt:          faker.CCNumber(),
			},
		},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()

			data, err := airportestRepo.Update(ctx, *test.Input)

			if test.WantErr || err != nil {
				t.Errorf("%s: got: %+v", test.Name, test.Input)
				return
			}

			if data.Id != test.Input.Id && data.Title != test.Input.Title &&
				data.Guid != test.Input.Guid &&
				data.CountryId != test.Input.CountryId &&
				data.CityId != test.Input.CityId &&
				data.Latitude != test.Input.Latitude &&
				data.Longitude != test.Input.Longitude &&
				data.Radius != test.Input.Radius &&
				data.Image != test.Input.Image &&
				data.Address != test.Input.Address &&
				data.TimezoneId != test.Input.TimezoneId &&
				data.Country != test.Input.Country &&
				data.City != test.Input.City &&
				data.SearchText != test.Input.SearchText &&
				data.Code != test.Input.Code &&
				data.ProductCount != test.Input.ProductCount && data.Gmt != test.Input.Gmt {
				t.Errorf("%s: no mistmach data got %+v", test.Name, data)
			}

		})
	}
}

// func TestDeleteAirport(t *testing.T) {

// 	tests := []struct {
// 		Name    string
// 		Input   *models.AirportPrimaryKey
// 		Output  error
// 		WantErr bool
// 	}{
// 		{
// 			Name:  "Case 1",
// 			Input: &models.AirportPrimaryKey{Id: "c286497b-57b4-496b-a367-aa897e329cef"},
// 		},
// 		{
// 			Name:  "Case 2",
// 			Input: &models.AirportPrimaryKey{Id: "8885c460-3566-46a7-8eb4-b3be41e8a9b0"},
// 		},
// 		{
// 			Name:  "Case 3",
// 			Input: &models.AirportPrimaryKey{Id: "a5895862-543c-44c5-a2f5-b8c741580ff6"},
// 		},
// 		{
// 			Name:  "Case 4",
// 			Input: &models.AirportPrimaryKey{Id: "7525732f-2f91-4a8b-967f-3dd0b5919e83"},
// 		},
// 	}

// 	for _, test := range tests {

// 		t.Run(test.Name, func(t *testing.T) {
// ctx := context.Background()

// 			data, err := airportestRepo.Delete(ctx,*test.Input)

// 			if test.WantErr || err != nil {
// 				t.Errorf("%s: got: %+v", test.Name, test.Input)
// 				return
// 			}

// 			if data != "Deleted" {
// 				t.Errorf("%s: no mistmach data got %+v", test.Name, data)
// 			}

// 		})
// 	}
// }
