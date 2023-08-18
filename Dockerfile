# Build stage
FROM golang:alpine as build

WORKDIR /build
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o test-task /build/cmd/server/main.go

# Deploy stage
FROM alpine:latest
COPY --from=build /build/ /
EXPOSE 8000
CMD ["/test-task"]
