swagger:
	cd api/swagger && swagger generate server -A pokemon-Api -f ./swagger.yml && mv cmd/pokemon-api-server/main.go ../../cmd && rm -r cmd