FROM ghcr.io/webassembly/wasi-sdk:wasi-sdk-25

RUN apt-get update && apt-get install -y binaryen git patch

ADD buildtools/wasm/version.txt version.txt
RUN git clone --recurse-submodules --shallow-submodules https://github.com/grpc/grpc.git --branch $(cat version.txt | awk '{$1=$1};1') /workspace

ENV CMAKE_TOOLCHAIN_FILE=/opt/wasi-sdk/share/cmake/wasi-sdk-pthread.cmake
ENV CFLAGS="-D_WASI_EMULATED_MMAN -D_WASI_EMULATED_PROCESS_CLOCKS -D_WASI_EMULATED_SIGNAL -DABSL_HAVE_MMAP"
ENV CXXFLAGS="$CFLAGS -fno-exceptions"
ENV LDFLAGS="-lwasi-emulated-process-clocks -lwasi-emulated-mman -lwasi-emulated-signal -Wl,--max-memory=4294967296 -Wl,--global-base=1024"

WORKDIR /workspace

ADD buildtools/wasm/patch-absl.txt patch-absl.txt
RUN cd third_party/abseil-cpp && patch -p1 < ../../patch-absl.txt
ADD buildtools/wasm/patch-protobuf.txt patch-protobuf.txt
RUN cd third_party/protobuf && patch -p1 < ../../patch-protobuf.txt

# libprotoc, the library that code generator plugins link against, includes the command line interface, including
# subprocess execution. Code generator plugins never need this, so we can remove the files to allow the project to
# build with wasi-sdk.
RUN rm third_party/protobuf/src/google/protobuf/compiler/subprocess.* third_party/protobuf/src/google/protobuf/compiler/command_line_interface.*
RUN sed -i '/src\/google\/protobuf\/compiler\/subprocess\./d' third_party/protobuf/src/file_lists.cmake
RUN sed -i '/src\/google\/protobuf\/compiler\/command_line_interface\./d' third_party/protobuf/src/file_lists.cmake

RUN cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_C_FLAGS="$CFLAGS" -DCMAKE_CXX_FLAGS="$CXXFLAGS" -DCMAKE_EXE_LINKER_FLAGS="$LDFLAGS" -S . -B build
RUN cmake --build build --target plugins -j 2

RUN for f in build/*_plugin; do wasm-opt -o "$f".wasm --low-memory-unused --flatten --rereloop --converge -O3 "$f"; done

CMD ["bash", "-c", "cp build/*_plugin.wasm /out/"]
