install_deps:
	go mod download

# Standard go test
test:
	go test ./... -v -race

# Make sure no unnecessary dependecies are present
go-mod-tidy:
	go mod tidy -v
	git diff-index --quiet HEAD

# Run all tests & linters in CI
ci: test go-mod-tidy


generate_oto:
	oto \
		-template ./vendor/oto/templates/server.go.plush \
		-out ./api/generated/oto.gen.go \
		-ignore Ignorer -pkg generated ./api/definitions
	gofmt -w ./api/oto.gen.go ./api/oto.gen.go