package httpserver


type UrlRequest struct {
	Url string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}