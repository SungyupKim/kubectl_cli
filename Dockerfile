FROM golang:alpine
RUN apk add --no-cache git make musl-dev go
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]
