package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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

type Topic struct {
	ID string `json:"id"`
	Topic string `json:"topic"`
}

type Verse struct {
	Reference string `json:"reference"`
	Verse string `json:"verse"`
	Supports bool `json:"supports"`
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

	db, err = sql.Open("sqlite3", "./verses.db")
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

	// allow CORS
	r.Use(CORSMiddleware())

	// Register routes
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbum)
	r.GET("/topics", getTopics)
	r.GET("/topic/:topic_id", getTopic)

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

func getTopics(g *gin.Context) {
	// Declare a slice called "topics" of type Album to hold the results
	var topics[]Topic

	// Query the database
	rows, err := db.Query("SELECT * FROM topics")
	if err != nil {
		fmt.Printf("%v\n", err)
		g.AbortWithStatus(500)
		return
	}

	// Loop through the returned rows and add each to the "topics" slice
	for rows.Next() {
		// Declare a variable of type Topic to hold the row
		var topic Topic

		// Use Rows.Scan() to insert values into "alb"
		if err := rows.Scan(&topic.ID, &topic.Topic); err != nil {
			g.AbortWithStatus(500)
			return
		}

		// Add the album from the row to "topics"
		topics = append(topics, topic)
	}

	// Send "topics" as JSON
	g.JSON(http.StatusOK, topics)
}

func getTopic(g *gin.Context) {
	topicID := g.Param("topic_id")

	var verses[]Verse

	// Query the database
	rows, err := db.Query(`
		SELECT
			bible_verses.verse_reference,
			bible_verses.verse_text,
			map.supports
		FROM
			map
		INNER JOIN
			bible_verses ON map.verse_id = bible_verses.verse_id
		WHERE
			map.topic_id = ?;
		`, topicID)
	if err != nil {
		fmt.Printf("%v\n", err)
		g.AbortWithStatus(500)
		return
	}

	// Loop through the returned rows and add each to the "verses" slice
	for rows.Next() {
		// Declare a variable of type Verse to hold the row
		var verse Verse

		// Use Rows.Scan() to insert values into "verse"
		if err := rows.Scan(&verse.Reference, &verse.Verse, &verse.Supports); err != nil {
			g.AbortWithStatus(500)
			return
		}

		// Add the verse from the row to "verses"
		verses = append(verses, verse)
	}

	// Send "albums" as JSON
	g.JSON(http.StatusOK, verses)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}