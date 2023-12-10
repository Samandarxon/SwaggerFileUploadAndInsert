package postgres

import (
	"context"
	"essy_travel/config"
	"essy_travel/storage"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *pgxpool.Pool
	city    *CityRepo
	country *CountryRepo
	airport *AirportRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
			cfg.PostgresHost,
			cfg.PostgresUser,
			cfg.PostgresDatabase,
			cfg.PostgresPassword,
			cfg.PostgresPort,
		),
	)
	if err != nil {
		return nil, err
	}

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	return &Store{
		db: pgxpool,
	}, nil
}

func (s *Store) City() storage.CityRepoI {
	if s.city == nil {
		s.city = NewCityRepo(s.db)
	}
	return s.city
}

func (s *Store) Airport() storage.AirportRepoI {
	if s.airport == nil {
		s.airport = NewAirportRepo(s.db)
	}
	return s.airport
}

func (s *Store) Country() storage.CountryRepoI {
	if s.country == nil {
		s.country = NewCountryRepo(s.db)
	}
	return s.country
}
