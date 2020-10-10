# 0001-http-google-cloud-functions-in-go

## Prerequisite

__note:-__ skip this section, if `gcloud` Google Cloud SDK has already been installed and setup

1. install gcloud sdk, see [instructions](https://cloud.google.com/sdk/docs/install)

1. initial gcloud sdk, see [instructions](https://cloud.google.com/sdk/docs/initializing)

    ```bash
    gcloud init
    # or
    gcloud init --console-only
    ```

## Set up

1. verify current Google Cloud Project

    ```bash
    gcloud config get-value project
    # or to switch project
    gcloud config set project <project-id>
    ```

1. ensure the Cloud Functions API is enabled

    ```bash
    gcloud services enable cloudfunctions.googleapis.com
    ```

1. project structure

    ```bash
      .
      ├── Makefile       # Makefile
      ├── README.md      # read me document
      ├── cmd
      │   └── main.go    # binary to run function locally
      ├── go.mod         # go module definition.
      ├── go.sum         # hash of specific module versions
      ├── gopher.go      # go file for Gopher function
      ├── gopher_test.go # test file for Gopher function
      ├── gopher.png     # image resource, the gopher!
      ├── hello.go       # go file for HelloGopher function
      └── hello_test.go  # test file for HelloGopher function
    ```

## Ex1: HelloGopher function

1. create a function `HelloGopher` file `hello.go`

    ```golang
    package hellogo

    import (
      "fmt"
      "net/http"
    )

    // HelloGopher prints "Hello, Gopher."
    func HelloGopher(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintln(w, "Hello, Gopher.")
    }
    ```

1. deploy the funcion to Google Cloud Finctions

    ```bash
    make deploy-hello-gopher
    # or
    gcloud functions deploy HelloGopher \
        --region asia-northeast1 \
        --runtime go111 \
        --trigger-http \
        --allow-unauthenticated
    # for windows
    gcloud functions deploy HelloGopher --region asia-northeast1 --runtime go111 --trigger-http --allow-unauthenticated
    ```

1. review result

    ```text
    Deploying function (may take a while - up to 2 minutes)...done.
    availableMemoryMb: 256
    entryPoint: HelloGopher
    httpsTrigger:
      url: https://region-my-project.cloudfunctions.net/HelloGopher
    ```

1. open the url in browser

   ```bash
   curl https://region-my-project.cloudfunctions.net/HelloGopher
   ```

## Ex2: Gopher function

![gopher](./gopher.png)

1. create a function `Gopher` in file `gopher.go`

    ```golang
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
      // Uses directory "serverless_function_source_code" as defined in the Go Functions Framework Buildpack.
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

    ```

    __note:-__ this function read local image file `gopher.png`. the file will be located in `serverless_function_source_code` folder on server, see [Go Functions Framework Buildpack](https://github.com/GoogleCloudPlatform/buildpacks/blob/56eaad4dfe6c7bd0ecc4a175de030d2cfab9ae1c/cmd/go/functions_framework/main.go#L38).

1. deploy the funcion to Google Cloud Finctions

    ```bash
    make deploy-gopher
    # or
    gcloud functions deploy Gopher \
        --region asia-northeast1 \
        --runtime go111 \
        --trigger-http \
        --allow-unauthenticated
    # for windows
    gcloud functions deploy Gopher --region asia-northeast1 --runtime go111 --trigger-http --allow-unauthenticated
    ```

1. review result

    ```text
    Deploying function (may take a while - up to 2 minutes)...done.
    availableMemoryMb: 256
    entryPoint: Gopher
    httpsTrigger:
      url: https://region-my-project.cloudfunctions.net/Gopher
    ```

1. open the url in browser

   ```bash
   curl https://region-my-project.cloudfunctions.net/Gopher
   ```

## Ex3: test HelloGopher

1. create a `TestHelloGopher` in file `hello_test.go`

    ```golang
    package hellogo

    import (
      "io/ioutil"
      "net/http"
      "net/http/httptest"
      "testing"

      "github.com/stretchr/testify/assert"
    )

    func TestHelloGopher(t *testing.T) {
      rr := httptest.NewRecorder()
      req := httptest.NewRequest("GET", "/", nil)
      HelloGopher(rr, req)
      if r := rr.Result(); r.StatusCode != http.StatusOK {
        t.Errorf("Gopher StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusOK)
      } else {
        responseData, _ := ioutil.ReadAll(r.Body)
        assert.Equal(t, "Hello, Gopher.\n", string(responseData))
      }
    }
    ```

1. run test

    ```bash
    make test
    # or
    go test
    ```

## Ex4: test Gopher

1. create a `TestGopher` in file `gopher_test.go`

    ```golang
    package hellogo

    import (
      "net/http"
      "net/http/httptest"
      "testing"

      "github.com/stretchr/testify/assert"
    )

    func TestGopher(t *testing.T) {
      rr := httptest.NewRecorder()
      req := httptest.NewRequest("GET", "/", nil)
      Gopher(rr, req)
      if r := rr.Result(); r.StatusCode != http.StatusOK {
        t.Errorf("Gopher StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusOK)
      } else {
        contentType := r.Header.Get("Content-Type")
        assert.Equal(t, "image/png", contentType)
      }
    }
    ```

1. run test

    ```bash
    make test
    # or
    go test
    ```

## Ex5: run locally

1. create a `main` function in `cmd/main.go`

    ```golang
    package main

    import (
      "log"
      "net/http"
      "os"

      "github.com/foxfoxio/learn-go/0001-http-google-cloud-functions-in-go/functions"
    )

    func main() {
      http.HandleFunc("/Gopher", hellogo.Gopher)
      http.HandleFunc("/HelloGopher", hellogo.HelloGopher)

      port := os.Getenv("PORT")
      if port == "" {
        port = "8080"
      }
      log.Printf("Listening on port %s", port)
      if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
      }
    }
    ```

1. start local server

    ```bash
    make run-local
    # or
    go run cmd/main.go
    ```
