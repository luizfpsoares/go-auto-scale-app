FROM golang:1.24.2-alpine AS build
WORKDIR /opt/app
COPY . .
RUN go get ./cmd/api
RUN env GOOS=linux GOARCH=amd64 go build -o auto-scale-app cmd/api/main.go

FROM alpine:3.21.3
WORKDIR /opt/app
COPY --from=build /opt/app/auto-scale-app .
ENTRYPOINT [ "./auto-scale-app" ]