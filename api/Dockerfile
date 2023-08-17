FROM golang:1.20

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download

COPY ../cmd cmd/
COPY ../src src/

RUN go install github.com/cosmtrek/air@v1.27.3

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]