mocks:
	@mockery --all --keeptree --case underscore

test:
	go test ./delivery/resthandler ./usecase ./repository  -coverprofile=cover.out
	go tool cover -html=cover.out
