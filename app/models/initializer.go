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
		{ID: "bu1", Name: "Burek", ExpirationDate: time.Now().Add(2 * 24 * time.Hour).UTC().Format("02-01-2006"), Price: 2, Quantity: 10},
		{ID: "pi1", Name: "Pizza", ExpirationDate: time.Now().Add(2 * 24 * time.Hour).UTC().Format("02-01-2006"), Price: 3, Quantity: 10},
		{ID: "eg1", Name: "Egg", ExpirationDate: time.Now().Add(10 * 24 * time.Hour).UTC().Format("02-01-2006"), Price: 2, Quantity: 10},
		{ID: "or1", Name: "Orange", ExpirationDate: time.Now().Add(20 * 24 * time.Hour).UTC().Format("02-01-2006"), Price: 3, Quantity: 10},
	}

	gems := []Product{
		{ID: "ru1", Name: "Ruby", Price: 15, Quantity: 10},
		{ID: "to1", Name: "Topaz", Price: 9, Quantity: 3},
		{ID: "op1", Name: "Opal", Price: 33, Quantity: 11},
	}

	movies := []Product{
		{ID: "ews1", Name: "Eyes Wide Shut", Price: 26, Quantity: 12},
		{ID: "im1", Name: "Iron Man", Price: 24, Quantity: 12},
		{ID: "sh1", Name: "Shrek", Price: 25, Quantity: 12},
	}

	allProducts := append(foodItems, gems...)
	allProducts = append(allProducts, movies...)

	traders := []Trader{
		{ID: "tr1", TraderType: Food, PIB: "pib1", Products: getIds(foodItems)},
		{ID: "tr2", TraderType: Gem, PIB: "pib2", Products: getIds(gems)},
		{ID: "tr3", TraderType: Movie, PIB: "pib3", Products: getIds(movies)},
	}

	users := []User{
		{ID: "jj1", Name: "Jovan", LastName: "Jovanovski", Email: "jova@jovski.com", AccountBalance: 150},
		{ID: "ik1", Name: "Ivica", LastName: "Kolar", Email: "boorgir@the.com", AccountBalance: 0},
		{ID: "sg1", Name: "Satoru", LastName: "Gojo", Email: "jujutsu@kaisen.com", AccountBalance: 2000},
	}

	return InitialChainState{Products: allProducts, Traders: traders, Users: users}
}
