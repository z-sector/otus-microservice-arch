package repo

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/wuzyk/otus-microservice-arch/06/internal/domain"
)

type OrderRow struct {
	bun.BaseModel `bun:"table:orders"`

	ID              int    `bun:"id,pk,autoincrement"`
	ProductID       int    `bun:"product_id"`
	Quantity        int    `bun:"quantity"`
	ShippingAddress string `bun:"shipping_address"`
	RequestID       string `bun:"request_id,nullzero"`
}

type OrderRepo struct {
	db *bun.DB
}

func NewOrderRepo(db *bun.DB) *OrderRepo {
	return &OrderRepo{db}
}

func (r *OrderRepo) New(ctx context.Context, order *domain.Order) error {
	orderRow := OrderRow{
		ProductID:       order.ProductID,
		Quantity:        order.Quantity,
		ShippingAddress: order.ShippingAddress,
		RequestID:       order.RequestID,
	}

	_, err := r.db.
		NewInsert().
		Model(&orderRow).
		Ignore().
		Exec(ctx)
	if err != nil {
		return err
	}

	order.ID = orderRow.ID

	if order.ID == 0 {
		existingOrderRow := OrderRow{}

		q := r.db.
			NewSelect().
			Table("orders").
			Where("request_id = ?", order.RequestID)

		_, err = q.Exec(ctx, &existingOrderRow)
		if err != nil {
			return err
		}

		order.ID = existingOrderRow.ID
	}

	return err
}
