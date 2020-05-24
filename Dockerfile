FROM golang:alpine
COPY ./main.go /
WORKDIR /
RUN CGO_ENABLED=0 go build -o ./webserver

FROM scratch
WORKDIR /
COPY --from=0 /webserver .
EXPOSE 8080
ENTRYPOINT ["/webserver"]
