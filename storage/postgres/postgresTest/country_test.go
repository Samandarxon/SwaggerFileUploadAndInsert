package postgres_test

import (
	"context"
	"essy_travel/models"
	"testing"

	"github.com/google/uuid"
)

func TestCreateCountry(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateCountry
		Output  *models.Country
		WantErr bool
	}{{
		Name: "Case 1",
		Input: &models.CreateCountry{
			Guid:      uuid.NewString(),
			Title:     "Notebook",
			Code:      "AS",
			Continent: "EU",
		},
	}}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()

			data, err := countryTestRepo.Create(ctx, *test.Input)
			if test.WantErr || err != nil {
				t.Errorf("%s: got: %+v", test.Name, test.Input)
				return
			}

			if data.Code != test.Input.Code && data.Title != test.Input.Title &&
				data.Continent != test.Input.Continent {
				t.Errorf("%s: no mistmach data got %+v", test.Name, data)
			}

		})
	}
}

func TestGetByIdCountry(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CountryPrimaryKey
		Output  *models.Country
		WantErr bool
	}{
		{
			Name:  "Case 1",
			Input: &models.CountryPrimaryKey{Id: "176b5b7e-7923-448e-a884-55c2be5589ce"},
		}, {
			Name:  "Case 2",
			Input: &models.CountryPrimaryKey{Id: "6c8d345b-bf84-4825-b6f7-70f369cb595e"},
		}, {
			Name:  "Case 3",
			Input: &models.CountryPrimaryKey{Id: "a3860a41-d834-469e-880e-d1303746de3e"},
		}, {
			Name:  "Case 4",
			Input: &models.CountryPrimaryKey{Id: "2a3d9631-0de7-4287-93be-0f2e9b288749"},
		}, {
			Name:  "Case 5",
			Input: &models.CountryPrimaryKey{Id: "c6a9e69d-5122-4dfa-8acf-83823b1b2063"},
		}, {
			Name:  "Case 6",
			Input: &models.CountryPrimaryKey{Id: "4e83a5c8-a0cb-4725-8757-e28fec0c4e5c"},
		}, {
			Name:  "Case 7",
			Input: &models.CountryPrimaryKey{Id: "888fa9da-ff73-4839-8428-8c5e809d3970"},
		}, {
			Name:  "Case 8",
			Input: &models.CountryPrimaryKey{Id: "7540bb71-9969-4bf7-8a7e-76e0981aff2e"},
		}, {
			Name:  "Case 9",
			Input: &models.CountryPrimaryKey{Id: "1b116a64-45cc-46f9-a94d-723663f7aa12"},
		}, {
			Name:  "Case 10",
			Input: &models.CountryPrimaryKey{Id: "210eeb20-80d1-483e-b877-3200655ecc33"},
		}, {
			Name:  "Case 11",
			Input: &models.CountryPrimaryKey{Id: "fcc77c22-b60d-4590-ac4a-eb19aacb4379"},
		}, {
			Name:  "Case 12",
			Input: &models.CountryPrimaryKey{Id: "4192884d-5f81-4524-aee5-38f911b92bbd"},
		}, {
			Name:  "Case 13",
			Input: &models.CountryPrimaryKey{Id: "626cea43-c0b1-4f3d-8118-96462348050d"},
		}, {
			Name:  "Case 14",
			Input: &models.CountryPrimaryKey{Id: "78a1e666-5054-4b33-89c7-c252c5b6767f"},
		}, {
			Name:  "Case 15",
			Input: &models.CountryPrimaryKey{Id: "7938b8e1-18b1-4a90-bbe4-bbf867821fa4"},
		},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()

			data, err := countryTestRepo.GetById(ctx, *test.Input)

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

func TestUpdateCountry(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.UpdateCountry
		Output  *models.Country
		WantErr bool
	}{
		{
			Name:  "Case 1",
			Input: &models.UpdateCountry{Id: "176b5b7e-7923-448e-a884-55c2be5589ce", Guid: uuid.NewString(), Title: "Book", Code: "ER", Continent: "rwerwersdf"},
		},
		{
			Name:  "Case 2",
			Input: &models.UpdateCountry{Id: "6c8d345b-bf84-4825-b6f7-70f369cb595e", Guid: uuid.NewString(), Title: "Apple", Code: "UK", Continent: "rwesdsd"},
		},
		{
			Name:  "Case 3",
			Input: &models.UpdateCountry{Id: "a3860a41-d834-469e-880e-d1303746de3e", Guid: uuid.NewString(), Title: "Phone", Code: "QW", Continent: "adasd"},
		},
		{
			Name:  "Case 4",
			Input: &models.UpdateCountry{Id: "2a3d9631-0de7-4287-93be-0f2e9b288749", Guid: uuid.NewString(), Title: "PC", Code: "EN", Continent: "sdas"},
		},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()
			data, err := countryTestRepo.Update(ctx, *test.Input)

			if test.WantErr || err != nil {
				t.Errorf("%s: got: %+v", test.Name, test.Input)
				return
			}

			if data.Id != test.Input.Id && data.Code != test.Input.Code && data.Title != test.Input.Title &&
				data.Continent != test.Input.Continent {
				t.Errorf("%s: no mistmach data got %+v", test.Name, data)
			}

		})
	}
}

// func TestDeleteCountry(t *testing.T) {

// 	tests := []struct {
// 		Name    string
// 		Input   *models.CountryPrimaryKey
// 		Output  error
// 		WantErr bool
// 	}{
// 		{
// 			Name:  "Case 1",
// 			Input: &models.CountryPrimaryKey{Id: "e50fa25e-bb8c-494e-a468-d27652846068"},
// 		},
// 		{
// 			Name:  "Case 2",
// 			Input: &models.CountryPrimaryKey{Id: "874be166-2fc8-4a35-a474-8ff7ceeaffd2"},
// 		},
// 		{
// 			Name:  "Case 3",
// 			Input: &models.CountryPrimaryKey{Id: "30f874d3-4da5-4f81-8e2d-c873868b0423"},
// 		},
// 		{
// 			Name:  "Case 4",
// 			Input: &models.CountryPrimaryKey{Id: "50012da7-d012-4156-9590-71bf21dc0145"},
// 		},
// 	}

// 	for _, test := range tests {

// 		t.Run(test.Name, func(t *testing.T) {
// ctx := context.Background()
// 			data, err := countryTestRepo.Delete(ctx,*test.Input)

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
