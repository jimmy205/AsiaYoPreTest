package currencyBin

import (
	"errors"

	"github.com/shopspring/decimal"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type ICurrency interface {
	FindRate(from, to string) (float64, error)
	Transfer(rate float64, amount float64) (string, error)
}

var (
	CurrencyNotFoundErr error = errors.New("currency not found.")
	AmountLessThan0Err  error = errors.New("amount must greater than or equal 0.")
	RateLessThan0Err    error = errors.New("rate must greater than or equal 0.")
)

// 先暫時寫死
var exchangeRate = map[string]map[string]float64{
	"TWD": {
		"TWD": 1,
		"JPY": 3.669,
		"USD": 0.03281,
	},
	"JPY": {
		"TWD": 0.26956,
		"JPY": 1,
		"USD": 0.00885,
	},
	"USD": {
		"TWD": 30.44,
		"JPY": 111.801,
		"USD": 1,
	},
}

type currency struct {
}

func NewCurrency() ICurrency {
	return &currency{}
}

func (c *currency) FindRate(from, to string) (float64, error) {

	rate, ok := exchangeRate[from][to]
	if !ok {
		return 0, CurrencyNotFoundErr
	}

	return rate, nil
}

func (c *currency) addComma(amount float64) string {
	return message.NewPrinter(language.English).Sprintf("%.2f", amount)
}

func (c *currency) Transfer(rate, amount float64) (string, error) {

	if amount < 0 {
		return "", AmountLessThan0Err
	}
	if rate < 0 {
		return "", RateLessThan0Err
	}

	res, _ := decimal.NewFromFloat(amount).Mul(
		decimal.NewFromFloat(rate),
	).Round(2).Float64()

	return c.addComma(res), nil
}
