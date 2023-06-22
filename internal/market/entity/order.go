package entity

type OrderType string

type Order struct {
	ID 				string
	Investor 		*Investor //esse é nosso investor de investor.go
	Asset 			*Asset //esse é nosso Asset de asset.go
	Shares 			int
	PendingShares 	int
	Price 			float64
	OrderType 		string
	Status 			string
	Transactions 	[]*Transaction
}

func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order{
	return &Order{
		ID: orderID,
		Investor: investor,
		Asset: asset,
		Shares: shares,
		PendingShares: shares,
		Price: price,
		OrderType: orderType,
		Status: "OPEN", //TODA ORDEM AO CRIADA, SERÁ SEMPRE ABERTA, DEPOIS DISSO, CLOSED
		Transactions: []*Transaction{},
	}
}