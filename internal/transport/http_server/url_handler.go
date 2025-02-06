package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s HttpServer) Create(w http.ResponseWriter, r *http.Request) {
	var urlReq UrlRequest

	if err := json.NewDecoder(r.Body).Decode(&urlReq); err != nil {
		ErrorResponse("invalid request", w, r, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	
	if err := urlReq.Validation(); err != nil {
		fmt.Println(err)
		ErrorResponse(fmt.Sprintf("invalid request: %s", err), w, r, http.StatusBadRequest)
		return
	}

	alias, err := s.UrlService.Create(r.Context(), urlReq.Url, urlReq.Alias)
	if err != nil {
		ErrorResponse("server error", w, r, http.StatusInternalServerError)
		return
	}

	ResponseOk(ResponseSuccess{
		Status: "Success",
		Alias: alias,
	}, w, r)
}

func (s HttpServer) RedirectByAlias(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	alias := query["alias"]

	if alias == "" {
		ErrorResponse("alias is empty", w, r, http.StatusBadRequest)
		return
	}

	url, err := s.UrlService.GetUrl(r.Context(), alias)
	if err != nil {
		ErrorResponse("server error", w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w,r,url, http.StatusMovedPermanently)
}