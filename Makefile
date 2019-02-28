.PHONY: doc test all clean

all: doc test

doc: README.md

README.md: doc.go .godocdown.tmpl
	godocdown --output=README.md

test: *.go
	go test -race -cover ./...

clean:
	$(RM) *~
