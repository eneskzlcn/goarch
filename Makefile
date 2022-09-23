run:
	go build -o bin/goarch . && ./bin/goarch

clean:
	rm -rf .dev && rm -rf server && rm -rf .cd && rm -rf bin && rm -rf cmd