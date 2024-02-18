package inventory

import "github.com/Rayato159/gql-example/graph/model"

func QueryInventory(playerID string) *model.Inventory {
	return &model.Inventory{
		PlayerID: playerID,
	}
}
