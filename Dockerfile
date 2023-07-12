FROM golang:1.20

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

EXPOSE 5000

ENTRYPOINT ["./apiserver"]