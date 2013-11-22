package braintree

type (
	ApiErrors struct {
		ErrorMessage     string     `xml:"message,omitempty"`
		GenericErrors    []ApiError `xml:"errors>errors>error,omitempty"`
		CreditCardErrors []ApiError `xml:"errors>credit-card>errors>error,omitempty"`
	}
	ApiError struct {
		Code      string `xml:"code,omitempty"`
		Attribute string `xml:"attribute,omitempty"`
		Message   string `xml:"message,omitempty"`
	}
)

func (aes ApiErrors) ErrorCount() int {
	return len(aes.GenericErrors) +
		len(aes.CreditCardErrors)
}

func (aes ApiErrors) Success() bool {
	return aes.ErrorCount() == 0 && aes.ErrorMessage == ""
}
