FROM alpine:3.13


WORKDIR /opt/myapp/
COPY linuxExe .
COPY ./templates ./templates

RUN chmod +x linuxExe

EXPOSE 80

ENTRYPOINT [ "./linuxExe" ]