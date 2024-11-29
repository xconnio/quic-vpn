run-server:
	go run ./cmd/server

run-client:
	go run ./cmd/client

gen-certs:
	openssl req -x509 -newkey rsa:4096 -nodes -keyout key.pem -out cert.pem -days 365
