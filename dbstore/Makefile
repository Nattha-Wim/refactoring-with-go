.PHONY: build test run proto clean docker-build

NAME=app
DB_FILE=app.db.bin
PB_DIR=./pb

# version settings
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
PROJECT?=github.com/anuchito/dbstore

build:
	go build -ldflags "-X ${PROJECT}/pkg/version.Release=${RELEASE} \
	-X ${PROJECT}/pkg/version.Commit=${COMMIT} -X ${PROJECT}/pkg/version.BuildTime=${BUILD_TIME}" \
	-o ${NAME} "${PROJECT}/cmd/server"
# delete cache and race when test
test:
	go clean -testcache && go test -v -race github.com/anuchito/dbstore/...

run: build
	PORT=8080 DB_FILENAME=${DB_FILE} ./app

proto:
	protoc -I ${PB_DIR} ${PB_DIR}/db.proto --go_out=${PB_DIR}

clean:
	rm -f ${NAME} ${DB_FILE} ./pkg/db/db.test.bin

docker-build:
	GOOS=linux go build -o ${NAME} "${PROJECT}/cmd/server"
	docker build -t gerlacdt/db-example:latest .
	rm -f ${NAME}
