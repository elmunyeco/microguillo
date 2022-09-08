FROM golang:1.18.1-alpine as build-env 

RUN apk add --no-cache git

WORKDIR /src

COPY . .

RUN go mod download 

RUN CGO_ENABLED=0 go build -o ./out/main /src/main.go


FROM scratch 
COPY --from=build-env /src/out/main /main
EXPOSE 8080
CMD  ["./main"]