# Go Database Connection

## How to Install Project
### Run with docker image
```bash
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=sckshuhari -d mysql:5.7
```

### Install MySql Workbench
```bash
https://dev.mysql.com/downloads/workbench/?utm_source=tuicool
```

### Install go mysql driver library
```bash
go get github.com/go-sql-driver/mysql
```