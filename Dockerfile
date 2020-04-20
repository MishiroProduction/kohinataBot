FROM golang:1.13 as build
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0
WORKDIR /usr/src/kohinataBot
COPY . .
RUN go build -o kohinataBot main.go

FROM alpine:3.11.5
COPY --from=build /usr/src/kohinataBot/kohinataBot /kohinataBot
RUN apk --no-cache add ca-certificates
ENTRYPOINT [ "/kohinataBot" ]