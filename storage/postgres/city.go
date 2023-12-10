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

type CityRepo struct {
	db *pgxpool.Pool
}

func NewCityRepo(db *pgxpool.Pool) *CityRepo {
	return &CityRepo{
		db: db,
	}
}

func (a *CityRepo) Create(ctx context.Context, req models.CreateCity) (*models.City, error) {

	var id = uuid.New().String()
	var query = `
	INSERT INTO city(
		"id",
		"guid",
		"title",
		"country_id",
		"city_code",
		"latitude",
		"longitude",
		"offset",
		"timezone_id",
		"country_name",
		updated_at
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW())
`
	// fmt.Println(query)
	_, err := a.db.Exec(ctx, query,
		id,
		req.Guid,
		req.Title,
		req.CountryId,
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		req.TimezoneId,
		req.CountryName,
	)
	if err != nil {
		return &models.City{}, err
	}
	return a.GetById(ctx, models.CityPrimaryKey{Id: id})
}

func (c *CityRepo) GetById(ctx context.Context, req models.CityPrimaryKey) (*models.City, error) {

	var (
		city  = models.City{}
		query = `
		SELECT 
			"id",
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			created_at,
			updated_at 
		FROM city WHERE id=$1`
	)
	var (
		Id          sql.NullString
		Guid        sql.NullString
		Title       sql.NullString
		CountryId   sql.NullString
		CityCode    sql.NullString
		Latitude    sql.NullString
		Longitude   sql.NullString
		Offset      sql.NullString
		TimezoneId  sql.NullString
		CountryName sql.NullString
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&Guid,
		&Title,
		&CountryId,
		&CityCode,
		&Latitude,
		&Longitude,
		&Offset,
		&TimezoneId,
		&CountryName,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	city = models.City{
		Id:          Id.String,
		Guid:        Guid.String,
		Title:       Title.String,
		CountryId:   CountryId.String,
		CityCode:    CityCode.String,
		Latitude:    Latitude.String,
		Longitude:   Longitude.String,
		Offset:      Offset.String,
		TimezoneId:  TimezoneId.String,
		CountryName: CountryName.String,
		CreatedAt:   CreatedAt.String,
		UpdatedAt:   UpdatedAt.String,
	}
	return &city, nil
}

func (c *CityRepo) GetList(ctx context.Context, req models.GetListCityRequest) (*models.GetListCityResponse, error) {
	var (
		citys  = models.GetListCityResponse{}
		where  = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"

		query = `
		SELECT 
			COUNT(*) OVER(),
			"id",
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			"created_at",
			"updated_at"
		FROM city
		`
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
	// fmt.Println(query)
	for rows.Next() {
		var (
			Id          sql.NullString
			Guid        sql.NullString
			Title       sql.NullString
			CountryId   sql.NullString
			CityCode    sql.NullString
			Latitude    sql.NullString
			Longitude   sql.NullString
			Offset      sql.NullString
			TimezoneId  sql.NullString
			CountryName sql.NullString
			CreatedAt   sql.NullString
			UpdatedAt   sql.NullString
		)
		err = rows.Scan(
			&citys.Count,
			&Id,
			&Guid,
			&Title,
			&CountryId,
			&CityCode,
			&Latitude,
			&Longitude,
			&Offset,
			&TimezoneId,
			&CountryName,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		citys.Cities = append(citys.Cities, models.City{
			Id:          Id.String,
			Guid:        Guid.String,
			Title:       Title.String,
			CountryId:   CountryId.String,
			CityCode:    CityCode.String,
			Latitude:    Latitude.String,
			Longitude:   Longitude.String,
			Offset:      Offset.String,
			TimezoneId:  TimezoneId.String,
			CountryName: CountryName.String,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		})
	}
	// fmt.Println(citys)
	defer rows.Close()
	return &citys, nil
}

func (c *CityRepo) Update(ctx context.Context, req models.UpdateCity) (*models.City, error) {
	var (
		query = `
			UPDATE city SET 
				guid = $2,
				title = $3,
				country_id = $4,
				city_code = $5,
				latitude = $6,
				longitude = $7,
				"offset" = $8,
				timezone_id = $9,
				country_name = $10,
				updated_at = NOW() 
			WHERE id = $1`
	)

	_, err := c.db.Exec(ctx, query,
		req.Id,
		req.Guid,
		req.Title,
		req.CountryId,
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		req.TimezoneId,
		req.CountryName,
	)

	// fmt.Println("*******************", query)
	if err != nil {
		return nil, err
	}

	return c.GetById(ctx, models.CityPrimaryKey{Id: req.Id})
}

func (c *CityRepo) Delete(ctx context.Context, req models.CityPrimaryKey) (string, error) {

	_, err := c.db.Exec(ctx, `DELETE FROM city WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *CityRepo) Upload(ctx context.Context, req models.UploadCity) error {
	defer os.Remove(req.DST)
	// fileName, err := helpers.Upload(&req.File, req.Base)
	// if err != nil {
	// 	return err
	// }

	dates, err := helpers.ReadCity(req.DST)
	if err != nil {
		return err
	}
	var (
		// str   string
		query = `
		INSERT INTO city(
					"id",
					"guid",
					"title",
					"country_id",
					"city_code",
					"latitude",
					"longitude",
					"offset",
					"timezone_id",
					"country_name",
					"updated_at"
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW())
`
	)

	wg := sync.WaitGroup{}
	t_time := time.Now()
	for _, data := range dates {

		wg.Add(1)
		go func(wg *sync.WaitGroup, data models.UpCity) {
			defer wg.Done()
			data.Id = uuid.New().String()
			// fmt.Printf("%+v", data)
			fmt.Println(query)
			_, _ = c.db.Exec(ctx, query,
				data.Id,
				data.Guid,
				data.Title,
				data.CountryId,
				data.CityCode,
				data.Latitude,
				data.Longitude,
				data.Offset,
				data.TimezoneId,
				data.CountryName,
			)
			// if err != nil {
			// 	return err
			// }
		}(&wg, data)
		wg.Wait()
	}
	fmt.Println(time.Since(t_time))
	return nil
}
