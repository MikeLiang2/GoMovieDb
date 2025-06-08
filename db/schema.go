package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	schema := `
    CREATE TABLE IF NOT EXISTS movies (
        id INTEGER PRIMARY KEY,
        title TEXT,
        year INTEGER,
        rating REAL
    );

    CREATE TABLE IF NOT EXISTS genres (
        movie_id INTEGER,
        genre TEXT,
        FOREIGN KEY (movie_id) REFERENCES movies(id)
    );

	CREATE TABLE IF NOT EXISTS my_collection (
    	movie_id INTEGER,
    	my_rating REAL,
    	location TEXT,
    	notes TEXT,
    	FOREIGN KEY (movie_id) REFERENCES movies(id)
	);
    `

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	log.Println("Database schema created.")
	return db, nil
}
