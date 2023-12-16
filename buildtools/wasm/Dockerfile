FROM ghcr.io/wasilibs/wasix-sdk:sha-fc94d60

RUN apt-get update && apt-get install -y binaryen git patch

RUN git clone --recursive https://github.com/grpc/grpc.git --branch v1.60.0 /workspace

ENV CXXFLAGS "${CXXFLAGS} -O3 -pthread -D_WASI_EMULATED_MMAN -D_WASI_EMULATED_PROCESS_CLOCKS"
ENV LDFLAGS $LDFLAGS -lwasi-emulated-process-clocks -lwasi-emulated-mman -Wl,--max-memory=4294967296

WORKDIR /workspace
RUN cmake -S . -B build
RUN cmake --build build --target plugins

RUN for f in build/*_plugin; do wasm-opt -o "$f".wasm --low-memory-unused --flatten --rereloop --converge -O3 "$f"; done

CMD ["bash", "-c", "cp build/*_plugin.wasm /out/"]
