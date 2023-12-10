package postgres_test

import (
	"context"
	"essy_travel/config"
	"essy_travel/storage/postgres"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	countryTestRepo *postgres.CountryRepo
	airportestRepo  *postgres.AirportRepo
	citytestRepo    *postgres.CityRepo
)

func TestMain(m *testing.M) {
	cfg := config.Load()

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
		panic(err)
		return
	}

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)

	countryTestRepo = postgres.NewCountryRepo(pgxpool)
	airportestRepo = postgres.NewAirportRepo(pgxpool)
	citytestRepo = postgres.NewCityRepo(pgxpool)

	os.Exit(m.Run())
}
