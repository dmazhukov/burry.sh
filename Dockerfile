#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/github.com/dmazhukov/burry.sh
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...
# # RUN go-wrapper download   # "go get -d -v ./..."
# # RUN go-wrapper install    # "go install -v ./..."

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/burry.sh /burry

ENTRYPOINT ["/burry"] 
LABEL Name=burry.sh Version=1.0
