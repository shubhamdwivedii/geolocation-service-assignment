package sql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"

	// go get "github.com/Masterminds/squirrel"
	"github.com/shubhamdwivedii/geolocation-service-assignment/models"
)

// Move to models ??
type Storage struct {
	sync.Mutex
	db *sql.DB
}

func NewStorage() (*Storage, error) {
	var err error
	s := new(Storage)

	// DB_URL := os.Getenv("DB_URL")
	DB_URL := "root:hesoyam@tcp(127.0.0.1:3306)/geolocation"
	s.db, err = initDb(DB_URL)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func initDb(connection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connection)

	if err != nil {
		log.Println("Unable to open connection to DB", err.Error())
		return nil, err
	}
	log.Println("Connected to DB successfully...")
	return db, nil
}

func (s *Storage) AddGeodata(gloc models.Geolocation) error {
	// check if all fields of gloc are valid (not empty)
	defer s.Unlock()
	s.Lock()

	query := sq.Insert("geolocation").
		Columns("ip", "ccode", "country", "city", "latitude", "longitude", "mystery").
		Values(gloc.IP, gloc.CCode, gloc.Country, gloc.City, gloc.Latitude, gloc.Longitude, gloc.MValue)

	_, err := query.RunWith(s.db).Exec()

	qry, _, _ := query.ToSql()
	fmt.Println("query >>>>", qry)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetGeodata(ip string) (*models.Geolocation, error) {
	defer s.Unlock()
	s.Lock()

	var gloc models.Geolocation

	query, args, err := sq.Select("*").From("geolocation").Where(sq.Eq{"ip": ip}).ToSql()

	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(query, args...).Scan(
		&gloc.IP,
		&gloc.CCode,
		&gloc.Country,
		&gloc.City,
		&gloc.Latitude,
		&gloc.Longitude,
		&gloc.MValue,
	)

	if err != nil {
		return nil, err
	}

	return &gloc, nil
}
