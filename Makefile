format:
	@find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/Riotam/sample-mvc-api-go"
verify:
	@staticcheck ./... && go vet ./...
unit-test:
	@go test ./... -coverprofile=./test_results/cover.out && go tool cover -html=./test_results/cover.out -o ./test_results/cover.html
serve:
	@docker-compose -f build/docker-compose.yml up -d
down:
	@docker-compose -f build/docker-compose.yml down