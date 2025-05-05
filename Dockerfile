FROM golang:1.24.1
COPY . /app
WORKDIR /app

COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

RUN apt update && apt install -y --no-install-recommends nano git  \
    && apt-get clean && rm -rf /var/lib/apt/lists/* \
RUN go mod download

EXPOSE 3080

ENTRYPOINT [ "/usr/local/bin/entrypoint.sh" ]