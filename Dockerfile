FROM golang:1.22.5-nanoserver-ltsc2022 AS build
WORKDIR C:/src
COPY . .
RUN go build -o q394-client.exe .

FROM mcr.microsoft.com/windows/nanoserver:ltsc2022
COPY --from=build C:/src/q394-client.exe C:/q394-client.exe
ENTRYPOINT ["C:\\q394-client.exe"]
