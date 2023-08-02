package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	// "os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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
	// cfg := mysql.Config{
	// 	User: os.Getenv("MYSQL_USER"),
	// 	Passwd: os.Getenv("MYSQL_PASSWORD"),
	// 	Net: "tcp",
	// 	Addr: ":3306",
	// 	DBName: os.Getenv("MYSQL_DATABASE"),
	// 	AllowNativePasswords: true,
	// }
	var err error
	
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)  // for example /home/user

	// Get a database handle

	db, err = sql.Open("sqlite3", "./records.db")
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to ensure connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")
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
	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		fmt.Printf("%v\n", err)
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
	g.JSON(http.StatusOK, albums)
}

func getAlbum(g *gin.Context) {
	// Capture the variable passed in through the route in a variable
	id := g.Param("id")

	// Declare a slice called "albums" of type Album to hold the results
	var alb Album

	// Query the database (DB.QueryRow dos not return an error)
	row := db.QueryRow("SELECT * FROM albums WHERE id = ?", id)

	// Use Rows.Scan() to insert the values from the row into the "album" variable
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		// If no rows are found, return an message
		if err == sql.ErrNoRows {
			g.JSON(http.StatusOK, []Album{})
			return
		}
	}

	// Send "album" as JSON
	g.JSON(http.StatusOK, alb)
}