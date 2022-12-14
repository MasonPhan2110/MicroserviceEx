package endpoints

import "github.com/MasonPhan2110/MicroserviceEx/internal"

type GetRequest struct {
	Filters []internal.Filter `json:"filters,omitempty"`
}

type GetResponse struct {
	Documents []internal.Document `json:"documents"`
	Err       string              `json:"err,omitempty"`
}

type StatusRequest struct {
	TicketId string `json:"ticketId"`
}

type StatusResponse struct {
	Status internal.status `json:"status"`
	Err    string          `json:"err,omitempty"`
}

type WatermarkRequest struct {
	TicketId string `json:"ticketId"`
	Mark     string `json:"mark"`
}

type WatermarkResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

type AddDocumentRequest struct {
	Document *internal.Document `json:"document"`
}

type AddDocumentResponse struct {
	TicketId string `json:"ticketId"`
	Err      string `json:"err,omitempty`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
