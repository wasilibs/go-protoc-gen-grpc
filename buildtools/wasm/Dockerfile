FROM ghcr.io/wasilibs/wasix-sdk:sha-fc94d60

RUN apt-get update && apt-get install -y binaryen git patch

ADD buildtools/wasm/version.txt version.txt
RUN git clone --recursive https://github.com/grpc/grpc.git --branch $(cat version.txt | awk '{$1=$1};1') /workspace

ENV CXXFLAGS "${CXXFLAGS} -O3 -pthread -D_WASI_EMULATED_MMAN -D_WASI_EMULATED_PROCESS_CLOCKS -DABSL_HAVE_MMAP"
ENV LDFLAGS $LDFLAGS -lwasi-emulated-process-clocks -lwasi-emulated-mman -Wl,--max-memory=4294967296

WORKDIR /workspace

ADD buildtools/wasm/patch-absl.txt patch-absl.txt

RUN cd third_party/abseil-cpp && patch -p1 < ../../patch-absl.txt

RUN cmake -S . -B build
RUN cmake --build build --target plugins

RUN for f in build/*_plugin; do wasm-opt -o "$f".wasm --low-memory-unused --flatten --rereloop --converge -O3 "$f"; done

CMD ["bash", "-c", "cp build/*_plugin.wasm /out/"]
