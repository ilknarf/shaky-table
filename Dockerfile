FROM golang:alpine
RUN apk add --no-cache --update gcc musl-dev
WORKDIR /build
COPY . .
RUN go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=0 /build/server ./
EXPOSE 8080
CMD ["./server"]
