run-local:
	go run cmd/main.go

build-local:
	go build -o ./dist/main cmd/main.go

build-local-linux:
	GOOS=linux GOARCH=amd64 go build -o ./dist/main-linux-x64 cmd/main.go

test:
	go test -v

deploy-hello-gopher:
	gcloud functions deploy HelloGopher \
		--region asia-northeast1 \
		--runtime go111 \
		--trigger-http \
		--allow-unauthenticated

deploy-gopher:
	gcloud functions deploy Gopher \
		--region asia-northeast1 \
		--runtime go111 \
		--trigger-http \
		--allow-unauthenticated
