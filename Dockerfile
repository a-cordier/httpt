
FROM golang:1-alpine as build

WORKDIR /opt
COPY src .
RUN go build -o httpt .

FROM alpine:latest as run

COPY --from=build /opt/httpt .

# Openshift compliance
ENV USER_ID=1001
RUN chmod +x /httpt
USER $USER_ID

EXPOSE 8080

CMD ["/httpt"]
