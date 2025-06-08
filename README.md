# Go IMDb Movie Database Application

This is a sample personal movie database application built in **Go** using a **pure-Go implementation of SQLite** (`modernc.org/sqlite`). It loads real IMDb-like movie data and allows local storage, querying, and management of personal movie collections.

## Dataset

The application uses six comma-delimited text files provided by Northwestern University, based on open IMDb data:
https://arch.library.northwestern.edu/concern/datasets/3484zh40n?locale=en

In this project, we use:
- `IMDb-movies.csv`: general movie data
- `IMDb-movies_genres.csv`: genre mappings

## Database Schema

Implemented in SQLite with the following tables:

### `movies`
| Column  | Type   | Description          |
|---------|--------|----------------------|
| id      | INT    | Movie ID (primary)   |
| title   | TEXT   | Movie title          |
| year    | INT    | Year of release      |
| rating  | REAL   | IMDb-style rating    |

### `genres`
| Column     | Type   | Description                    |
|------------|--------|--------------------------------|
| movie_id   | INT    | FK to movies.id                |
| genre      | TEXT   | Genre of the movie             |

### `my_collection`
| Column     | Type   | Description                    |
|------------|--------|--------------------------------|
| movie_id   | INT    | FK to movies.id                |
| my_rating  | REAL   | My personal rating             |
| location   | TEXT   | Where I store the movie        |
| notes      | TEXT   | Personal notes (e.g., "classic")|

## How to Run

### Prerequisites
- Go 1.18+ (no need to install SQLite)
- Clone this repository
- Place CSV files into a `IMDb/` folder

### Run the app  
go run main.go

(Current main function will:  
Remove any existing imdb.db  
Create database schema  
Load movies and genres from CSV  
Insert 15 example movies into my_collection  
Execute SQL queries and print results  
)  

## Sample Query Results
## Top Genres by Average Rating
Top Genres by Average Rating:  
 - Film-Noir       : 6.70  
 - Animation       : 6.56  
 - Documentary     : 6.50  
 - Adult           : 6.43  
 - Music           : 6.42  

## My Movie Collection
 - $ (1971) | My Rating: 9.5 | Location: External HDD | Notes: One of my favorite classics  
 - '60s Pop Rock Reunion (2004) | My Rating: 9.2 | Location: BluRay Shelf | Notes: Deep and meaningful  

## How the SQLite Relational Database Was Set Up
To set up the SQLite relational database, I used the pure Go library modernc.org/sqlite, avoiding cgo and external dependencies. The database is initialized automatically on runtime. Schema definitions were created using standard SQL CREATE TABLE statements for movies, genres, and a custom my_collection table. The whole process can be found in schema.go file and have no difference with create a database in MYSQL. The application parses the raw CSV files and uses prepared statements with batched transactions to efficiently populate the database. 

## Adding a Personal Movie Collection Table
To make the database more relevant for personal use, I added a my_collection table with the following fields:
- movie_id: foreign key referencing the primary movies table
- my_rating: a personal numeric rating
- location: where the movie is stored (e.g., BluRay shelf, external SSD, NAS)
- notes: any personal annotation or reason for collecting the movie
For this version, this table needs to be populated manually or programmatically to record which movies I own, where I store them, and how much I like them. Insertion of data is in main.go file.  

The main purpose of this local movie database is to provide a lightweight, fully controllable archive of movies that I have watched or collected. It allows me to track:
What I own, What I rated highly, Where my files are stored and What genres or directors I tend to favor. Unlike IMDb, this database can be customized, private, and built around personal needs or tastes. For instance, this project shows a simple example with a personal movie collection table, which provides some simple options like maintaining a list of movie I watched.  
  
Unlike IMDb, which offers global, crowdsourced data, this application focuses on personal movie management. It allows users to record their own ratings, notes, and storage locations—features IMDb does not support. In addition, the app is fully local and ad-free, making it fast, private, and ideal for offline use. It's not just a read-only reference, but a customizable and programmable tool tailored to individual preferences.

### Possible User Interactions and my future thought
Beyond direct SQL queries, a practical Go movie application could offer: A CLI or web-based interface to search by genre, year, rating, or storage location. A “Top Picks” list based on personal ratings. The ability to filter movies by availability (e.g., "Show all movies I have stored on my Google Drive") and feature to export or backup selected movies into a report or JSON. Some advanced feature (not supported but possible) includes read the comments, note of the movie collection table, using NLP plan deeplearning algorithm to learn my own preferences and provide recommendations when opening or searching for movies. I think this requires more attributes or features from each movies to make things interesting, such as some descriptions of the movie plot and the location where the story takes place. In future version, I could begain with some simple algorithm I'm familar with like K-mean clustering. This will make this application truly different from IMdm.  


## GenAI Tools Used
This project was built with support from ChatGPT-4o, which helped:
- Design the schema and optimize table structure
- Handle malformed CSV input
- Implement batch insert with SQLite transactions
- Write JOIN queries and CLI output formatting
- Draft this README structure  
https://chatgpt.com/share/6844f74f-0cbc-8008-88d0-2bb3277e6002
