package validations

type ValidateError struct {
	Key   string `json:"key"`
	Error string `json:"error"`
	Value string `json:"value"`
}

type MinMaxValidationFields struct {
	Min        *int   `json:"min"`
	Max        *int   `json:"max"`
	FieldName  string `json:"fieldName"`
	FieldValue string `json:"fieldValue"`
}

func IntPtr(i int) *int {
	return &i
}
