FROM golang:alpine
RUN apk add --no-cache git make musl-dev go
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main . 
EXPOSE 3010
CMD ["/app/main"]
