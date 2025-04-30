FROM alpine:20240329

ADD files/ /root/

RUN apk add fio iozone hdparm bash screen lshw inotify-tools iotop curl \
    && apk add gcc autoconf make g++ zlib-dev popt-dev \
    && apk cache purge \
    && chmod 0755 /root/mdbench.py \
    && wget https://www.samba.org/ftp/tridge/dbench/dbench-4.0.tar.gz \
    && tar xzf dbench-4.0.tar.gz \
    && cd dbench-4.0 \
    && ./autogen.sh && ./configure && make && make install \
    && cd .. && rm -rf dbench*

WORKDIR /root

ENTRYPOINT [ "/bin/bash" ]
