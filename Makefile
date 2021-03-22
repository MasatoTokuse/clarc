test:
	go test `go list ./... | grep -v app/models`
sqlboiler:
	cd app && sqlboiler mysql