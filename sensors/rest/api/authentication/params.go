package authentication

import (
	"encoding/json"
	"fmt"
	"log"
)

//LoginParams parameters sent in login request
type LoginParams struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

//RegistrationParams parameters sent in to register request
type RegistrationParams struct {
	Email    string `json:"email`
	Username string `json:"username"`
	Password string `json:"password"`
}

//ValidateRegisterParams validates registration parameters
func ValidateRegisterParams(params *RegistrationParams) (bool, error) {

	var pMap map[string]interface{}
	in, _ := json.Marshal(params)
	_ = json.Unmarshal(in, &pMap)

	var err error
	for k, v := range pMap {
		if v == "" {
			log.Printf("checking %v value passed %v\n", k, v)
			err = ErrInvalidParams{message: fmt.Sprintf("%v is required", k)}
			break
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}
