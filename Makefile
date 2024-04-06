BINARY_FILE = ./binary/app
all: compile run

compile:
	go build -v -o ${BINARY_FILE} cmd/todo-app/main.go

run:
	${BINARY_FILE}


