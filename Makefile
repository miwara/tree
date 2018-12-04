BINALY="tree"

all:
	go build -o $(BINALY) -v
clean:
	rm $(BINALY)
