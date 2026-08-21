FROM ubuntu

RUN apt update && \
    apt install -y gdb-multiarch llvm-20 wget

RUN wget https://go.dev/dl/go1.26.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go*.tar.gz && \
    rm go*.tar.gz

RUN wget https://github.com/tinygo-org/tinygo/releases/download/v0.41.1/tinygo_0.41.1_amd64.deb && \
	dpkg -i tinygo*.deb && \
    rm tinygo*.deb

ENV PATH="$PATH:/usr/local/go/bin"
