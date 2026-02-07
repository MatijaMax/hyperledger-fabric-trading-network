package models

type TraderType string

const (
	Food  TraderType = "FOOD"
	Gem   TraderType = "GEM"
	Movie TraderType = "MOVIE"
)

type Trader struct {
	ID             string     `json:"id"`
	TraderType     TraderType `json:"trader_type"`
	PIB            string     `json:"pib"`
	Products       []string   `json:"products"`
	Receipts       []string   `json:"receipts"`
	AccountBalance uint       `json:"account_balance"`
}

func (p Trader) GetID() string {
	return p.ID
}
