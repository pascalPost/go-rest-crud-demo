package main

import "errors"

type Db struct {
	articles []*Article
}

func NewDb() *Db {
	return &Db{
		articles: articles,
	}
}

// fixture data
var articles = []*Article{
	{ID: "1", Title: "Hi", Slug: "hi"},
	{ID: "2", Title: "sup", Slug: "sup"},
	{ID: "3", Title: "alo", Slug: "alo"},
	{ID: "4", Title: "bonjour", Slug: "bonjour"},
	{ID: "5", Title: "whats up", Slug: "whats-up"},
}

func (d *Db) GetArticles() []*Article {
	return articles
}

func (d *Db) GetArticle(id string) (*Article, error) {
	println(id)

	for _, a := range articles {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("article not found.")
}
