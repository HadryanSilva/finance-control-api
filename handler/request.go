package handler

import "fmt"

type AccountRequest struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *AccountRequest) Validate() error {
	if r.Name == "" && r.Balance >= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if r.Balance < 0 {
		return errParamIsRequired("balance", "float64")
	}

	return nil
}
