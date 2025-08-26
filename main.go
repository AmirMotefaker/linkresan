package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
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
	// اتصال به پایگاه داده
	connStr := "user=postgres password=12345678 dbname=linkresan_db sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err) // اضافه شده
	}
	defer db.Close()

	router := gin.Default()

	// روت صفحه اصلی
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

	// روت برای ایجاد لینک کوتاه و ذخیره در دیتابیس
	router.POST("/shorten", func(c *gin.Context) {
		log.Println("Received a request to shorten a URL.") // اضافه شده
		originalURL := c.PostForm("url")
		log.Println("Original URL received:", originalURL) // اضافه شده

		shortCode := generateShortCode(urlLength)

		// درج لینک در پایگاه داده
		_, err := db.Exec("INSERT INTO links (short_code, original_url) VALUES ($1, $2)", shortCode, originalURL)
		if err != nil {
			log.Println("Error saving link to database:", err) // اضافه شده
			c.String(http.StatusInternalServerError, "Error saving link to database.")
			return
		}
		log.Println("Successfully saved link with short code:", shortCode) // اضافه شده

		shortURL := "http://localhost:8080/" + shortCode

		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	})

	// روت برای هدایت به لینک اصلی
	router.GET("/:shortCode", func(c *gin.Context) {
		shortCode := c.Param("shortCode")
		log.Println("Received a redirect request for short code:", shortCode) // اضافه شده

		var originalURL string
		err := db.QueryRow("SELECT original_url FROM links WHERE short_code = $1", shortCode).Scan(&originalURL)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("Short code not found in database:", shortCode) // اضافه شده
				c.String(http.StatusNotFound, "404 Not Found")
				return
			}
			log.Println("Error querying the database:", err) // اضافه شده
			c.String(http.StatusInternalServerError, "Internal Server Error.")
			return
		}
		log.Println("Redirecting to:", originalURL) // اضافه شده
		c.Redirect(http.StatusMovedPermanently, originalURL)
	})

	router.Run(":8080")
}