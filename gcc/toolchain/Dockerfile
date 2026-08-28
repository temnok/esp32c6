FROM gcc:16.2.0-trixie

RUN apt update && \
    apt install -y \
        autoconf \
        automake \
        autotools-dev \
        bc \
        bison \
        bzip2 \
        ccache \
        cmake \
        cpio \
        curl \
        device-tree-compiler \
        e2fsprogs \
        flex \
        gawk \
        git \
        gperf \
        libelf-dev \
        libexpat-dev \
        libglib2.0-dev \
        libgmp-dev \
        libmpc-dev \
        libmpfr-dev \
        libslirp-dev \
        libssl-dev \
        libtool \
        make \
        ninja-build \
        patchutils \
        python3 \
        python3-pip \
        python3-tomli \
        sudo \
        texinfo \
        wget \
        xz-utils \
        zlib1g-dev \
    ;


RUN groupadd -r user && \
    useradd -m -r -g user user

USER user
WORKDIR /home/user

RUN git clone --depth 1 --single-branch --branch 2026.08.27 \
    https://github.com/riscv/riscv-gnu-toolchain toolchain-source && \
    cd toolchain-source && \
    ./configure --prefix=/home/user/toolchain --with-arch=rv32imac_zicsr --with-abi=ilp32 && \
    make --jobs=$(nproc) && \
    cd .. && \
    rm -rf toolchain-source

ENV PATH="$PATH:/home/user/toolchain/bin"
