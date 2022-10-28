package response

import (
	"time"

	"github.com/cperdiansyah/golang-h8-dts/assignment-2/model/domain"
)

type OrderResponse struct {
	ID           uint          `json:"id"`
	CustomerName string        `json:"customer_name"`
	OrderedAt    time.Time     `json:"ordered_at"`
	Items        []domain.Item `json:"items"`
}
