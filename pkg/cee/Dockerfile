FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    python3 python3-pip \
    gcc g++ \
    curl \
    time \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /sandbox

COPY run.sh /sandbox/run.sh
RUN chmod +x /sandbox/run.sh

ENTRYPOINT ["/sandbox/run.sh"]
CMD ["/sandbox/code"]