FROM alpine:3.21

LABEL maintainer="Dmitry Mozzherin"

ENV LAST_FULL_REBUILD=2025-10-01

WORKDIR /bin

COPY ./out/bin/gnapi /bin

ENTRYPOINT [ "gnapi" ]

CMD ["-p", "8888"]
