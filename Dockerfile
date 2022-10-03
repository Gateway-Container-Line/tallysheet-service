FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o tallysheet-service

RUN export GO111MODULE=on

ENTRYPOINT ["/app/tallysheet-service"]

CMD ["go", "run" , "main.go"]