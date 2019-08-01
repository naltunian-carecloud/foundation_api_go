
FROM golang
ARG token
RUN git config --global url."https://service-carecloud:${token}@github.com/".insteadOf "https://github.com/"
WORKDIR /go/src/github.com/CareCloud/foundation_api_go/
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build .

FROM alpine
ENV CONFIG_HTTP_PORT=80
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /go/src/github.com/CareCloud/foundation_api_go/ .
ENTRYPOINT ["/foundation_api_go"]