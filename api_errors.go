package braintree

type (
	ApiErrors struct {
		ErrorMessage     string         `xml:"message,omitempty"`
		GenericErrors    *ApiErrorArray `xml:"errors>errors>error,omitempty"`
		CreditCardErrors *ApiErrorArray `xml:"errors>credit-card>errors>error,omitempty"`
	}

	ApiErrorArray []ApiError
	ApiError      struct {
		Code      string `xml:"code,omitempty"`
		Attribute string `xml:"attribute,omitempty"`
		Message   string `xml:"message,omitempty"`
	}
)

func (ae ApiErrorArray) Count() int {
	return len(ae)
}

func (aes ApiErrors) ErrorCount() int {
	var total int
	if aes.GenericErrors != nil {
		total += aes.GenericErrors.Count()
	}

	if aes.CreditCardErrors != nil {
		total += aes.CreditCardErrors.Count()
	}
	return total
}

func (aes ApiErrors) Success() bool {
	return aes.ErrorCount() == 0 && aes.ErrorMessage == ""
}

func (aea *ApiErrorArray) For(attribute string) []string {
	var errs []string
	if aea != nil {
		for _, e := range *aea {
			if e.Attribute == attribute {
				errs = append(errs, e.Message)
			}
		}
	}
	return errs
}
