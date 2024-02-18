package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"

	"github.com/Rayato159/gql-example/graph/model"
	"github.com/Rayato159/gql-example/inventory"
	"github.com/Rayato159/gql-example/item"
)

// Inventory is the resolver for the inventory field.
func (r *queryResolver) Inventory(ctx context.Context, playerID string) (*model.Inventory, error) {
	return inventory.QueryInventory(playerID), nil
}

// Item is the resolver for the item field.
func (r *queryResolver) Item(ctx context.Context, playerID string) ([]*model.Item, error) {
	return item.QueryItem(playerID), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }