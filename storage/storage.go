package storage

import (
	"context"
	"essy_travel/models"
)

type StorageI interface {
	City() CityRepoI
	Airport() AirportRepoI
	Country() CountryRepoI
}

type CountryRepoI interface {
	Create(ctx context.Context, req models.CreateCountry) (*models.Country, error)
	Update(ctx context.Context, req models.UpdateCountry) (*models.Country, error)
	Upload(ctx context.Context, req models.UploadCountry) error
	GetById(ctx context.Context, req models.CountryPrimaryKey) (*models.Country, error)
	GetList(ctx context.Context, req models.GetListCountryRequest) (*models.GetListCountryResponse, error)
	Delete(ctx context.Context, req models.CountryPrimaryKey) (string, error)
}

type CityRepoI interface {
	Create(ctx context.Context, req models.CreateCity) (*models.City, error)
	Update(ctx context.Context, req models.UpdateCity) (*models.City, error)
	Upload(ctx context.Context, req models.UploadCity) error
	GetById(ctx context.Context, req models.CityPrimaryKey) (*models.City, error)
	GetList(ctx context.Context, req models.GetListCityRequest) (*models.GetListCityResponse, error)
	Delete(ctx context.Context, req models.CityPrimaryKey) (string, error)
}

type AirportRepoI interface {
	Create(ctx context.Context, req models.CreateAirport) (*models.Airport, error)
	Update(ctx context.Context, req models.UpdateAirport) (*models.Airport, error)
	Upload(ctx context.Context, req models.UploadAirport) error
	GetById(ctx context.Context, req models.AirportPrimaryKey) (*models.Airport, error)
	GetList(ctx context.Context, req models.GetListAirportRequest) (*models.GetListAirportResponse, error)
	Delete(ctx context.Context, req models.AirportPrimaryKey) (string, error)
}
