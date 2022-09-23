run:
	go build -o bin/goarch . && ./bin/goarch

clean:
	rm -rf .dev server .cd bin cmd postgres logger