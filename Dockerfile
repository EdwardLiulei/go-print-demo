From golang:1.17.7-stretch as gobuilder
WORKDIR /demo
COPY . /demo
RUN go build -o /demo/print

FROM amazonlinux:latest
WORKDIR /demo
COPY --from=gobuilder /demo/print /demo/print
CMD ["/demo/print", "--p", "hello world"]