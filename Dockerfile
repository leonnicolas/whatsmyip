FROM scratch
COPY ./simple_webserver /
EXPOSE 8080
ENTRYPOINT ["/simple_webserver"]
