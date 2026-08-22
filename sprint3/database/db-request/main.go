package main

import (
	"context"
	"database/sql"
	"fmt"
	"yandex-learning-project/sprint3/database/db-request/model"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "sprint3/database/data/video.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRowContext(context.Background(),
		"SELECT title, likes, comments_disabled "+
			"FROM videos ORDER BY likes DESC LIMIT 1")
	var (
		title  string
		likes  int
		comdis bool
	)
	// порядок переменных должен соответствовать порядку колонок в запросе
	err = row.Scan(&title, &likes, &comdis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s | %d | %t \r\n", title, likes, comdis)

	desc, err := getDesc(context.Background(), db, "0EbFotkXOiA")
	if err != nil {
		panic(err)
	}

	fmt.Printf(desc)

	videos, err := QueryVideos(context.Background(), db, 10)
	if err != nil {
		panic(err)
	}

	for _, video := range videos {
		fmt.Printf("%s, %s, %d", video.Id, video.Title, video.Views)
	}
}

func getDesc(ctx context.Context, db *sql.DB, id string) (string, error) {
	row := db.QueryRowContext(ctx,
		"SELECT description FROM videos WHERE video_id = ?", id)
	var desc sql.NullString

	err := row.Scan(&desc)
	if err != nil {
		return "", err
	}
	if desc.Valid {
		return desc.String, nil
	}
	return "-----", nil
}

func QueryVideos(ctx context.Context, db *sql.DB, limit int) ([]model.Video, error) {
	videos := make([]model.Video, 0, limit)

	rows, err := db.QueryContext(ctx, "SELECT video_id, title, views from videos ORDER BY views LIMIT ?", limit)
	if err != nil {
		return nil, err
	}

	// обязательно закрываем перед возвратом функции
	defer rows.Close()

	// пробегаем по всем записям
	for rows.Next() {
		var v model.Video
		err = rows.Scan(&v.Id, &v.Title, &v.Views)
		if err != nil {
			return nil, err
		}

		videos = append(videos, v)
	}

	// проверяем на ошибки
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return videos, nil
}
