FROM golang:1.24.5 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /q394-policy-client .

FROM gcr.io/distroless/static-debian12
COPY --from=build /q394-policy-client /q394-policy-client
ENTRYPOINT ["/q394-policy-client"]
