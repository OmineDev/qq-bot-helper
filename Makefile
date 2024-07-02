# c/c++ toolchain
CGO_ENABLED:=1
# if CGO_ENABLED then set c/c++ compilers
ifeq ($(CGO_ENABLED),1)
	MACOS_AMD64_CC:=/usr/bin/clang 
	MACOS_ARM64_CC:=/usr/bin/clang
	WINDOWS_AMD64_CC:=/opt/homebrew/bin/x86_64-w64-mingw32-gcc
	WINDOWS_X86_CC:=/opt/homebrew/bin/i686-w64-mingw32-gcc
	LINUX_AMD64_CC:=/opt/homebrew/bin/x86_64-unknown-linux-gnu-gcc
	LINUX_ARM64_CC:=/opt/homebrew/bin/aarch64-unknown-linux-gnu-gcc
	ANDROID_NDK_HOME:=$(shell brew --prefix)/share/android-ndk
	ANDROID_ARM64_CC:=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang
endif
# end c/c++ toolchain

# define build target, project source, and exec name
PROJECT_DIR:= $(shell pwd)/../go-cqhttp
MAIN_DIR:= ${PROJECT_DIR} # or  ${PROJECT_DIR}/clis/remote_omega/access_point, etc.
PROJECT_NAME:= $(notdir ${MAIN_DIR})
SRCS_GO := $(foreach dir, $(shell find $(PROJECT_DIR) -type d), $(wildcard $(dir)/*.go $(dir)/*.c))
EXEC_PREFIX:=cqhttp-
# end define build target, project source, and exec name

# define go build flags
GO_CGO_FLAGS_COMMON:=CGO_ENABLED=0
ifeq ($(CGO_ENABLED),1)
	GO_CGO_FLAGS_COMMON :=CGO_CFLAGS=${CGO_DEF} CGO_ENABLED=1
endif
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
${${TYPE}_EXEC}_CC:=${LINUX_AMD64_CC}
${${TYPE}_EXEC}_TRIPLE:=GOOS=linux GOARCH=amd64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=linux-arm64
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}
${${TYPE}_EXEC}_CC:=${LINUX_ARM64_CC}
${${TYPE}_EXEC}_TRIPLE:=GOOS=linux GOARCH=arm64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=windows-x86
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}.exe
${${TYPE}_EXEC}_CC:=${WINDOWS_X86_CC}
${${TYPE}_EXEC}_TRIPLE:=GOOS=windows GOARCH=386
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=windows
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}.exe
${${TYPE}_EXEC}_CC:=${WINDOWS_AMD64_CC}
${${TYPE}_EXEC}_TRIPLE:=GOOS=windows GOARCH=amd64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=macos
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}
${${TYPE}_EXEC}_CC:=${MACOS_AMD64_CC}
${${TYPE}_EXEC}_TRIPLE:=GOOS=darwin GOARCH=amd64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=macos-arm64
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}
${${TYPE}_EXEC}_CC:=${MACOS_ARM64_CC}
${${TYPE}_EXEC}_TRIPLE:=GOOS=darwin GOARCH=arm64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}

TYPE:=android
${TYPE}_EXEC:=${OUTPUT_DIR}/${EXEC_PREFIX}${TYPE}
${${TYPE}_EXEC}_CC:=${ANDROID_ARM64_CC}
${${TYPE}_EXEC}_TRIPLE:=GOOS=android GOARCH=arm64
${TYPE}: ${${TYPE}_EXEC}
EXECS:=${EXECS} ${${TYPE}_EXEC}


${OUTPUT_DIR}:
	@echo make output dir $@
	@mkdir -p $@
	

.PHONY: ${EXECS}
${EXECS}: ${OUTPUT_DIR}/${EXEC_PREFIX}%: ${OUTPUT_DIR} ${SRCS_GO}
		@${GO_CGO_FLAGS_COMMON} ${$@_TRIPLE} CC=${$@_CC}  go build ${GO_BUILD_FLAGS_COMMON} -o $@ ${MAIN_DIR}

	@md5sum $@ | cut -d' ' -f1 > $@.hash
	@go run ../utils/compressor/main.go -in $@ -out $@.brotli
	@echo "\033[32mbuild $@ Done \033[0m\t" `cat $@.hash`

execs:${EXECS}

checkout:
	cd ${PROJECT_DIR} && git checkout . && git checkout master

all: checkout ${EXECS}

upload: ${EXECS}
	cp ./binary/cqhttp-windows.exe.hash ./binary/cqhttp-windows.hash
	../uploader-omega -l ./binary -r cqhttp

clean:
	rm -f ${OUTPUT_DIR}/*