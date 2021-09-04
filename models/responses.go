package models

type JsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Body    string `json:"body"`
}
