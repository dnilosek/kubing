CMDPATH=cmd
BINPATH=bin
GOSRC=./...
COVERDIR=.cover
COVERFILE=.cover.out
VERSIONFILE=version
TARGETS=visits
BINARIES=$(patsubst %, $(BINPATH)/%, $(TARGETS))

ENV		?= test
PORT		?= 8080
BUILD_VERSION	?= $(shell cat $(VERSIONFILE) | head -n 1)
BUILD_NUMBER	?= 0
DOCKER_TAG	?= $(ENV)

.DEFAULT_GOAL=test

dep:
	@go get -v -t -d $(GOSRC)

test:
	@go test -v -race -coverprofile $(COVERFILE) $(GOSRC)

vet:
	@go vet $(GOSRC)

cover: test
	@mkdir -p $(COVERDIR)
	@go tool cover -html=$(COVERFILE) -o $(COVERDIR)/index.html
	@cd $(COVERDIR) && python -m SimpleHTTPServer $(PORT)

run:
	@go run cmd/visits.go

build: $(BINARIES)

$(BINARIES): $(BINPATH)/%: $(CMDPATH)/%.go
	        @CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo \
                -o $@ $(CMDPATH)/$*.go

clean:
	@rm -rf $(BINPATH)

image:  build
	docker build -t visits -f build/docker/dockerfile .

.PHONY: dep test vet cover run build clean image
