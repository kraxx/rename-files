NAME =	rename_this

DEPS =	

OUTPUT_DIR =	./
OUTPUT_BINARY =	${OUTPUT_DIR}${NAME}

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


.PHONY: install build serve test clean heroku docker re # invalidate these commands if they exist outside this script
.SILENT: # Prepends everything with @ (command executed without printing to stdout)
all: install build
install:
	echo "${YELLOW_LIGHT_BOLD}Installing dependencies${END_COLOUR}"
	go get ${DEPS}
build:
	echo "${YELLOW_LIGHT_BOLD}Building binary${END_COLOUR}"
	go build -o ${OUTPUT_BINARY} -ldflags "-X main.version=$(TAG)" .
clean:
	echo "${YELLOW_LIGHT_BOLD}Cleaning installations and binary${END_COLOUR}"
	go clean
	rm -f ${OUTPUT_BINARY}

re: clean all
