package graph

import "ddd_sample/src/application/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	rankings []*model.Ranking
	items []*model.Item
}
