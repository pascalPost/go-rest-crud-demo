package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/swgui/v5emb"
	"log"
	"net/http"
)

func main() {
	// config mock
	c := NewConfig(true)

	// db mock
	db := NewDb()

	s := NewService(c, db)

	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	reflector.Spec.Info.
		WithTitle("Test API").
		WithVersion("1.2.3").
		WithDescription("Test API description.")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Get("/articles", s.GetArticles)
	getOp, _ := reflector.NewOperationContext(http.MethodGet, "/articles")
	getOp.AddRespStructure(new([]*Article))
	reflector.AddOperation(getOp)

	r.Get("/articles/{id}", s.GetArticle)

	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}

	r.Get("/api/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(reflector.Spec)
		if err != nil {
			if s.c.Dev() {
				http.Error(w, "Internal error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	})

	if c.Dev() {
		r.Handle("/api/docs", v5emb.New(
			reflector.Spec.Title(),
			"http://localhost:3000/api/openapi.json",
			"/api/docs",
		))
	}

	fmt.Println(string(schema))

	http.ListenAndServe(":3000", r)
}

type Service struct {
	db *Db
	c  *Config
}

func NewService(c *Config, db *Db) Service {
	return Service{db: db, c: c}
}

func (s *Service) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles := s.db.GetArticles()
	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (s *Service) GetArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := s.db.GetArticle(id)

	if err != nil {
		if s.c.Dev() {
			http.Error(w, "Internal error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		if s.c.Dev() {
			http.Error(w, "Internal error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
