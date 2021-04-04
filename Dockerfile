# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git gcc ca-certificates
COPY go.mod go.sum main.go /src/
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main

# final stage
FROM scratch
COPY --from=build-env /src/main /
ENTRYPOINT ["/main"]