FROM alpine:3.5

MAINTAINER Kenny Chen

LABEL Description="Cybervision Beats plugin"

RUN apk update && \
    apk upgrade && \
    apk add \
        bash \
        ca-certificates \
    && rm -rf /var/cache/apk/*


RUN mkdir /plugin
COPY icube-cybervision-plugin.yml /plugin/
COPY fields.yml /plugin/
COPY ./release/linux/amd64/icube-cybervision-plugin /plugin/icube-cybervision-plugin

COPY settings.json /tmp/settings.json

WORKDIR /plugin

ENTRYPOINT ["/plugin/icube-cybervision-plugin", "-e", "-d", "*"]
