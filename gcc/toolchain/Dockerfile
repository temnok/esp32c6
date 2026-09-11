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

RUN git clone --depth 1 --single-branch --branch 2026.08.27 \
    https://github.com/riscv/riscv-gnu-toolchain toolchain-source && \
    cd toolchain-source && \
    ./configure --prefix=/opt/riscv --enable-multilib && \
    make --jobs=$(nproc) && \
    cd .. && \
    rm -rf toolchain-source

ENV PATH="$PATH:/opt/riscv/bin"

RUN apt update && \
    apt install -y \
        binutils-dev \
        build-essential \
        g++-multilib \
        gcc-multilib \
        genromfs \
        gettext \
        kconfig-frontends \
        libexpat1-dev \
        libisl-dev \
        libncurses5-dev \
        libncursesw5-dev \
        picocom \
        pkg-config \
        u-boot-tools \
        util-linux \
        xxd

RUN pip install esptool --break-system-packages
