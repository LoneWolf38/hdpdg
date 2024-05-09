build:
	go mod tidy && go mod vendor
	go build -o ./bin/hdpdg .

clean:
	rm -rf bin

test:
	go test ./... -v

build-image:
	docker build -t hdpgd/kafka-go .
