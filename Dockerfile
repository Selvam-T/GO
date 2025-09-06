FROM golang:1.25

WORKDIR /usr/src/app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        nano \
        curl \
        git \
        && rm -rf /var/lib/apt/lists/*

CMD ["bash"]
