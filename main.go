package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float32 `json:"price"`
}

func init() {
	// Capture connection properties
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "recordings",
		AllowNativePasswords: true,
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to ensure connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func main() {
	// Create a new router
	r := gin.Default()

	// Register routes
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbum)

	// Run the server
	r.Run(":8080")
}

func getAlbums(g *gin.Context) {
	// Declare a slice called "albums" of type Album to hold the results
	var albums[]Album

	// Query the database
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		g.AbortWithStatus(500)
		return
	}

	// Loop through the returned rows and add each to the "albums" slice
	for rows.Next() {
		// Declare a variable of type Album to hold the row
		var alb Album

		// Use Rows.Scan() to insert values into "alb"
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			g.AbortWithStatus(500)
			return
		}

		// Add the album from the row to "albums"
		albums = append(albums, alb)
	}

	// Send "albums" as JSON
	g.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(g *gin.Context) {
	// Capture the variable passed in through the route in a variable
	id := g.Param("id")

	// Declare a slice called "albums" of type Album to hold the results
	var alb Album

	// Query the database (DB.QueryRow dos not return an error)
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	// Use Rows.Scan() to insert the values from the row into the "album" variable
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		// If no rows are found, return an message
		if err == sql.ErrNoRows {
			g.IndentedJSON(http.StatusOK, []Album{})
			return
		}
	}

	// Send "album" as JSON
	g.IndentedJSON(http.StatusOK, alb)
}