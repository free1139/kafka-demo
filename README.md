
# Install
Install docker
``` shell
sudo aptitude install docker.io
```

Using kafka in docker
[see](https://kafka.apache.org/quickstart)
```shell
sudo docker pull apache/kafka:3.8.0
sudo docker run -p 9092:9092 -d apache/kafka:3.8.0
sudo docker ps
```

Using mysql in docker
```shell
sudo docker pull mysql
sudo docker run -e MYSQL_ROOT_PASSWORD=123456 -p 3307:3306 -d mysql:latest
sudo docker ps

# login mysql and create dateabase named push
sudo aptitude install mysql-common
mysql -P 3307 -uroot -p
create database push;
exit
```

# Run a demo
shell 1
```shell
cd cmd/push-server
go run main.go
```

shell 2
```shell
cd cmd/push-client
go run main.go test.csv
```

