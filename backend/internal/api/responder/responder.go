// this is for standardizing the structure of responses
package responder

import (
	"encoding/json"
	"net/http"
)

type result string

const (
	result_Success result = "success"
	result_Error   result = "error"
)

// Response is a general structure for API responses
type Response struct {
	Result  result      `json:"result"`
	Message string      `json:"message,omitempty"` // Optional for success responses
	Data    interface{} `json:"data,omitempty"`    // Data is optional for error responses
	Errors  []Error     `json:"errors,omitempty"`  // Errors is optional and used for validation or other error details
}

// Error represents an individual error with a field and message.
type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// general method for responding with json
func respondJSON(w http.ResponseWriter, status int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// sends a success response(http.StatusOK) with optional data
func RespondSuccess(w http.ResponseWriter, data interface{}) {
	response := Response{
		Result: result_Success,
		Data:   data,
	}
	respondJSON(w, http.StatusOK, response)
}

// sends a success response with optional data and the passed status response
func RespondSuccessWithStatus(w http.ResponseWriter, status int, data interface{}) {
	response := Response{
		Result: result_Success,
		Data:   data,
	}
	respondJSON(w, status, response)
}

// sends an error response with an optional error message or validation errors
func RespondError(w http.ResponseWriter, status int, message string, errors []Error) {
	response := Response{
		Result:  result_Error,
		Message: message,
		Errors:  errors,
	}
	respondJSON(w, status, response)
}
