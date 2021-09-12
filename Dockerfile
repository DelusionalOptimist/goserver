# build using golang docker image
FROM golang:1.17 as builder
WORKDIR /goserver
ADD . .
RUN go build -o /goserver .

# store it in a lightweight distroless image
FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /goserver/goserver .
ENTRYPOINT ["/goserver"]
