package app

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jsmvaldivia/shorty/internal/web"
	"net/http"
)

type ShortenerController struct {
	Service ShortenerService
}

func (rs ShortenerController) GetPath() string {
	return "/"
}

func (rs ShortenerController) GetRoutes() http.Handler {
	router := chi.NewRouter()
	router.Post("/", rs.CreateShortUrl)
	return router
}

func (rs ShortenerController) CreateShortUrl(w http.ResponseWriter, req *http.Request) {
	var sur ShorUrlRequest
	var err error
	if err = json.NewDecoder(req.Body).Decode(&sur); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortUrl, err := rs.Service.CreateShortUrl(sur.FullUrl)

	w.WriteHeader(http.StatusCreated)

	render.JSON(w, req, NewUrlResponse(shortUrl))
}

type ShorUrlRequest struct {
	FullUrl string `json:"fullUrl"`
}

type ShorUrlResponse struct {
	ShortUrl string `json:"shortUrl"`
}

func NewUrlResponse(shortUrl string) *ShorUrlResponse {
	return &ShorUrlResponse{ShortUrl: shortUrl}
}

func NewShortenerController(shortenerService ShortenerService) web.HttpController {
	return ShortenerController{
		Service: shortenerService,
	}
}
