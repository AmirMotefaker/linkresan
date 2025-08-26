package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	urlLength   = 6
)

func generateShortCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	// خواندن اطلاعات اتصال از متغیر DATABASE_URL که توسط Render تزریق می‌شود
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set.")
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		html := `
		<!DOCTYPE html>
		<html lang="fa" dir="rtl">
		<head>
			<meta charset="UTF-8">
			<title>Linkresan.ir</title>
		</head>
		<body>
			<h1>به لینک‌رسان خوش آمدید</h1>
			<form action="/shorten" method="post">
				<label for="url">لینک بلند:</label><br>
				<input type="url" id="url" name="url" placeholder="مثال: https://www.google.com" size="50" required><br><br>
				<input type="submit" value="کوتاه کن!">
			</form>
		</body>
		</html>
		`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})

	router.POST("/shorten", func(c *gin.Context) {
		originalURL := c.PostForm("url")
		shortCode := generateShortCode(urlLength)

		_, err := db.Exec("INSERT INTO links (short_code, original_url) VALUES ($1, $2)", shortCode, originalURL)
		if err != nil {
			log.Println("Error saving link to database:", err)
			c.String(http.StatusInternalServerError, "Error saving link to database.")
			return
		}

		log.Println("Successfully created link:", shortCode, "for URL:", originalURL)
		shortURL := c.Request.Host + "/" + shortCode
		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	})

	router.GET("/:shortCode", func(c *gin.Context) {
		shortCode := c.Param("shortCode")

		var originalURL string
		err := db.QueryRow("SELECT original_url FROM links WHERE short_code = $1", shortCode).Scan(&originalURL)
		if err != nil {
			if err == sql.ErrNoRows {
				c.String(http.StatusNotFound, "404 Not Found")
				return
			}
			c.String(http.StatusInternalServerError, "Internal Server Error.")
			return
		}

		c.Redirect(http.StatusMovedPermanently, originalURL)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Starting server on port:", port)
	router.Run(":" + port)
}