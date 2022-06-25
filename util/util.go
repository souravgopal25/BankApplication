package util

//Constant for all supported currencies
const (
	INR = "INR"
	USD = "USD"
	EUR = "EUR"
	YEN = "YEN"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR, YEN, CAD:
		return true
	}
	return false
}
