FROM golang:1.23.2

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./ ./

RUN go mod download
RUN go build -o marketplace ./cmd/main.go

COPY ./migrations /migrations

ENTRYPOINT ["sh", "-c", "migrate -source file://./migrations -database postgres://root:123@postgres:5432/marketplace?sslmode=disable up 2 && ./marketplace"]

