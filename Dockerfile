FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN apk add --no-cache curl git
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags=jsoniter -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
ENV GIN_MODE=release
COPY --from=builder /build/main /app/
COPY --from=builder /build/conf /app/conf
WORKDIR /app
CMD ["./main"]