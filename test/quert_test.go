package test

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
	// 替换为你的 module 名
)

func TestTopGenresQuery(t *testing.T) {
	db, _ := sql.Open("sqlite", ":memory:")
	defer db.Close()

	db.Exec(`CREATE TABLE movies (id INTEGER PRIMARY KEY, title TEXT, year INTEGER, rating REAL)`)
	db.Exec(`CREATE TABLE genres (movie_id INTEGER, genre TEXT)`)

	db.Exec(`INSERT INTO movies VALUES (1, 'Movie A', 2020, 9.0), (2, 'Movie B', 2020, 8.0)`)
	db.Exec(`INSERT INTO genres VALUES (1, 'Drama'), (2, 'Drama')`)

	rows, err := db.Query(`
        SELECT g.genre, AVG(m.rating)
        FROM movies m
        JOIN genres g ON m.id = g.movie_id
        GROUP BY g.genre;
    `)
	if err != nil {
		t.Fatal("Query failed:", err)
	}

	count := 0
	for rows.Next() {
		count++
	}
	if count == 0 {
		t.Error("Expected at least one result row")
	}
}
