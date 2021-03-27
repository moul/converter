GOPKG ?=	moul.io/converter
DOCKER_IMAGE ?=	moul/converter
GOBINS ?=	./cmd/json2toml ./cmd/converter

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ converter -h' > .tmp/usage.txt
	converter -h 2>&1 >> .tmp/usage.txt
	echo 'foo@bar:~$$ echo -n "Hello World!" | converter _bytes-to-string rev' > .tmp/examples.txt
	echo -n 'Hello World!' | converter _bytes-to-string rev >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter md5' >> .tmp/examples.txt
	echo 'Hello World!' | converter md5 >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter md5 md5' >> .tmp/examples.txt
	echo 'Hello World!' | converter md5 md5 >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter md5 md5 md5' >> .tmp/examples.txt
	echo 'Hello World!' | converter md5 md5 md5 >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter _bytes-to-string rev _string-to-bytes md5 hex upper rev' >> .tmp/examples.txt
	echo 'Hello World!' | converter _bytes-to-string rev _string-to-bytes md5 hex upper rev >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter _bytes-to-string rev _string-to-bytes md5 hex upper rev base64-decode' >> .tmp/examples.txt
	echo 'Hello World!' | converter _bytes-to-string rev _string-to-bytes md5 hex upper rev base64-decode >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter _bytes-to-string rev _string-to-bytes md5 hex upper rev base64-decode bytes-to-string' >> .tmp/examples.txt
	echo 'Hello World!' | converter _bytes-to-string rev _string-to-bytes md5 hex upper rev base64-decode _bytes-to-string >> .tmp/examples.txt
	embedmd -w README.md
	sed -i 's/[[:blank:]]*$$//' README.md
	rm -rf .tmp
.PHONY: generate

lint:
	cd tool/lint; make
.PHONY: lint
