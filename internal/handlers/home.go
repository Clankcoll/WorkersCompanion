// internal/handlers/union_calculator.go
package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HomeSweetHome.")
}
