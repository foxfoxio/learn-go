package hellogo

import (
	"fmt"
	"net/http"
)

// HelloGopher prints "Hello, Gopher."
func HelloGopher(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Gopher.")
}
