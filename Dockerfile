FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o server server.go

FROM scratch
COPY --from=build-env /src/server /
EXPOSE 80
ENTRYPOINT ["/server"]
