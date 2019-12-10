package redemption

import (
	"time"
)

type RedeemResultType string

const (
	RedeemResultTypeError   RedeemResultType = "error"
	RedeemResultTypeSuccess RedeemResultType = "success"

	LinksPageCount = 50
)

type RedeemResult struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	ImageURL    string           `json:"image_url,omitempty"`
	Type        RedeemResultType `json:"type"`
	Error       string           `json:"error,omitempty"`
	ResultId    []string         `json:"result_id,omitempty"`
}

type Links []*Link

type Link struct {
	Link           string    `json:"link" bson:"link"`
	Code           string    `json:"code" bson:"code"`
	Provider       string    `json:"provider" bson:"provider"`
	Valid          bool      `json:"valid" bson:"valid"`
	ExpirationDate time.Time `json:"expiration_date" bson:"expiration_date"`
	CreatedDate    time.Time `json:"created_date" bson:"created_date"`
	Asset          Assets    `json:"asset" bson:"assets"`
}

type CreateLinks struct {
	Provider  string  `json:"provider" bson:"provider"`
	LinkCount int     `json:"link_count" bson:"link_count"`
	Assets    *Assets `json:"asset" bson:"asset"`
}

type Assets struct {
	Coin   uint     `json:"coin" bson:"coin"`
	Used   bool     `json:"used" bson:"used"`
	Assets []*Asset `json:"assets" bson:"assets"`
}

type Asset struct {
	Amount  int64  `json:"amount" bson:"amount"`
	TokenId string `json:"token_id,omitempty" bson:"token_id,omitempty"`
	Error   string `json:"error,omitempty" bson:"error,omitempty"`
}

type Redeem struct {
	Code      string    `json:"code" bson:"coin"`
	Observers Observers `json:"observers" bson:"observers"`
}

type Observer struct {
	Coin       uint     `json:"coin" bson:"coin"`
	PublicKeys []string `json:"publicKeys,omitempty" bson:"publicKeys,omitempty"`
	Addresses  []string `json:"addresses,omitempty" bson:"addresses,omitempty"`
}

type Observers []Observer

func (o Observers) GetCoinObservers(coin uint) Observers {
	filter := make(Observers, 0)
	for _, observer := range o {
		if coin == observer.Coin {
			filter = append(filter, observer)
		}
	}
	return filter
}

type CoinHosts []CoinHost

type CoinHost struct {
	Coin uint   `json:"coin" bson:"coin"`
	Host string `json:"host" bson:"host"`
}

type Success struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}