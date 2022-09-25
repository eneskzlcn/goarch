run:
	go build -o bin/goarch . && ./bin/goarch goarch ms

clean:
	rm -rf .dev server .cd bin cmd postgres logger seed rabbitmq kafka internal