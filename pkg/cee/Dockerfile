FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    python3 python3-pip \
    gcc g++ \
    curl \
    time \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /sandbox
RUN useradd -m sandbox
USER sandbox

CMD ["/bin/sh"]