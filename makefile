ifdef COMSPEC
	EXE_EXT := .exe
else
	EXE_EXT := 
endif

.PHONY: build
build:
	go build -o pokered$(EXE_EXT) ./cmd/main.go