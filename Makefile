# 静的解析
verify:
	@staticcheck ./... && go vet ./...

ut:
	@go test ./... -coverprofile=./test_results/cover.out
	@go tool cover -html=./test_results/cover.out -o ./test_results/cover.html
	@open ./test_results/cover.html

serve:
	@docker-compose -f build/docker-compose.yml up -d

down:
	@docker-compose -f build/docker-compose.yml down
