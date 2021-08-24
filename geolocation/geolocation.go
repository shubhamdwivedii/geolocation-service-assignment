package geolocation

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Geolocation struct {
	IP        string  `json:"ip"`
	CCode     string  `json:"country_code"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	MValue    int64   `json:"mystery_value"`
}

type GeoData struct {
	sync.Mutex
	db *sql.DB
}

func InitGeoData() (*GeoData, error) {
	// DB_URL := os.Getenv("DB_URL")
	DB_URL := "root:hesoyam@tcp(127.0.0.1:3306)/geolocation"

	URLS := strings.Split(DB_URL, "/")
	CONNECTION_URL := URLS[0]
	DB_NAME := URLS[1]

	// Add sql driver by: "go get github.com/go-sql-driver/mysql"
	fmt.Println("Drivers:", sql.Drivers())

	db, err := sql.Open("mysql", CONNECTION_URL+"/")
	if err != nil {
		log.Fatal("Unable to open connection to DB", err.Error())
		return nil, err
	} else {
		fmt.Println("Connected to DB...")
	}

	_, err = db.Exec("USE " + DB_NAME) // Replace with CREATE IF NOT EXISTS later
	if err != nil {
		fmt.Println("DB Not Exists", DB_NAME, err.Error())
		_, err = db.Exec("CREATE DATABASE " + DB_NAME)
		if err != nil {
			log.Fatal("Error: Creating Database", err.Error())
			return nil, err
		} else {
			fmt.Println("Database Created Successfully...")
			_, err = db.Exec("USE " + DB_NAME)
			if err != nil {
				log.Fatal("Error: Using Newly Creating DB", err.Error())
				return nil, err
			}
		}
	} else {
		fmt.Println("Database found...")
	}

	return &GeoData{
		db: db,
	}, checkTable("geolocations", db)
}

func checkTable(table string, db *sql.DB) error {
	results, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		fmt.Println("Table:", table, "Does Not Exists")

		create_table, err := db.Prepare("CREATE TABLE " + table + "(ip varchar(20), ccode varchar(3), country varchar(20), city varchar(20), latitude double, longitude double, mystery bigint, PRIMARY KEY (ip))")
		if err != nil {
			fmt.Println("Error: Creating Table Statement", err.Error())
			return err
		} else {
			_, err := create_table.Exec()
			if err != nil {
				fmt.Println("Error: Creating Table", err.Error())
				return err
			} else {
				fmt.Println("Table Created Successfully...")
				return nil
			}
		}
	} else {
		fmt.Println("Table", table, "Found...")
		results.Close()
		return nil
	}
}

func (geodata *GeoData) AddGeoData(location *Geolocation) error {
	if location.IP == "" || location.CCode == "" || location.Country == "" || location.City == "" {
		return errors.New("Invalid or empty data")
	}
	if location.Latitude == 0 || location.Longitude == 0 || location.MValue == 0 {
		return errors.New("Invalid or zero data")
	}

	defer geodata.Unlock()
	geodata.Lock()

	add_location, err := geodata.db.Prepare("INSERT INTO geolocations (ip, ccode, country, city, latitude, longitude, mystery) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = add_location.Exec(location.IP, location.CCode, location.Country, location.City, location.Latitude, location.Longitude, location.MValue)
	if err != nil {
		return err
	}
	return nil
}

func (geodata *GeoData) GetGeoData(ip string) (Geolocation, error) {
	defer geodata.Unlock()
	geodata.Lock()

	var geolocation Geolocation
	qry := fmt.Sprintf("SELECT * from geolocations where ip = '%v'", ip)
	err := geodata.db.QueryRow(qry).Scan(
		&geolocation.IP,
		&geolocation.CCode,
		&geolocation.Country,
		&geolocation.City,
		&geolocation.Latitude,
		&geolocation.Longitude,
		&geolocation.MValue,
	)
	if err != nil {
		return geolocation, err
	}
	return geolocation, nil
}

func ReadCSV(geodata *GeoData) {
	// Opening file
	pwd, _ := os.Getwd()
	fmt.Println("PWD", pwd)
	recordFile, err := os.Open(pwd + "/geolocation/sample.csv")
	// Path might not work on windows (use git-bash)

	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Initialize the reader
	reader := csv.NewReader(recordFile)

	// Read row by row
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println("Headers", header)
	// Add check for header ???

	for i := 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Println("Error:", err.Error())
		}
		var geoloc Geolocation
		geoloc.IP = record[0]
		geoloc.CCode = record[1]
		geoloc.Country = record[2]
		geoloc.City = record[3]
		geoloc.Latitude, err = strconv.ParseFloat(record[4], 64)
		if err != nil {
			fmt.Println("Error: parsing latitude", err.Error())
		}
		geoloc.Longitude, err = strconv.ParseFloat(record[5], 64)
		if err != nil {
			fmt.Println("Error: parsing latitude", err.Error())
		}
		geoloc.MValue, err = strconv.ParseInt(record[6], 10, 64)
		if err != nil {
			fmt.Println("Error: parsing latitude", err.Error())
		}
		fmt.Println("Geoloc:", geoloc)
		err = geodata.AddGeoData(&geoloc)
		if err != nil {
			fmt.Println("Error adding geo data", err.Error())
		}
		fmt.Println("RECORD:", record)
		fmt.Println("+++++++++++++++")
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("Error encountered closing file", err.Error())
	}
}
