NAME = rename_this

OUTPUT_DIR =	./
OUTPUT_BINARY =	${OUTPUT_DIR}${NAME}

TEST_DIR =		test
TEST_SCRIPT =	generate20files.sh

RED =				\033[31m
GREEN =				\033[32m
BLUE =				\033[34m
YELLOW =			\033[33m
MAGENTA =			\033[35m
GREY =				\033[37m
GREEN_LIGHT =		\033[92m
YELLOW_LIGHT =		\033[93m
YELLOW_BOLD =		\033[1;33m
YELLOW_LIGHT_BOLD =	\033[1;93m
MAGENTA_LIGHT =		\033[95m
BLINK =				\033[5m
GREEN_LIGHT_BLINK =	\033[5;92m
END_COLOUR =		\033[0m


.PHONY: build windows clean test re # invalidate these commands if they exist outside this script
.SILENT: # Prepends everything with @ (command executed without printing to stdout)
all: build
build:
	echo "${YELLOW_LIGHT_BOLD}Building binary${END_COLOUR}"
	go build -o ${OUTPUT_BINARY} -ldflags "-X main.version=$(TAG)" .
windows:
	echo "${YELLOW_LIGHT_BOLD}Building binary for ${GREEN_LIGHT_BLINK}WINDOWS${END_COLOUR}"
	GOOS=windows GOARCH=amd64 go build -o ${OUTPUT_BINARY}.exe -ldflags "-X main.version=$(TAG)" .
test: build
	echo "${YELLOW_LIGHT_BOLD}Generating and renaming test files${END_COLOUR}"
	mkdir -p ${TEST_DIR} && cd ${TEST_DIR} && ../${TEST_SCRIPT} && ../${OUTPUT_BINARY}
clean:
	echo "${YELLOW_LIGHT_BOLD}Cleaning installations and binary${END_COLOUR}"
	go clean
	rm -f ${OUTPUT_BINARY} ${OUTPUT_BINARY}.exe
	rm -rf ${TEST_DIR}
re: clean all
