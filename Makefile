NAME=bothan
OUT=bin
BUILD=go build
STRIP=-s -w
LINUX_LDFLAGS=-ldflags "${STRIP}"
WIN_LDFLAGS=-ldflags "${STRIP} -H windowsgui"

all: linux64 windows64 macos64 linux32 macos32 windows32 linux_arm linux_arm64

linux: linux64 linux32 linux_arm linux_arm64
windows: windows32 windows64
macos: macos32 macos64

linux64:
	$(eval GOOS=linux)
	$(eval GOARCH=amd64)
	GOOS=linux GOARCH=amd64 ${BUILD} ${LINUX_LDFLAGS} -o ${OUT}/${GOARCH}/${NAME}.${GOOS}

linux32:
	$(eval GOOS=linux)
	$(eval GOARCH=386)
	GOOS=linux GOARCH=386 ${BUILD} ${LINUX_LDFLAGS} -o ${OUT}/${GOARCH}/${NAME}.${GOOS}

linux_arm64:
	$(eval GOOS=linux)
	$(eval GOARCH=arm64)
	GOOS=linux GOARCH=arm64 ${BUILD} ${LINUX_LDFLAGS} -o ${OUT}/${GOARCH}/${NAME}.${GOOS}


linux_arm:
	$(eval GOOS=linux)
	$(eval GOARCH=arm)
	GOOS=linux GOARCH=arm ${BUILD} ${LINUX_LDFLAGS} -o ${OUT}/${GOARCH}/${NAME}.${GOOS}

windows64:
	$(eval GOOS=windows)
	$(eval GOARCH=amd64)
	GOOS=windows GOARCH=amd64 ${BUILD} ${WIN_LDFLAGS} -o ${OUT}/${GOARCH}/${NAME}.${GOOS}.exe

windows32:
	$(eval GOOS=windows)
	$(eval GOARCH=386)
	GOOS=windows GOARCH=386 ${BUILD} ${WIN_LDFLAGS} -o ${OUT}/${GOARCH}/${NAME}.${GOOS}.exe

macos64:
	$(eval GOOS=darwin)
	$(eval GOARCH=amd64)
	GOOS=darwin GOARCH=amd64 ${BUILD} ${LINUX_LDFLAGS} -o ${OUT}/${GOARCH}/${NAME}.${GOOS}

macos32:
	$(eval GOOS=darwin)
	$(eval GOARCH=386)
	GOOS=darwin GOARCH=386 ${BUILD} ${LINUX_LDFLAGS} -o ${OUT}/${GOARCH}/${NAME}.${GOOS}

clean:
	rm -rf ${OUT}
