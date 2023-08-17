FROM golang:1.20

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download

COPY ../cmd cmd/
COPY ../src src/

RUN ls -l
RUN go build -o ./app ./cmd

EXPOSE 8000

ENV GIN_MODE=release

CMD ["go", "run", "./cmd"]