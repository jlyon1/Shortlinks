FROM alpine:3.6

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apk add --no-cache \
    ca-certificates \
    musl-dev \
    git \
    go \
    tzdata


RUN mkdir -p "$GOPATH/src/docs/" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/docs
RUN git clone https://github.com/jlyon1/docs.git "./"
RUN go get

COPY . .

RUN go build -o docs

CMD ["./docs"]
