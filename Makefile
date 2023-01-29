



gen:
	go run github.com/99designs/gqlgen generate

server:
	go run .	

.PHONY:
	gen server
