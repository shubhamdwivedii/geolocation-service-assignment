package storage

import (
	"database/sql"
	"log"
	"regexp"

	// go get "github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
)

type SQLStorage struct {
	db *sql.DB
}

func NewSQLStorage(connection string) (Storage, error) {
	var err error
	s := new(SQLStorage)
	s.db, err = initDb(connection)
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
	return db, db.Ping()
}

func validateIP(ip string) bool {
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	return re.MatchString(ip)
}

func (s *SQLStorage) AddGeodata(gloc Geolocation) error {
	err := ValidateGeolocation(gloc)
	if err != nil {
		return err
	}

	query := sq.Insert("geolocation").
		Columns("ip", "ccode", "country", "city", "latitude", "longitude", "mystery").
		Values(gloc.IP, gloc.CCode, gloc.Country, gloc.City, gloc.Latitude, gloc.Longitude, gloc.MValue)

	_, err = query.RunWith(s.db).Exec()

	if err != nil {
		return err
	}
	return nil
}

func (s *SQLStorage) GetGeodata(ip string) (*Geolocation, error) {
	var gloc Geolocation

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
