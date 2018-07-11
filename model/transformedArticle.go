package model

import (
	"strconv"
	"time"
)

type TransformedArticle struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Date     string `json:"data"`
	Username string `json:"username"`
}

func NewTransformedArticle(article Article) *TransformedArticle {
	date := parseDate(article.UpdatedAt)
	res := TransformedArticle{
		Id:       article.ID,
		Title:    article.Title,
		Content:  article.Content,
		Date:     date,
		Username: article.Username,
	}

	return &res
}

func parseDate(data time.Time) string {
	return strconv.Itoa(data.Day()) + " " + data.Month().String() + " " + strconv.Itoa(data.Year()) + " " + strconv.Itoa(data.Hour()) + ":" +
		strconv.Itoa(data.Minute())
}
