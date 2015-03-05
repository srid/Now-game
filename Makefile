PROGNAME := Now

all:
	go get -d -v
	go build -v -o ./build/Now
	elm-make game.elm --output ./build/game.html

run:
	./build/Now

clean:
	rm -f ./build
