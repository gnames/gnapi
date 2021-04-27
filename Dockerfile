FROM alpine:3.13

MAINTAINER Dmitry Mozzherin

ENV LAST_FULL_REBUILD 2021-04-27

WORKDIR /bin

COPY ./gnapi/gnapi /bin

ENTRYPOINT [ "gnapi" ]

CMD ["-p", "8888"]
