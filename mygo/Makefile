.PHONY: doc docd cover test bench

doc:
	go doc mygo

docd:
	godoc -http=:6060

test:
	go test -v

cover:
	go test -coverprofile=./cove.out
	go tool cover -func=cove.out
	go tool cover -html=./cove.out

bench:
	go test -bench=. -benchmem 
