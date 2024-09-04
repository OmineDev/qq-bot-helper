# cgo flags
# no need to use cgo for macos, windows, linux
MACOS_AMD64_CC:=/osxcross/bin/x86_64-apple-darwin20.4-clang
MACOS_ARM64_CC:=/osxcross/bin/arm64-apple-darwin20.4-clang
WINDOWS_AMD64_CC:=/usr/bin/x86_64-w64-mingw32-gcc
LINUX_AMD64_CC:=/usr/bin/gcc 
LINUX_ARM64_CC:=/usr/bin/aarch64-linux-gnu-gcc 

# need cgo for android, or network related part will not functional properly
ANDROID_ARM64_CC:=/root/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang
ANDROID_ARM64_CXX:=${ANDROID_ARM64_CC}
ANDROID_ARM64_CGO_FLAGS=CGO_ENABLED=1 CC=${ANDROID_ARM64_CC} CXX=ANDROID_ARM64_CXX
# end cgo flags

# define build target, project source, and exec name
PROJECT_DIR:= $(shell pwd)/../go-cqhttp
MAIN_DIR:= ${PROJECT_DIR} # or  ${PROJECT_DIR}/clis/remote_omega/access_point, etc.
PROJECT_NAME:= $(notdir ${MAIN_DIR})
SRCS_GO := $(foreach dir, $(shell find $(PROJECT_DIR) -type d), $(wildcard $(dir)/*.go $(dir)/*.c))
EXEC_PREFIX:=cqhttp-
# end define build target, project source, and exec name

# define go build flags
GO_BUILD_FLAGS_COMMON :=-trimpath -ldflags "-s -w"
# end define go build flags

# define release dir and output dir
RELEASE_DIR:=$(shell pwd)/binary
OUTPUT_DIR:=${RELEASE_DIR}
# end define release dir and output dir

# find raknet/conn.go under mod path
# MOD_CACHE_PATH:=$(shell go env GOMODCACHE)
# RAKNET_CONN_PATH:= $(shell find ${MOD_CACHE_PATH}/github.com/sandertv -name "conn.go" -type f)
# $(foreach file, ${RAKNET_CONN_PATH}, $(shell chmod 777 $(file) && sed -i "" "s/currentProtocol byte = */currentProtocol byte = 8/g" $(file)))
# end find raknet/conn.go under mod path



TYPE:=linux
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}
${${TYPE}_EXEC}_GO_CGO_FLAGS:=
${${TYPE}_EXEC}_TRIPLE:=GOOS=linux GOARCH=amd64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=linux-arm64
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}
${${TYPE}_EXEC}_GO_CGO_FLAGS:=
${${TYPE}_EXEC}_TRIPLE:=GOOS=linux GOARCH=arm64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=windows-x86
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}.exe
${${TYPE}_EXEC}_GO_CGO_FLAGS:=
${${TYPE}_EXEC}_TRIPLE:=GOOS=windows GOARCH=386
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=windows
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}.exe
${${TYPE}_EXEC}_GO_CGO_FLAGS:=
${${TYPE}_EXEC}_TRIPLE:=GOOS=windows GOARCH=amd64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=macos
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}
${${TYPE}_EXEC}_GO_CGO_FLAGS:=
${${TYPE}_EXEC}_TRIPLE:=GOOS=darwin GOARCH=amd64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=macos-arm64
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}
${${TYPE}_EXEC}_GO_CGO_FLAGS:=
${${TYPE}_EXEC}_TRIPLE:=GOOS=darwin GOARCH=arm64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=android
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}
${${TYPE}_EXEC}_GO_CGO_FLAGS:=${ANDROID_ARM64_CGO_FLAGS}
${${TYPE}_EXEC}_TRIPLE:=GOOS=android GOARCH=arm64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}


${OUTPUT_DIR}:
	@echo make output dir $@
	@mkdir -p $@
	

.PHONY: ${EXECS}
${EXECS}: ${OUTPUT_DIR}/${EXEC_PREFIX}%: ${OUTPUT_DIR} ${SRCS_GO}
	cd ${MAIN_DIR} && ${GO_CGO_FLAGS_COMMON} ${$@_GO_CGO_FLAGS} ${$@_TRIPLE}  go build ${GO_BUILD_FLAGS_COMMON} -o $@ 

	@md5sum $@ | cut -d' ' -f1 > $@.hash
	@cd /workspace/buildenv && go run /workspace/buildenv/utils/compressor/main.go -in $@ -out $@.brotli
	@echo "\033[32mbuild $@ Done \033[0m\t" `cat $@.hash`

execs:${EXECS}

checkout:
	cd ${PROJECT_DIR} && git checkout . && git checkout master

all: checkout ${EXECS}

upload: ${EXECS}
	cp ./binary/cqhttp-windows.exe.hash ./binary/cqhttp-windows.hash
	/workspace/buildenv/uploader-omega -l ./binary -r cqhttp_lagrange

clean:
	rm -f ${OUTPUT_DIR}/*