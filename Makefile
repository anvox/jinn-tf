build:
	go build -o ./dist/jinn-tf ./main.go
	env GOOS=darwin GOARCH=amd64 go build -o ./dist/jinn-tf-darwin-amd64 ./main.go
	env GOOS=linux GOARCH=amd64 go build -o ./dist/jinn-tf-linux-amd64 ./main.go
	chmod a+x ./dist/jinn-*
test:
	go test ./...
