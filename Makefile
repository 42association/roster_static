# Makefile

# 'make run'コマンドでDockerコンテナをビルドし、起動します。
run:
	docker-compose up --build

# 'make down'コマンドでDockerコンテナを停止し、削除します。
down:
	docker-compose down

# 'make rebuild'コマンドでDockerコンテナを再ビルドし、起動します。
rebuild:
	docker-compose up --build --force-recreate

# 'make clean'コマンドでDocker関連の不要なデータをクリーンアップします。
clean:
	docker system prune -f
	docker volume prune -f
	docker network prune -f
