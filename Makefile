VERSION :=	$(shell cat .goxc.json | jq -c .PackageVersion | sed 's/"//g')
SOURCES :=	$(shell find . -name "*.go")
GOENV ?=	GO15VENDOREXPERIMENT=1
LOCAL_PKGS ?=	$(shell go list ./... | grep -v /vendor/)
VERSION :=	$(shell cat .goxc.json | grep "PackageVersion" | egrep -o "([0-9]{1,}\.)+[0-9]{1,}")
REV :=		$(shell git rev-parse HEAD || git ls-remote https://github.com/scaleway/scaleway-cli  | grep -F $(VERSION) | head -n1 | awk '{print $$1}' || echo "nogit")
TAG :=		$(shell git describe --tags --always || echo $(VERSION) || echo "nogit")
LDFLAGS =	"-X main.GITCOMMIT=$(TAG) \
		-X main.VERSION=$(VERSION) \
		-X main.BUILD_DATE=$(date +%s)"
GO ?=		$(GOENV) go


.PHONY: build
build: json2toml converter


json2toml converter: $(SOURCES)
	$(GO) build -ldflags $(LDFLAGS) -o $@ ./cmd/$@


.PHONY: test
test:
	$(GO) test  -ldflags $(LDFLAGS) -v .


.PHONY: cover
cover:
	rm -f profile.out
	$(GO) test -ldflags $(LDFLAGS) -covermode=count -coverpkg=. -coverprofile=profile.out .


.PHONY: convey
convey:
	$(GO) get github.com/smartystreets/goconvey
	goconvey -cover -port=10042 -workDir="$(realpath .)" -depth=1


.PHONY: install
install:
	$(GO) install ./cmd/converter


.PHONY: build-docker
build-docker: contrib/docker/.docker-container-built
	@echo "now you can 'docker push moul/converter'"


dist/latest/converter_latest_linux_386: $(SOURCES)
	mkdir -p dist
	rm -f dist/latest
	(cd dist; ln -s $(VERSION) latest)
	goxc -bc="linux,386" xc
	cp dist/latest/converter_$(VERSION)_linux_386 $@


contrib/docker/.docker-container-built: dist/latest/converter_latest_linux_386
	cp $< contrib/docker/converter
	docker build -t moul/converter:latest contrib/docker
	docker tag -f moul/converter:latest moul/converter:$(shell echo $(VERSION) | sed 's/\+/plus/g')
	docker run -it --rm moul/converter --list-filters
	docker inspect --type=image --format="{{ .Id }}" moul/converter > $@.tmp
	mv $@.tmp $@
