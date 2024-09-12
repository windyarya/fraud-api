FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /fraud-monitoring

FROM ubuntu:focal AS build-release-stage
WORKDIR /
COPY --from=build /fraud-monitoring /fraud-monitoring
EXPOSE 8081
ENTRYPOINT [ "/fraud-monitoring" ]