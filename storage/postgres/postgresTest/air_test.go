package postgres_test

import (
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
			data, err := airportestRepo.Create(*test.Input)
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
			Input: &models.AirportPrimaryKey{Id: "c62dee05-53c7-4cde-abef-50aff8c2fd97"},
		}, {
			Name:  "Case 2",
			Input: &models.AirportPrimaryKey{Id: "75f2f60b-95a9-48a6-9a93-18b6b372b278"},
		}, {
			Name:  "Case 3",
			Input: &models.AirportPrimaryKey{Id: "8aa0baff-3c20-464d-ba9a-cbf30ad5b68b"},
		}, {
			Name:  "Case 4",
			Input: &models.AirportPrimaryKey{Id: "2360d384-03be-46e9-98cd-2f4f32f71604"},
		}, {
			Name:  "Case 5",
			Input: &models.AirportPrimaryKey{Id: "3a4cfd5f-17a1-45e7-a8bc-82c44dc76628"},
		}, {
			Name:  "Case 6",
			Input: &models.AirportPrimaryKey{Id: "f0a4526f-0dd8-4105-acd4-64e2aed46bdd"},
		}, {
			Name:  "Case 7",
			Input: &models.AirportPrimaryKey{Id: "565ffd4e-6646-4ee1-9cc6-1c32340b921e"},
		}, {
			Name:  "Case 8",
			Input: &models.AirportPrimaryKey{Id: "df843f12-d04d-4ce6-85d0-c1489f6f4b40"},
		}, {
			Name:  "Case 9",
			Input: &models.AirportPrimaryKey{Id: "7effd7cd-d9a4-48de-a273-1879f0c134a4"},
		}, {
			Name:  "Case 10",
			Input: &models.AirportPrimaryKey{Id: "efd9dbb5-9d98-469d-9171-5a6853f71cee"},
		}, {
			Name:  "Case 11",
			Input: &models.AirportPrimaryKey{Id: "3a9b96de-2f21-4232-966c-47abda1458fa"},
		}, {
			Name:  "Case 12",
			Input: &models.AirportPrimaryKey{Id: "bab6242d-f248-4dc0-91ac-1125f44b6f71"},
		}, {
			Name:  "Case 13",
			Input: &models.AirportPrimaryKey{Id: "da3c709c-6179-4b94-8e0e-8999b708f073"},
		}, {
			Name:  "Case 14",
			Input: &models.AirportPrimaryKey{Id: "87ba0c90-286a-4194-b91f-a2c51fc259ce"},
		}, {
			Name:  "Case 15",
			Input: &models.AirportPrimaryKey{Id: "682d474e-05d6-4f07-bbaf-c3c4934e08f5"},
		},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			data, err := airportestRepo.GetById(*test.Input)

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
				Id:           "bab6242d-f248-4dc0-91ac-1125f44b6f71",
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
				Id:           "da3c709c-6179-4b94-8e0e-8999b708f073",
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
				Id:           "87ba0c90-286a-4194-b91f-a2c51fc259ce",
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
				Id:           "682d474e-05d6-4f07-bbaf-c3c4934e08f5",
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
			data, err := airportestRepo.Update(*test.Input)

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
// 			Input: &models.AirportPrimaryKey{Id: "bab6242d-f248-4dc0-91ac-1125f44b6f71"},
// 		},
// 		{
// 			Name:  "Case 2",
// 			Input: &models.AirportPrimaryKey{Id: "da3c709c-6179-4b94-8e0e-8999b708f073"},
// 		},
// 		{
// 			Name:  "Case 3",
// 			Input: &models.AirportPrimaryKey{Id: "87ba0c90-286a-4194-b91f-a2c51fc259ce"},
// 		},
// 		{
// 			Name:  "Case 4",
// 			Input: &models.AirportPrimaryKey{Id: "682d474e-05d6-4f07-bbaf-c3c4934e08f5"},
// 		},
// 	}

// 	for _, test := range tests {

// 		t.Run(test.Name, func(t *testing.T) {
// 			data, err := airportestRepo.Delete(*test.Input)

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
