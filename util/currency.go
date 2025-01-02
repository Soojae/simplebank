package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	KRW = "KRW"
)

// IsSupportedCurrency checks if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, KRW:
		return true
	default:
		return false
	}

}
