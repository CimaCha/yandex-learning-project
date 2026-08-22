package model

// Video — структура видео.
type Video struct {
	Id    string
	Title string
	Views int64
}

// limit — максимальное количество записей.
const limit = 20
