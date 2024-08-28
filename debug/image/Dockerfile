FROM alpine:20240329

#ADD files/ /root/

RUN apk add fio iozone hdparm bash screen lshw inotify-tools iotop curl \
    && apk add tcpdump mtr traceroute iputils-tracepath inetutils-telnet \
    && apk add iptables bind-tools \
    && apk cache purge

COPY files/ /files/

WORKDIR /root

ENTRYPOINT [ "/bin/bash" ]
