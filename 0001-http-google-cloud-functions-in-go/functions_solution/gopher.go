package hellogo

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Gopher prints a gopher.
func Gopher(w http.ResponseWriter, r *http.Request) {
	// Read the gopher image file.
	// Uses directory "serverless_function_source_code" as defined in the Go
	// Functions Framework Buildpack.
	// See https://github.com/GoogleCloudPlatform/buildpacks/blob/56eaad4dfe6c7bd0ecc4a175de030d2cfab9ae1c/cmd/go/functions_framework/main.go#L38.
	path := "serverless_function_source_code/gopher.png"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Fall back to the current working directory if that file doesn't exist.
		path = "gopher.png"
	}
	f, err := os.Open(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Write the gopher image to the response writer.
	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "image/png")
}
