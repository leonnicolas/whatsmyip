FROM golang:alpine as build
ARG ARCH=amd64
COPY . /wmip
WORKDIR /wmip
RUN CGO_ENABLED=0 GOARCH=$ARCH GOOS=linux go build  -o ./wmip

FROM scratch
WORKDIR /
COPY --from=build /wmip/wmip .
EXPOSE 8080
ENTRYPOINT ["/wmip"]
