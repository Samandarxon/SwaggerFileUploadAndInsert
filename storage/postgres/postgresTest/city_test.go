package postgres_test

import (
	"context"
	"essy_travel/models"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func TestCreateCity(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateCity
		Output  *models.City
		WantErr bool
	}{{
		Name: "Case 1",
		Input: &models.CreateCity{
			Guid:        uuid.NewString(),
			Title:       faker.LastName(),
			CountryId:   uuid.NewString(),
			CityCode:    strings.ToUpper(faker.FirstName()[:2]),
			Latitude:    fmt.Sprintf("%f", faker.Latitude()),
			Longitude:   fmt.Sprintf("%f", faker.Longitude()),
			Offset:      fmt.Sprintf("%d:%d", rand.Intn(10)+5, rand.Intn(5)+5),
			TimezoneId:  faker.UUIDHyphenated(),
			CountryName: strings.ToUpper(faker.LastName()[:2]),
		},
	}}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()

			data, err := citytestRepo.Create(ctx, *test.Input)
			if test.WantErr || err != nil {
				t.Errorf("%s: got: %+v", test.Name, test.Input)
				return
			}

			if data.Guid != test.Input.Guid &&
				data.Title != test.Input.Title &&
				data.CountryId != test.Input.CountryId &&
				data.CityCode != test.Input.CityCode &&
				data.Latitude != test.Input.Latitude &&
				data.Longitude != test.Input.Longitude &&
				data.Offset != test.Input.Offset &&
				data.TimezoneId != test.Input.TimezoneId &&
				data.CountryName != test.Input.CountryName {
				t.Errorf("%s: no mistmach data got %+v", test.Name, data)
			}
		})
	}
}

