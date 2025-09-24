# Makefile

.PHONY: fix
fix:
	golangci-lint run --fix
