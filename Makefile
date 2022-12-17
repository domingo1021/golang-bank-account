createdb:
	mysql --user=$(MYSQL_USER) --password=$(MYSQL_PASSWORD) -e 'CREATE DATABASE test_make'

.PHONY: createdb