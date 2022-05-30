package domain

import "github.com/kenethrrizzo/banking/dto"

const (
	WITHDRAWAL string = "withdrawal"
	DEPOSIT    string = "deposit"
)

type Transaction struct {
	Id        string  `db:"Id"`
	AccountId string  `db:"AccountId"`
	Amount    float64 `db:"Amount"`
	Type      string  `db:"Type"`
	Date      string  `db:"Date"`
}

func (t Transaction) IsWithdrawal() bool {
	return t.Type == WITHDRAWAL
}

func (t Transaction) IsDeposit() bool {
	return t.Type == DEPOSIT
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		Id:        t.Id,
		Type:      t.Type,
		Date:      t.Date,
		AccountId: t.AccountId,
		Amount:    t.Amount,
	}
}