VERSION :=	$(shell cat .goxc.json | jq -c .PackageVersion | sed 's/"//g')
SOURCES :=	$(shell find . -name "*.go")


.PHONY: build
build: json2toml converter


json2toml converter: $(SOURCES)
	go get ./...
	go build -o $@ ./cmd/$@


.PHONY: test
test:
	go get -t ./...
	go test -v ./...


.PHONY: cover
cover:
	rm -f profile.out
	go test -covermode=count -coverpkg=. -coverprofile=profile.out .


.PHONY: convey
convey:
	goconvey -cover -port=10042 -workDir="$(realpath .)" -depth=1


.PHONY: install
install:
	go install ./cmd/converter


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
