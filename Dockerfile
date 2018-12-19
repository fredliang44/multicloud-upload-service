FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags=jsoniter -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
COPY --from=builder /build/conf /app/conf
WORKDIR /app
CMD ["./main"]