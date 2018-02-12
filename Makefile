all:
	go install github.com/fatshaw/blockchain-sample
	go build
clean:
	go clean

test:
	go test


