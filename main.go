package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/swgui/v5emb"
	"log"
	"net/http"
	"os"
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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	r.Get("/articles", s.GetArticles)
	getOp, _ := reflector.NewOperationContext(http.MethodGet, "/articles")
	getOp.AddRespStructure(new([]*Article))
	err := reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	r.Get("/articles/{id}", s.GetArticle)

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
		schema, err := reflector.Spec.MarshalYAML()
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile("./api/openapi.yaml", schema, 0666)
		if err != nil {
			log.Fatal(err)
		}

		r.Mount("/api/docs", v5emb.New(
			reflector.Spec.Title(),
			"/api/openapi.json",
			"/api/docs/",
		))
		println("docs at http://localhost:3000/api/docs")
	}

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
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
