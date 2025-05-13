FROM golang:1.24.3-alpine

ENV BENTO4_PATH="/opt/bento4"
ENV PATH="$PATH:$BENTO4_PATH/bin"

RUN apk add --no-cache wget unzip bash ffmpeg

WORKDIR /tmp

# Baixa e extrai o Bento4 binário
RUN wget https://www.bok.net/Bento4/binaries/Bento4-SDK-1-6-0-641.x86_64-unknown-linux.zip && \
    unzip Bento4-SDK-1-6-0-641.x86_64-unknown-linux.zip && \
    mkdir -p $BENTO4_PATH && \
    cp -r Bento4-SDK-1-6-0-641.x86_64-unknown-linux/* $BENTO4_PATH && \
    rm -rf Bento4-SDK-1-6-0-641.x86_64-unknown-linux* && \
    chmod +x $BENTO4_PATH/bin/*

WORKDIR /go/src

# entrypoint opcional
ENTRYPOINT [ "tail", "-f", "/dev/null" ]
