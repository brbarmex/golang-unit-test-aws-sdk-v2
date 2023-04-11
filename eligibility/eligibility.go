package eligibility

import "context"

type PurchaseTransactions struct {
	ID          string  `json:"id"`
	PaymentType string  `json:"paymentType"`
	Amount      float64 `json:"amount"`
	CreatedAt   string  `json:"createdAt"`
}

type Customer struct {
	Id                   string                 `json:"id"`
	FullName             string                 `json:"fullName"`
	PhoneNumber          string                 `json:"phoneNumber"`
	Address              string                 `json:"address"`
	CreatedAt            string                 `json:"createdAt"`
	AccountKind          int                    `json:"kind"`
	PurchaseTransactions []PurchaseTransactions `json:"purchaseTransactions"`
}

type Service struct {
}

func (self *Service) Validate(ctx context.Context, customer *Customer) (bool, error) {
	return false, nil
}
