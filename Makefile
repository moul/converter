VERSION :=	$(shell cat .goxc.json | jq -c .PackageVersion | sed 's/"//g')
SOURCES :=	$(shell find . -name "*.go")
GOENV ?=	GO15VENDOREXPERIMENT=1
GO ?=		$(GOENV) go
GODEP ?=	$(GOENV) godep
LOCAL_PKGS ?=	$(shell go list ./... | grep -v /vendor/)


.PHONY: build
build: json2toml converter


json2toml converter: $(SOURCES)
	$(GO) get ./...
	$(GO) build -o $@ ./cmd/$@


.PHONY: test
test:
	$(GODEP) restore
	$(GO) get -t .
	$(GO) test -v .


.PHONY: godep-save
godep-save:
	$(GODEP) save $(LOCAL_PKGS)


.PHONY: cover
cover:
	rm -f profile.out
	$(GO) test -covermode=count -coverpkg=. -coverprofile=profile.out .


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
