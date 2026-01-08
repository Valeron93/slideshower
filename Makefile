.PHONY: dev dist

bun-install:
	@bun install --cwd frontend

dev: bun-install
	@mkdir -p tmp
	@bun run --cwd frontend dev --clearScreen false & air

dist: bun-install
	@mkdir -p tmp/dist
	@echo building frontend...
	@bun run --cwd frontend build
	@echo building application...
	@go build -o ./tmp/dist/slideshower -tags dist -ldflags="-s -w" ./cmd/slideshower
	@echo built: ./tmp/dist/slideshower

run-dist: dist
	@echo running dist...
	@./tmp/dist/slideshower