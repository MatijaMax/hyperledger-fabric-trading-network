package models

import (
	"time"
)

type InitialChainState struct {
	Products []Product
	Traders  []Trader
	Users    []User
}

func getIds[T Model](models []T) []string {
	ids := make([]string, 0)

	for _, model := range models {
		ids = append(ids, model.GetID())
	}

	return ids
}

func GetInitialChainState() InitialChainState {

	foodItems := []Product{
		{ID: ToProductID("bu1"), Name: "Burek", ExpirationDate: time.Now().Add(2 * 24 * time.Hour).UTC().Format("02-01-2006"), Price: 2, Quantity: 10, TraderID: "tr1"},
		{ID: ToProductID("pi1"), Name: "Pizza", ExpirationDate: time.Now().Add(2 * 24 * time.Hour).UTC().Format("02-01-2006"), Price: 3, Quantity: 10, TraderID: "tr1"},
		{ID: ToProductID("eg1"), Name: "Egg", ExpirationDate: time.Now().Add(10 * 24 * time.Hour).UTC().Format("02-01-2006"), Price: 2, Quantity: 10, TraderID: "tr1"},
		{ID: ToProductID("or1"), Name: "Orange", ExpirationDate: time.Now().Add(20 * 24 * time.Hour).UTC().Format("02-01-2006"), Price: 3, Quantity: 10, TraderID: "tr1"},
	}

	gems := []Product{
		{ID: ToProductID("ru1"), Name: "Ruby", Price: 15, Quantity: 10, TraderID: "tr2"},
		{ID: ToProductID("to1"), Name: "Topaz", Price: 9, Quantity: 3, TraderID: "tr2"},
		{ID: ToProductID("op1"), Name: "Opal", Price: 33, Quantity: 11, TraderID: "tr2"},
	}

	movies := []Product{
		{ID: ToProductID("si1"), Name: "Eyes Wide Shut", Price: 26, Quantity: 12, TraderID: "tr3"},
		{ID: ToProductID("pi1"), Name: "Iron Man", Price: 24, Quantity: 12, TraderID: "tr3"},
		{ID: ToProductID("br1"), Name: "Shrek 2", Price: 25, Quantity: 12, TraderID: "tr3"},
	}

	allProducts := append(foodItems, gems...)
	allProducts = append(allProducts, movies...)

	traders := []Trader{
		{ID: ToTraderID("tr1"), TraderType: Food, PIB: "pib1", Products: getIds(foodItems), Receipts: make([]string, 0)},
		{ID: ToTraderID("tr2"), TraderType: Gem, PIB: "pib2", Products: getIds(gems), Receipts: make([]string, 0)},
		{ID: ToTraderID("tr3"), TraderType: Movie, PIB: "pib3", Products: getIds(movies), Receipts: make([]string, 0)},
	}

	users := []User{
		{ID: ToUserID("jj1"), Name: "Jovan", LastName: "Jovanovski", Email: "jova@jovski.com", AccountBalance: 150, ReceiptsID: make([]string, 0)},
		{ID: ToUserID("ik1"), Name: "Ivica", LastName: "Kolar", Email: "boorgir@the.com", AccountBalance: 0, ReceiptsID: make([]string, 0)},
		{ID: ToUserID("sg1"), Name: "Satoru", LastName: "Gojo", Email: "jujutsu@kaisen.com", AccountBalance: 2000, ReceiptsID: make([]string, 0)},
	}

	return InitialChainState{Products: allProducts, Traders: traders, Users: users}
}
