// Util package contains files that help develop functionalities to the app
package util

// HttpStatus contains the http status id and message
type HttpStatus struct {
	ID      int    `json:"http_status_id"`
	Message string `json:"http_status_message"`
}
