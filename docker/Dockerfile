FROM golang:1.17-bullseye

RUN apt-get update && \
    apt-get install -y inotify-tools

COPY docker/entrypoint.sh /
RUN chmod 755 /entrypoint.sh
ENTRYPOINT [ "/entrypoint.sh"]
