package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var links = make(map[string]string)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	urlLength   = 6
)

// تابع برای تولید کد کوتاه و تصادفی
func generateShortCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	router := gin.Default()

	// روت صفحه اصلی برای نمایش فرم
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
			<p>لینک‌های بلند خود را به آدرس‌های کوتاه و زیبا تبدیل کنید.</p>
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

	// روت برای دریافت لینک بلند و ایجاد لینک کوتاه
	router.POST("/shorten", func(c *gin.Context) {
		originalURL := c.PostForm("url")
		shortCode := generateShortCode(urlLength)

		// ذخیره لینک در حافظه (آینده به دیتابیس تغییر می‌کند)
		links[shortCode] = originalURL

		shortURL := "http://localhost:8080/" + shortCode

		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	})

	// روت برای ریدایرکت (هدایت) به لینک اصلی
	router.GET("/:shortCode", func(c *gin.Context) {
		shortCode := c.Param("shortCode")
		originalURL, ok := links[shortCode]
		if !ok {
			c.String(http.StatusNotFound, "404 Not Found")
			return
		}
		c.Redirect(http.StatusMovedPermanently, originalURL)
	})

	router.Run(":8080")
}