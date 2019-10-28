package utills

import (
	"encoding/json"
)

// ApplicationError - Warp error for appication error response
type ApplicationError struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

func (e *ApplicationError) Error() ([]byte, error) {
	return json.Marshal(e)
}
