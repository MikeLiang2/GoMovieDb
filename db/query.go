package db

import (
	"database/sql"
	"fmt"
)

func QueryTopGenres(db *sql.DB, limit int) error {
	query := `
        SELECT g.genre, AVG(m.rating) AS avg_rating
        FROM movies m
        JOIN genres g ON m.id = g.movie_id
        WHERE m.rating IS NOT NULL
        GROUP BY g.genre
        ORDER BY avg_rating DESC
        LIMIT ?;
    `

	rows, err := db.Query(query, limit)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Top Genres by Average Rating:")
	for rows.Next() {
		var genre string
		var avg float64
		err := rows.Scan(&genre, &avg)
		if err != nil {
			return err
		}
		fmt.Printf(" - %-15s : %.2f\n", genre, avg)
	}

	return nil
}

func QueryMyCollection(db *sql.DB) error {
	query := `
        SELECT m.title, m.year, c.my_rating, c.location, c.notes
        FROM my_collection c
        JOIN movies m ON c.movie_id = m.id
        ORDER BY c.my_rating DESC;
    `

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("\n My Movie Collection:")
	for rows.Next() {
		var title string
		var year int
		var rating float64
		var location, notes string

		err := rows.Scan(&title, &year, &rating, &location, &notes)
		if err != nil {
			return err
		}

		fmt.Printf(" - %s (%d) | My Rating: %.1f | Location: %s | Notes: %s\n",
			title, year, rating, location, notes)
	}

	return nil
}
