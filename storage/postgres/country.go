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

type CountryRepo struct {
	db *pgxpool.Pool
}

func NewCountryRepo(db *pgxpool.Pool) *CountryRepo {
	return &CountryRepo{
		db: db,
	}
}

func (a *CountryRepo) Create(ctx context.Context, req models.CreateCountry) (*models.Country, error) {
	var id = uuid.New().String()
	var query = `
	INSERT INTO country(
			"id",
			"guid",
			"title",
			"code",
			"continent",
			"updated_at"
) VALUES ($1,$2,$3,$4,$5,NOW())
`
	// fmt.Println(query, req)
	_, err := a.db.Exec(ctx, query,
		id,
		req.Guid,
		req.Title,
		req.Code,
		req.Continent,
	)
	if err != nil {
		return &models.Country{}, err
	}
	return a.GetById(ctx, models.CountryPrimaryKey{Id: id})
}

func (c *CountryRepo) GetById(ctx context.Context, req models.CountryPrimaryKey) (*models.Country, error) {

	var (
		country = models.Country{}
		query   = `
		SELECT 
			"id",
			"guid",
			"title",
			"code",
			"continent",
			"created_at",
			"updated_at"
		FROM country WHERE id=$1`
	)
	var (
		Id        sql.NullString
		Guid      sql.NullString
		Title     sql.NullString
		Code      sql.NullString
		Continent sql.NullString
		CreatedAt sql.NullString
		UpdatedAt sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&Guid,
		&Title,
		&Code,
		&Continent,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	country = models.Country{
		Id:        Id.String,
		Guid:      Guid.String,
		Title:     Title.String,
		Code:      Code.String,
		Continent: Continent.String,
		CreatedAt: CreatedAt.String,
		UpdatedAt: UpdatedAt.String,
	}
	return &country, nil
}

func (c *CountryRepo) GetList(ctx context.Context, req models.GetListCountryRequest) (*models.GetListCountryResponse, error) {
	var (
		countrys = models.GetListCountryResponse{}
		where    = " WHERE TRUE "
		offset   = " OFFSET 0"
		limit    = " LIMIT 10"
		query    = `
		SELECT 
			COUNT(*) OVER(),
			"id",
			"guid",
			"title",
			"code",
			"continent",
			"created_at",
			"updated_at"
		FROM country
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
	fmt.Println(query)
	for rows.Next() {
		var (
			Id        sql.NullString
			Guid      sql.NullString
			Title     sql.NullString
			Code      sql.NullString
			Continent sql.NullString
			CreatedAt sql.NullString
			UpdatedAt sql.NullString
		)

		rows.Scan(
			&countrys.Count,
			&Id,
			&Guid,
			&Title,
			&Code,
			&Continent,
			&CreatedAt,
			&UpdatedAt,
		)
		countrys.Countries = append(countrys.Countries, models.Country{
			Id:        Id.String,
			Guid:      Guid.String,
			Title:     Title.String,
			Code:      Code.String,
			Continent: Continent.String,
			CreatedAt: CreatedAt.String,
			UpdatedAt: UpdatedAt.String,
		})
	}
	rows.Close()
	return &countrys, nil
}

func (c *CountryRepo) Update(ctx context.Context, req models.UpdateCountry) (*models.Country, error) {
	var (
		query = `
			UPDATE country SET 
				"guid" = $2,
				"title" = $3,
				"code" = $4,
				"continent" = $5,
				updated_at = NOW() 
			WHERE id = $1`
	)
	_, err := c.db.Exec(ctx, query,
		req.Id,
		req.Guid,
		req.Title,
		req.Code,
		req.Continent,
	)
	if err != nil {
		return nil, err
	}

	return c.GetById(ctx, models.CountryPrimaryKey{Id: req.Id})
}

func (c *CountryRepo) Delete(ctx context.Context, req models.CountryPrimaryKey) (string, error) {

	_, err := c.db.Exec(ctx, `DELETE FROM country WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *CountryRepo) Upload(ctx context.Context, req models.UploadCountry) error {
	defer os.Remove(req.DST)

	// fileName, err := helpers.Upload(&req.File, req.Base)
	// if err != nil {
	// 	return err
	// }

	dates, err := helpers.ReadCountry(req.DST)
	if err != nil {
		return err
	}
	var (
		// str   string
		query = `
		INSERT INTO country(
					"id",
					"guid",
					"title",
					"code",
					"continent",
					"updated_at"
) VALUES ($1,$2,$3,$4,$5,NOW())
`
	)

	wg := sync.WaitGroup{}
	t_time := time.Now()
	for _, data := range dates {

		wg.Add(1)
		go func(wg *sync.WaitGroup, data models.UpCountry) {
			defer wg.Done()
			data.Id = uuid.New().String()
			// fmt.Printf("%+v", data)
			//  fmt.Println(query)
			_, _ = c.db.Exec(ctx, query,
				data.Id,
				data.Guid,
				data.Title,
				data.Code,
				data.Continent,
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
