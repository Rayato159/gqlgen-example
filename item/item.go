package item

import "github.com/Rayato159/gql-example/graph/model"

var items = map[string][]*model.Item{
	"1": {
		{
			ItemID: "1",
			Name:   "Sword",
		},
		{
			ItemID: "2",
			Name:   "Kanata",
		},
		{
			ItemID: "3",
			Name:   "Bow",
		},
	},
}

func QueryItem(playerID string) []*model.Item {
	return items[playerID]
}
