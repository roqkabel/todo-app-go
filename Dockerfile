FROM golang:1.22.0-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum  ./


# RUN go mod download 
RUN GOPROXY="https://goproxy.io" go mod download
RUN go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app cmd/todo-app/main.go

CMD ["app"]

EXPOSE 5050


