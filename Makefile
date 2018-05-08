GOCACHE?=

ENABLE_MONGODB_TESTS?=false
TEST_PSMDB_VERSION?=latest
TEST_RS_NAME?=rs
TEST_MONGODB_DOCKER_UID?=1001
TEST_PRIMARY_PORT?=65217
TEST_SECONDARY1_PORT?=65218
TEST_SECONDARY2_PORT?=65219

$(GOPATH)/bin/gocoverutil:
	go get -u github.com/AlekSi/gocoverutil 

$(GOPATH)/bin/glide:
	go get -u github.com/Masterminds/glide

vendor: $(GOPATH)/bin/glide glide.lock
	glide install

test:
	GOCACHE=$(GOCACHE) ENABLE_MONGODB_TESTS=$(ENABLE_MONGODB_TESTS) go test -v ./...

test-full-prepare:
	TEST_RS_NAME=$(TEST_RS_NAME) \
	TEST_PSMDB_VERSION=$(TEST_PSMDB_VERSION) \
	TEST_PRIMARY_PORT=$(TEST_PRIMARY_PORT) \
	TEST_SECONDARY1_PORT=$(TEST_SECONDARY1_PORT) \
	TEST_SECONDARY2_PORT=$(TEST_SECONDARY2_PORT) \
	docker-compose up -d
	test/init-test-replset-wait.sh

test-full: vendor $(GOPATH)/bin/gocoverutil
	ENABLE_MONGODB_TESTS=true \
	TEST_RS_NAME=$(TEST_RS_NAME) \
	TEST_PRIMARY_PORT=$(TEST_PRIMARY_PORT) \
	GOCACHE=$(GOCACHE) $(GOPATH)/bin/gocoverutil test -v ./...

test-full-clean:
	docker-compose down
	rm -rf vendor 2>/dev/null || true

clean:
	rm -f coverage.txt 2>/dev/null || true
