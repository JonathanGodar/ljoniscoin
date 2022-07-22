


build: 
	go build -o ./bin/peer ./src/peer/*.go

brun: build
	./bin/peer

build-watcher: 
	go build -o ./bin/watcher ./src/watcher/*.go

brun-watcher: build-watcher
	./bin/watcher
