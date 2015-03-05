PROGNAME := Now

all:
	go get -d -v
	gofmt -w .
	elm-make Game.elm --output ./build/elm/Game.html

	go get github.com/jteeuwen/go-bindata/...
	go-bindata -o Game.go build/elm/
	go build -v -o ./build/Now

run:
	./build/Now

clean:
	rm -rf ./build
