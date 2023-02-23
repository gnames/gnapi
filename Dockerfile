FROM alpine:3.17

MAINTAINER Dmitry Mozzherin

ENV LAST_FULL_REBUILD 2023-02-23

WORKDIR /bin

COPY ./out/bin/gnapi /bin

ENTRYPOINT [ "gnapi" ]

CMD ["-p", "8888"]
