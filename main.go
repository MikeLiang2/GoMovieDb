package main

import (
	"godb/db"
	"log"
	"os"
)

func main() {
	dbPath := "imdb.db"

	if _, err := os.Stat(dbPath); err == nil {
		err := os.Remove(dbPath)
		if err != nil {
			log.Fatal("Failed to remove existing database:", err)
		}
		log.Println("Old database removed.")
	}

	database, err := db.InitDatabase(dbPath)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	defer database.Close()

	err = db.LoadMovies(database, "IMDb/IMDb-movies.csv")
	if err != nil {
		log.Fatal("Failed to load movies:", err)
	}

	err = db.LoadGenres(database, "IMDb/IMDb-movies_genres.csv")
	if err != nil {
		log.Fatal("Failed to load genres:", err)
	}

	_, err = database.Exec(`
    INSERT INTO my_collection (movie_id, my_rating, location, notes) VALUES
    (2, 9.5, 'External HDD', 'One of my favorite classics'),
    (5, 8.0, 'BluRay Shelf', 'A visual masterpiece'),
    (12, 7.5, 'Google Drive', 'Watched during college'),
    (18, 9.0, 'External HDD', 'Emotional and powerful'),
    (25, 6.5, 'DVD Box', 'Rewatchable but flawed'),
    (31, 8.8, 'NAS Storage', 'Sci-fi at its best'),
    (37, 7.2, 'Google Drive', 'Charming and nostalgic'),
    (42, 9.2, 'BluRay Shelf', 'Deep and meaningful'),
    (47, 8.4, 'External HDD', 'Great performances'),
    (53, 7.9, 'Cloud Folder', 'Feel-good movie'),
    (61, 8.7, 'BluRay Shelf', 'Classic thriller'),
    (66, 9.1, 'External SSD', 'Rewatched multiple times'),
    (70, 7.3, 'DVD Box', 'Slow burn but worth it'),
    (77, 8.9, 'USB Stick', 'Unexpectedly great'),
    (83, 9.0, 'NAS Storage', 'Strong script and cast');
	`)
	if err != nil {
		log.Println("Insert into my_collection failed:", err)
	}

	// top 10 genres by average movie rating
	err = db.QueryTopGenres(database, 10)
	if err != nil {
		log.Fatal("Query failed:", err)
	}

	// see all movies in my collection
	err = db.QueryMyCollection(database)
	if err != nil {
		log.Fatal("Collection query failed:", err)
	}
}
