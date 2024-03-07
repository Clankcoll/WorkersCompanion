// internal/handlers/union_calculator.go
package handlers

import (
	"fmt"
	"net/http"
)

// HomeOffice handles requests for the time you can spend @home
func HomeOfficeCalculator(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Time @ Homeoffice aka HO Calc will be implemented here.")
}
