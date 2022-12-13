package watermark

import (
	"context"

	"github.com/MasonPhan2110/MicroserviceEx/internal"
)

type Service interface {
	// Get the list of all documents

	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
	Status(ctx context.Context, ticketId string) (internal.Status, error)
	Watermark(ctx context.Context, ticketId, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
