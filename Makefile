run:
	go build -o bin/goarch  ./cmd/goarch && ./bin/goarch

clean:
	rm -rf .dev
	#&& rm -rf server