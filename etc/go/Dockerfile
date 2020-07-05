FROM golang:latest

WORKDIR /go/src/task-man
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate