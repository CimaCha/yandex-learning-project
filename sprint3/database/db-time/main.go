package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Trend struct {
	T     time.Time
	Count int
}

func TrendingCount(db *sql.DB) ([]Trend, error) {
	trends := make([]Trend, 0)
	date := new(string)

	rows, err := db.Query("SELECT trending_date, COUNT(trending_date) FROM videos GROUP BY trending_date")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// порядок переменных должен соответствовать порядку колонок в запросе
	for rows.Next() {
		trend := Trend{}
		// все теги должны автоматически преобразоваться в слайс v.Tags
		err = rows.Scan(date, &trend.Count)
		if err != nil {
			return nil, err
		}
		var t time.Time
		if t, err = time.Parse("06.02.01", *date); err != nil {
			return nil, err
		}
		trend.T = t
		trends = append(trends, trend)
	}

	return trends, nil
}

func main() {
	db, err := sql.Open("sqlite", "sprint3/database/data/video.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	trends, err := TrendingCount(db)
	for _, t := range trends {
		fmt.Printf("Trending date: %s, Count: %d\n", t.T, t.Count)
	}
}
