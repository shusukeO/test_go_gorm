.PHONY: run stop clean restart mysql-up mysql-down mysql-status go-run help

# デフォルトのDB_TYPE
DB_TYPE ?= mysql

# MySQLサーバーが起動しているかチェック
mysql-status:
	@if docker-compose ps mysql | grep -q "Up"; then \
		echo "MySQL is running"; \
		exit 0; \
	else \
		echo "MySQL is not running"; \
		exit 1; \
	fi

# MySQLサーバーを起動
mysql-up:
	@echo "Starting MySQL server..."
	@docker-compose up -d
	@echo "Waiting for MySQL to be ready..."
	@sleep 10
	@echo "MySQL is ready!"

# MySQLサーバーを停止
mysql-down:
	@echo "Stopping MySQL server..."
	@docker-compose down

# MySQLサーバーとデータを完全削除
clean:
	@echo "Stopping and removing MySQL server and data..."
	@docker-compose down -v
	@docker-compose rm -f
	@echo "Cleanup completed!"

# Goアプリケーションを実行
go-run:
	@echo "Running Go application with DB_TYPE=$(DB_TYPE)..."
	@DB_TYPE=$(DB_TYPE) go run main.go

# MySQLが起動していない場合は起動してからGoアプリケーションを実行
# MySQLが起動している場合はGoアプリケーションのみ実行
run:
	@if [ "$(DB_TYPE)" = "mysql" ]; then \
		if ! $(MAKE) mysql-status > /dev/null 2>&1; then \
			echo "MySQL is not running. Starting MySQL..."; \
			$(MAKE) mysql-up; \
		else \
			echo "MySQL is already running. Restarting Go application..."; \
		fi; \
	fi
	@$(MAKE) go-run DB_TYPE=$(DB_TYPE)

# MySQLサーバーを停止（データは保持）
stop:
	@$(MAKE) mysql-down

# MySQLサーバーを再起動
restart:
	@$(MAKE) mysql-down
	@$(MAKE) mysql-up

# ヘルプ
help:
	@echo "Available commands:"
	@echo "  make run [DB_TYPE=mysql|sqlite] - Run application (default: mysql)"
	@echo "  make stop                       - Stop MySQL server"
	@echo "  make clean                      - Stop MySQL and remove all data"
	@echo "  make restart                    - Restart MySQL server"
	@echo "  make mysql-up                   - Start MySQL server only"
	@echo "  make mysql-down                 - Stop MySQL server only"
	@echo "  make mysql-status               - Check MySQL status"
	@echo "  make go-run [DB_TYPE=mysql|sqlite] - Run Go app only"
	@echo ""
	@echo "Examples:"
	@echo "  make run                        - Start MySQL (if needed) and run with MySQL"
	@echo "  make run DB_TYPE=sqlite         - Run with SQLite"
	@echo "  make run DB_TYPE=mysql          - Start MySQL (if needed) and run with MySQL"
	@echo "  make stop                       - Stop MySQL (data preserved)"
	@echo "  make clean                      - Complete cleanup (data removed)"
