package graphql

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Products(ctx context.Context, distinct_on []ProductsSelectColumn, limit *int, offset *int, order_by []Products_order_by, where *Products_bool_exp) ([]Products, error) {
	panic("not implemented")
}
func (r *queryResolver) ProductsByPk(ctx context.Context, uuid string) (*Products, error) {
	panic("not implemented")
}
