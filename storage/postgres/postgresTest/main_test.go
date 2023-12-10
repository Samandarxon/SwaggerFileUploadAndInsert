package postgres_test

import (
	"database/sql"
	"essy_travel/config"
	"essy_travel/storage/postgres"
	"fmt"
	"os"
	"testing"
)

var (
	countryTestRepo *postgres.CountryRepo
	airportestRepo  *postgres.AirportRepo
	citytestRepo    *postgres.CityRepo
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}

	countryTestRepo = postgres.NewCountryRepo(db)
	airportestRepo = postgres.NewAirportRepo(db)
	citytestRepo = postgres.NewCityRepo(db)

	os.Exit(m.Run())
}
