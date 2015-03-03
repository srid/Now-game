PROGNAME := Now

all:
	cabal sandbox init
	cabal install --only-dependencies
	cabal build

run:
	./dist/build/${PROGNAME}/${PROGNAME}
