package scraper

import (
	"encoding/csv"
	"os"
)

func SaveToCSV(blog Blog) {
	file, err := os.OpenFile("blogs.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // 0644 -> unix file permissions
	if err != nil {
		panic(err) // handle this gracefully next time
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{blog.Title, blog.Date, blog.Author})
}
