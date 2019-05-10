package api

import "fmt"

type e struct {
	Text string `json:"error"`
}

// Error wrap an error in a struct to encode in JSON
func Error(err error) e {
	return e{fmt.Sprintf("%v", err)}
}
