FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
ADD conf /build/conf
WORKDIR /build
RUN apk add --no-cache curl git
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags=jsoniter -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM alpine
ENV GIN_MODE=release
RUN apk update && \
    apk add --no-cache ca-certificates
COPY --from=builder /build/main /app/
COPY --from=builder /build/conf /app/conf
WORKDIR /app
CMD ["./main"]