func TestGetByIdCity(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CityPrimaryKey
		Output  *models.City
		WantErr bool
	}{
		{
			Name:  "Case 1",
			Input: &models.CityPrimaryKey{Id: "046f203b-dbcb-4ad8-b35c-c6e9503f2f1f"},
		}, {
			Name:  "Case 2",
			Input: &models.CityPrimaryKey{Id: "28d86b1d-6a5e-48ae-a438-9ef9f09ce32f"},
		}, {
			Name:  "Case 3",
			Input: &models.CityPrimaryKey{Id: "45d74e7f-80fa-46cf-a2dd-20e5f68b62b0"},
		}, {
			Name:  "Case 4",
			Input: &models.CityPrimaryKey{Id: "759576a0-945b-4992-bb7a-23151b190263"},
		}, {
			Name:  "Case 5",
			Input: &models.CityPrimaryKey{Id: "cc85aa53-d567-44fe-8415-15fc42b2f801"},
		}, {
			Name:  "Case 6",
			Input: &models.CityPrimaryKey{Id: "acd64aa8-75a0-4943-b146-0723f3ed8420"},
		}, {
			Name:  "Case 7",
			Input: &models.CityPrimaryKey{Id: "7c25b22c-1b4a-4bad-bc2d-e6941c8b5131"},
		}, {
			Name:  "Case 8",
			Input: &models.CityPrimaryKey{Id: "d9c9a115-61cc-4391-9c3f-e3ecc28afaa7"},
		}, {
			Name:  "Case 9",
			Input: &models.CityPrimaryKey{Id: "3d3aab45-6274-4ea0-976e-4e6e6e34e124"},
		}, {
			Name:  "Case 10",
			Input: &models.CityPrimaryKey{Id: "6e36543a-51eb-405c-900f-d90a0c09b46b"},
		}, {
			Name:  "Case 11",
			Input: &models.CityPrimaryKey{Id: "943a3607-9c5b-4224-b71a-d2350b07d9be"},
		}, {
			Name:  "Case 12",
			Input: &models.CityPrimaryKey{Id: "e2609e57-ca80-43ad-a365-325c7e5f0944"},
		}, {
			Name:  "Case 13",
			Input: &models.CityPrimaryKey{Id: "9a34c859-6ed0-4ece-8ced-81b4bc39a37f"},
		}, {
			Name:  "Case 14",
			Input: &models.CityPrimaryKey{Id: "d0fc1fc3-1248-42ce-b10e-77a3dccfa57f"},
		}, {
			Name:  "Case 15",
			Input: &models.CityPrimaryKey{Id: "afee2a06-8ed6-4955-9c42-e61a4031ae3d"},
		},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()

			data, err := citytestRepo.GetById(ctx, *test.Input)

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

func TestUpdateCity(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.UpdateCity
		Output  *models.City
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateCity{
				Id:          "d67de33d-534e-402f-beb6-7c289cfed150",
				Guid:        uuid.NewString(),
				Title:       faker.LastName(),
				CountryId:   uuid.NewString(),
				CityCode:    strings.ToUpper(faker.FirstName()[:2]),
				Latitude:    fmt.Sprintf("%f", faker.Latitude()),
				Longitude:   fmt.Sprintf("%f", faker.Longitude()),
				Offset:      fmt.Sprintf("%d:%d", rand.Intn(10)+5, rand.Intn(5)+5),
				TimezoneId:  faker.UUIDHyphenated(),
				CountryName: strings.ToUpper(faker.LastName()[:2]),
			},
		},
		{
			Name: "Case 2",
			Input: &models.UpdateCity{
				Id:          "4522f4ef-1bd8-435f-b3c2-1500fcf2cb39",
				Guid:        uuid.NewString(),
				Title:       faker.LastName(),
				CountryId:   uuid.NewString(),
				CityCode:    strings.ToUpper(faker.FirstName()[:2]),
				Latitude:    fmt.Sprintf("%f", faker.Latitude()),
				Longitude:   fmt.Sprintf("%f", faker.Longitude()),
				Offset:      fmt.Sprintf("%d:%d", rand.Intn(10)+5, rand.Intn(5)+5),
				TimezoneId:  faker.UUIDHyphenated(),
				CountryName: strings.ToUpper(faker.LastName()[:2]),
			},
		},
		{
			Name: "Case 3",
			Input: &models.UpdateCity{
				Id:          "9f1cc281-9d01-4c55-a457-4d120e4c8218",
				Guid:        uuid.NewString(),
				Title:       faker.LastName(),
				CountryId:   uuid.NewString(),
				CityCode:    strings.ToUpper(faker.FirstName()[:2]),
				Latitude:    fmt.Sprintf("%f", faker.Latitude()),
				Longitude:   fmt.Sprintf("%f", faker.Longitude()),
				Offset:      fmt.Sprintf("%d:%d", rand.Intn(10)+5, rand.Intn(5)+5),
				TimezoneId:  faker.UUIDHyphenated(),
				CountryName: strings.ToUpper(faker.LastName()[:2]),
			},
		},
		{
			Name: "Case 4",
			Input: &models.UpdateCity{
				Id:          "6b1fb740-33d4-4526-8142-9a135f734690",
				Guid:        uuid.NewString(),
				Title:       faker.LastName(),
				CountryId:   uuid.NewString(),
				CityCode:    strings.ToUpper(faker.FirstName()[:2]),
				Latitude:    fmt.Sprintf("%f", faker.Latitude()),
				Longitude:   fmt.Sprintf("%f", faker.Longitude()),
				Offset:      fmt.Sprintf("%d:%d", rand.Intn(10)+5, rand.Intn(5)+5),
				TimezoneId:  faker.UUIDHyphenated(),
				CountryName: strings.ToUpper(faker.LastName()[:2]),
			},
		},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()

			data, err := citytestRepo.Update(ctx, *test.Input)

			if test.WantErr || err != nil {
				t.Errorf("%s: got: %+v", test.Name, test.Input)
				return
			}

			if data.Id != test.Input.Id &&
				data.Guid != test.Input.Guid &&
				data.Title != test.Input.Title &&
				data.CountryId != test.Input.CountryId &&
				data.CityCode != test.Input.CityCode &&
				data.Latitude != test.Input.Latitude &&
				data.Longitude != test.Input.Longitude &&
				data.Offset != test.Input.Offset &&
				data.TimezoneId != test.Input.TimezoneId &&
				data.CountryName != test.Input.CountryName {
				t.Errorf("%s: no mistmach data got %+v", test.Name, data)
			}

		})
	}
}

// func TestDeleteCity(t *testing.T) {

// 	tests := []struct {
// 		Name    string
// 		Input   *models.CityPrimaryKey
// 		Output  error
// 		WantErr bool
// 	}{
// 		{
// 			Name:  "Case 1",
// 			Input: &models.CityPrimaryKey{Id: "02ce0602-e08d-4e18-ac0c-b536fc7176b0"},
// 		},
// 		{
// 			Name:  "Case 2",
// 			Input: &models.CityPrimaryKey{Id: "c7c69f7a-04fa-46fb-9511-6d47165340cd"},
// 		},
// 		{
// 			Name:  "Case 3",
// 			Input: &models.CityPrimaryKey{Id: "c261ac67-976e-4d8e-b806-27c58802caea"},
// 		},
// 		{
// 			Name:  "Case 4",
// 			Input: &models.CityPrimaryKey{Id: "633a6b06-f53c-4d87-9ee4-bdf078ca8348"},
// 		},
// 	}

// 	for _, test := range tests {

// 		t.Run(test.Name, func(t *testing.T) {
// ctx := context.Background()

// 			data, err := citytestRepo.Delete(ctx,*test.Input)

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
