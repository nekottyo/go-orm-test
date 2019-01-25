up:
	docker network create mariadb || true
	docker rm -f mariadb || true
	docker run --name mariadb --net mariadb -e MYSQL_ROOT_PASSWORD=hogehoge -p 3306:3306 -v $(CURDIR)/sql:/docker-entrypoint-initdb.d -d mariadb

down:
	docker rm -f mariadb
	docker network rm mariadb
