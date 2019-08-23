CMDPATH=cmd
BINPATH=bin
GOSRC=./...
COVERDIR=.cover
COVERFILE=.cover.out
VERSIONFILE=version
TARGETS=visitor-count
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
	@go run cmd/visitor-count.go

build: $(BINARIES)

$(BINARIES): $(BINPATH)/%: $(CMDPATH)/%.go
	@CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo \
		-o $@ $(CMDPATH)/$*.go

clean:
	@rm -rf $(BINPATH)

image:  build
	docker build -t visitor-count-$(DOCKER_TAG) -f build/docker/dockerfile .

compose: build
	docker-compose -f build/docker/docker-compose.yml up --build -d

compose-down: build
	docker-compose -f build/docker/docker-compose.yml down

.PHONY: dep test vet cover run build clean image
