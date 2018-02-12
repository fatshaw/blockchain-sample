test-all:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic
	go install github.com/fatshaw/blockchain-sample
	go build

clean:
	go clean
