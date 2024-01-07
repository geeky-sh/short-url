swag:
	rm ./swagger.yaml
	swagger generate spec -w cmd/server -o ./swagger.yaml
	swagger serve -F=swagger swagger.yaml
