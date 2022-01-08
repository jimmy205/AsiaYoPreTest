package currencyBin_test

import (
	"asiayo/business/currencyBin"
	"testing"
)

type transferTestCase struct {
	name      string
	rate      float64
	amount    float64
	expect    string
	expectErr error
}

func TestTransfer_Normal(t *testing.T) {
	currency := currencyBin.NewCurrency()

	testCase := transferTestCase{
		name:      "test transfer normal",
		rate:      0.5,
		amount:    1,
		expect:    "0.50",
		expectErr: nil,
	}

	t.Run(testCase.name, func(t *testing.T) {
		got, gotErr := currency.Transfer(testCase.rate, testCase.amount)
		if testCase.expect != got || testCase.expectErr != gotErr {
			t.Fatalf(
				"[ expect: %s, got: %s ] , [ expectErr: %v, got: %v ]",
				testCase.expect, got, testCase.expectErr, gotErr,
			)
		}
	})
}

func TestTransfer_ThousandComma(t *testing.T) {
	currency := currencyBin.NewCurrency()

	testCase := transferTestCase{
		name:      "test transfer thousand comma",
		rate:      100,
		amount:    1234,
		expect:    "123,400.00",
		expectErr: nil,
	}

	t.Run(testCase.name, func(t *testing.T) {
		got, gotErr := currency.Transfer(testCase.rate, testCase.amount)
		if testCase.expect != got || testCase.expectErr != gotErr {
			t.Fatalf(
				"[ expect: %s, got: %s ] , [ expectErr: %v, got: %v ]",
				testCase.expect, got, testCase.expectErr, gotErr,
			)
		}
	})
}

func TestTransferAmount_LessThan0(t *testing.T) {
	currency := currencyBin.NewCurrency()

	testCase := transferTestCase{
		name:      "test transfer amount less than 0",
		rate:      0.5,
		amount:    -1,
		expect:    "",
		expectErr: currencyBin.AmountLessThan0Err,
	}

	t.Run(testCase.name, func(t *testing.T) {
		got, gotErr := currency.Transfer(testCase.rate, testCase.amount)
		if testCase.expect != got || testCase.expectErr.Error() != gotErr.Error() {
			t.Fatalf(
				"[ expect: %s, got: %s ] , [ expectErr: %v, got: %v ]",
				testCase.expect, got, testCase.expectErr, gotErr,
			)
		}
	})
}

func TestTransferRate_LessThan0(t *testing.T) {
	currency := currencyBin.NewCurrency()

	testCase := transferTestCase{
		name:      "test transfer rate less than 0",
		rate:      -0.5,
		amount:    1,
		expect:    "",
		expectErr: currencyBin.RateLessThan0Err,
	}

	t.Run(testCase.name, func(t *testing.T) {
		got, gotErr := currency.Transfer(testCase.rate, testCase.amount)
		if testCase.expect != got || testCase.expectErr.Error() != gotErr.Error() {
			t.Fatalf(
				"[ expect: %s, got: %s ] , [ expectErr: %v, got: %v ]",
				testCase.expect, got, testCase.expectErr, gotErr,
			)
		}
	})
}

type findRateTestCase struct {
	name      string
	from      string
	to        string
	expect    float64
	expectErr error
}

func TestFindRate_Normal(t *testing.T) {
	currency := currencyBin.NewCurrency()

	testCase := findRateTestCase{
		name:      "test find rate normal",
		from:      "TWD",
		to:        "USD",
		expect:    0.03281,
		expectErr: nil,
	}
	t.Run(testCase.name, func(t *testing.T) {
		got, gotErr := currency.FindRate(testCase.from, testCase.to)
		if testCase.expect != got || testCase.expectErr != gotErr {
			t.Fatalf(
				"[ expect: %.6f, got: %.6f ] , [ expectErr: %v, got: %v ]",
				testCase.expect, got, testCase.expectErr, gotErr,
			)
		}
	})
}

func TestFindRate_CurrencyNotExist(t *testing.T) {
	currency := currencyBin.NewCurrency()

	testCase := findRateTestCase{
		name:      "test find rate currency not exist",
		from:      "Not Currency",
		to:        "Not Currency 2",
		expect:    0,
		expectErr: currencyBin.CurrencyNotFoundErr,
	}
	t.Run(testCase.name, func(t *testing.T) {
		got, gotErr := currency.FindRate(testCase.from, testCase.to)
		if testCase.expect != got || testCase.expectErr != gotErr {
			t.Fatalf(
				"[ expect: %.6f, got: %.6f ] , [ expectErr: %v, got: %v ]",
				testCase.expect, got, testCase.expectErr, gotErr,
			)
		}
	})
}
