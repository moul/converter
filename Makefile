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
