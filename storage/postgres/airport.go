package postgres

import (
	"context"
	"database/sql"
	"essy_travel/models"
	"essy_travel/pkg/helpers"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AirportRepo struct {
	db *pgxpool.Pool
}

func NewAirportRepo(db *pgxpool.Pool) *AirportRepo {
	return &AirportRepo{
		db: db,
	}
}

func (a *AirportRepo) Create(ctx context.Context, req models.CreateAirport) (*models.Airport, error) {
	var id = uuid.New().String()
	var query = `
	INSERT INTO airport(
			"id",
			"guid",
			"title",
			"country_id",
			"city_id",
			"latitude",
			"longitude",
			"radius",
			"image",
			"address",
			"timezone_id",
			"country",
			"city",
			"search_text",
			"code",
			"product_count",
			"gmt",
			"updated_at"
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,NOW())
`
	// fmt.Println(query, req)
	_, err := a.db.Exec(ctx, query,
		id,
		req.Guid,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Latitude,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Address,
		req.TimezoneId,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
	)
	if err != nil {
		return &models.Airport{}, err
	}
	return a.GetById(ctx, models.AirportPrimaryKey{Id: id})
}

func (c *AirportRepo) GetById(ctx context.Context, req models.AirportPrimaryKey) (*models.Airport, error) {

	var (
		airport = models.Airport{}
		query   = `
		SELECT 
				"id",
				"guid",
				"title",
				"country_id",
				"city_id",
				"latitude",
				"longitude",
				"radius",
				"image",
				"address",
				"timezone_id",
				"country",
				"city",
				"search_text",
				"code",
				"product_count",
				"gmt",
				"created_at",
				"updated_at"
		FROM airport WHERE id=$1`
	)
	var (
		Id           sql.NullString
		Guid         sql.NullString
		Title        sql.NullString
		CountryId    sql.NullString
		CityId       sql.NullString
		Latitude     sql.NullFloat64
		Longitude    sql.NullFloat64
		Radius       sql.NullString
		Image        sql.NullString
		Address      sql.NullString
		TimezoneId   sql.NullString
		Country      sql.NullString
		City         sql.NullString
		SearchText   sql.NullString
		Code         sql.NullString
		ProductCount sql.NullInt64
		Gmt          sql.NullString
		CreatedAt    sql.NullString
		UpdatedAt    sql.NullString
	)
	fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&Guid,
		&Title,
		&CountryId,
		&CityId,
		&Latitude,
		&Longitude,
		&Radius,
		&Image,
		&Address,
		&TimezoneId,
		&Country,
		&City,
		&SearchText,
		&Code,
		&ProductCount,
		&Gmt,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	airport = models.Airport{
		Id:           Id.String,
		Guid:         Guid.String,
		Title:        Title.String,
		CountryId:    CountryId.String,
		CityId:       CityId.String,
		Latitude:     Latitude.Float64,
		Longitude:    Longitude.Float64,
		Radius:       Radius.String,
		Image:        Image.String,
		Address:      Address.String,
		TimezoneId:   TimezoneId.String,
		Country:      Country.String,
		City:         City.String,
		SearchText:   SearchText.String,
		Code:         Code.String,
		ProductCount: int(ProductCount.Int64),
		Gmt:          Gmt.String,
		CreatedAt:    CreatedAt.String,
		UpdatedAt:    UpdatedAt.String,
	}
	return &airport, nil
}

func (c *AirportRepo) GetList(ctx context.Context, req models.GetListAirportRequest) (*models.GetListAirportResponse, error) {
	var (
		airports = models.GetListAirportResponse{}

		Id           sql.NullString
		Guid         sql.NullString
		Title        sql.NullString
		CountryId    sql.NullString
		CityId       sql.NullString
		Latitude     sql.NullFloat64
		Longitude    sql.NullFloat64
		Radius       sql.NullString
		Image        sql.NullString
		Address      sql.NullString
		TimezoneId   sql.NullString
		Country      sql.NullString
		City         sql.NullString
		SearchText   sql.NullString
		Code         sql.NullString
		ProductCount sql.NullInt64
		Gmt          sql.NullString
		CreatedAt    sql.NullString
		UpdatedAt    sql.NullString

		where  = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		query  = `
				SELECT 
				COUNT(*) OVER(),
				"id",
				"guid",
				"title",
				"country_id",
				"city_id",
				"latitude",
				"longitude",
				"radius",
				"image",
				"address",
				"timezone_id",
				"country",
				"city",
				"search_text",
				"code",
				"product_count",
				"gmt",
				"created_at",
				"updated_at"
				FROM airport`
	)
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += where + offset + limit

	rows, err := c.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	fmt.Println(query)
	for rows.Next() {
		var airport models.Airport

		err := rows.Scan(
			&airports.Count,
			&Id,
			&Guid,
			&Title,
			&CountryId,
			&CityId,
			&Latitude,
			&Longitude,
			&Radius,
			&Image,
			&Address,
			&TimezoneId,
			&Country,
			&City,
			&SearchText,
			&Code,
			&ProductCount,
			&Gmt,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		airport = models.Airport{
			Id:           Id.String,
			Guid:         Guid.String,
			Title:        Title.String,
			CountryId:    CountryId.String,
			CityId:       CityId.String,
			Latitude:     Latitude.Float64,
			Longitude:    Longitude.Float64,
			Radius:       Radius.String,
			Image:        Image.String,
			Address:      Address.String,
			TimezoneId:   TimezoneId.String,
			Country:      Country.String,
			City:         City.String,
			SearchText:   SearchText.String,
			Code:         Code.String,
			ProductCount: int(ProductCount.Int64),
			Gmt:          Gmt.String,
			CreatedAt:    CreatedAt.String,
			UpdatedAt:    UpdatedAt.String,
		}
		// fmt.Printf("%+v\n", airport)
		airports.Airports = append(airports.Airports, airport)
	}
	defer rows.Close()
	return &airports, nil
}

func (c *AirportRepo) Update(ctx context.Context, req models.UpdateAirport) (*models.Airport, error) {
	var (
		query = `
			UPDATE airport SET
				"title"=$2,
				"country_id"=$3,
				"city_id"=$4,
				"latitude"=$5,
				"longitude"=$6,
				"radius"=$7,
				"image"=$8,
				"address"=$9,
				"timezone_id"=$10,
				"country"=$11,
				"city"=$12,
				"search_text"=$13,
				"code"=$14,
				"product_count"=$15,
				"gmt"=$16,
				updated_at = NOW() 
			WHERE id = $1`
	)
	_, err := c.db.Exec(ctx, query,
		req.Id,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Latitude,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Address,
		req.TimezoneId,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
	)
	if err != nil {
		return nil, err
	}

	return c.GetById(ctx, models.AirportPrimaryKey{Id: req.Id})
}

func (c *AirportRepo) Delete(ctx context.Context, req models.AirportPrimaryKey) (string, error) {

	_, err := c.db.Exec(ctx, `DELETE FROM airport WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *AirportRepo) Upload(ctx context.Context, req models.UploadAirport) error {
	defer os.Remove(req.DST)
	// fileName, err := helpers.Upload(&req.File, req.Base)
	// if err != nil {
	// 	return err
	// }

	dates, err := helpers.ReadAirport(req.DST)
	// fmt.Println(dates)
	if err != nil {
		return err
	}
	var (
		// str   string
		query = `
			INSERT INTO airport(
				"id",
				"guid",
				"title",
				"country_id",
				"city_id",
				"latitude",
				"longitude",
				"radius",
				"image",
				"address",
				"timezone_id",
				"country",
				"city",
				"search_text",
				"code",
				"product_count",
				"gmt",
				"updated_at"
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,NOW())
	`
	)

	t_time := time.Now()
	for _, data := range dates {
		wg := sync.WaitGroup{}
		wg.Add(1)

		go func(wg *sync.WaitGroup, data models.UpAirport) {
			defer wg.Done()
			data.Id = uuid.New().String()
			fmt.Printf("%+v", data)
			// fmt.Println(query)
			_, err = c.db.Exec(ctx, query,
				data.Id,
				data.Guid,
				data.Title,
				data.CountryId,
				data.CityId,
				data.Latitude,
				data.Longitude,
				data.Radius,
				data.Image,
				data.Address,
				data.TimezoneId,
				data.Country,
				data.City,
				data.SearchText,
				data.Code,
				data.ProductCount,
				data.Gmt,
			)
			if err != nil {
				panic(err)
			}
		}(&wg, data)
		// fmt.Println(data)
		wg.Wait()
	}
	fmt.Println(time.Since(t_time))
	return nil
}
