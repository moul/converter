GOPKG ?=	moul.io/converter
DOCKER_IMAGE ?=	moul/converter
GOBINS ?=	./cmd/json2toml ./cmd/converter

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ converter -h' > .tmp/usage.txt
	converter -h 2>&1 >> .tmp/usage.txt
	echo 'foo@bar:~$$ echo -n "Hello World!" | converter reverse' > .tmp/examples.txt
	echo -n 'Hello World!' | converter reverse >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter md5' >> .tmp/examples.txt
	echo 'Hello World!' | converter md5 >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter md5 md5' >> .tmp/examples.txt
	echo 'Hello World!' | converter md5 md5 >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter md5 md5 md5' >> .tmp/examples.txt
	echo 'Hello World!' | converter md5 md5 md5 >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter reverse md5 upper reverse' >> .tmp/examples.txt
	echo 'Hello World!' | converter reverse md5 upper reverse >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter reverse md5 upper reverse base64-decode' >> .tmp/examples.txt
	echo 'Hello World!' | converter reverse md5 upper reverse base64-decode >> .tmp/examples.txt
	echo 'foo@bar:~$$ echo "Hello World!" | converter reverse md5 upper reverse base64-decode bytes-to-string' >> .tmp/examples.txt
	echo 'Hello World!' | converter reverse md5 upper reverse base64-decode bytes-to-string >> .tmp/examples.txt
	embedmd -w README.md
	rm -rf .tmp
.PHONY: generate

lint:
	cd tool/lint; make
.PHONY: lint
