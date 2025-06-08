package db

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadMovies(db *sql.DB, filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	_, _ = reader.Read() // skip header

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO movies (id, title, year, rating) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Skipping bad line:", err)
			continue
		}

		id, _ := strconv.Atoi(record[0])
		title := strings.TrimSpace(record[1])
		year, _ := strconv.Atoi(record[2])

		var rating sql.NullFloat64
		if record[3] != "NULL" {
			f, _ := strconv.ParseFloat(record[3], 64)
			rating = sql.NullFloat64{Float64: f, Valid: true}
		}

		_, err = stmt.Exec(id, title, year, rating)
		if err != nil {
			log.Println("Insert failed:", err)
		}

		count++
		if count%10000 == 0 {
			log.Printf("Inserted %d records...", count)
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	log.Println("Movies loaded.")
	return nil
}

func LoadGenres(db *sql.DB, filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	_, _ = reader.Read() // skip header

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO genres (movie_id, genre) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Skipping bad line:", err)
			continue
		}

		movieID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println("Invalid movie ID:", record[0])
			continue
		}

		genre := strings.TrimSpace(record[1])

		_, err = stmt.Exec(movieID, genre)
		if err != nil {
			log.Println("Insert genre failed:", err)
		}

		count++
		if count%10000 == 0 {
			log.Printf("Inserted %d genres...", count)
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	log.Println("Genres loaded.")
	return nil
}
