# Makefile

# デフォルトのターゲット
all: up

# コンテナをビルドして起動
up:
	docker-compose up --build

# コンテナを停止
down:
	docker-compose down

# コンテナのログを表示
logs:
	docker-compose logs -f

# コンテナをビルド
build:
	docker-compose build

# ヘルプメッセージを表示
help:
	@echo "make - デフォルトのターゲット (同じく 'make up') を実行して、コンテナをビルドして起動します。"
	@echo "make down - コンテナを停止します。"
	@echo "make logs - コンテナのログを表示します。"
	@echo "make build - コンテナをビルドします。"
