# Nome do alvo para rodar todos os testes
.PHONY: run-tests

# Alvo para rodar os testes
run-tests:
	go test ./... -v
