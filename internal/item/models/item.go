package models

type Item struct {
	ItemId string `json:"item_id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

type TableKeys struct {
	PartitionKey string `dynamodbav:"partition_key" json:"-"`
	SortKey      string `dynamodbav:"sort_key, omitempty" json:"-"`
}

type DBItem struct {
	TableKeys
	ItemId string  `dynamodbav:"item_id" json:"item_id"`
	Name   string  `dynamodbav:"name" json:"name"`
	Price  float32 `dynamodbav:"price" json:"price"`
}

func (I *DBItem) ToItem() *Item {
	if I == nil {
		return nil
	}
	return &Item{
		ItemId: I.ItemId,
		Name:   I.Name,
		Price:  int(I.Price),
	}
}

func (I *Item) ToDBItem() *DBItem {
	return &DBItem{
		TableKeys: TableKeys{
			PartitionKey: "ITEM#" + I.ItemId,
			SortKey:      "ITEM",
		},
		ItemId: I.ItemId,
		Name:   I.Name,
		Price:  float32(I.Price),
	}
}
