server.prod: static/static_generated.go
	GOOS=linux GOARCH=amd64 go build -o server.prod .

static/static_generated.go: main.go $(filter-out static/static_generated.go,$(wildcard static/*))
	cd static && go generate
