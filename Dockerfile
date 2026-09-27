FROM golang:1.22.5 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /q394-client .

FROM gcr.io/distroless/static-debian12
COPY --from=build /q394-client /q394-client
ENTRYPOINT ["/q394-client"]
