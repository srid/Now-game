PROGNAME := Now

all:
	go get -d -v
	go build -v
	./Now

clean:
	rm -f ./Now
