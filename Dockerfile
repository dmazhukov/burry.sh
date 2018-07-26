FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY burry ./burry
ENTRYPOINT ["./burry"] 
LABEL Name=burry.sh Version=1.0
