xhappen

proto:

	protoc --proto_path=./api --proto_path=./third_party --go_out=paths=source_relative:./api --go-grpc_out=paths=source_relative:./api --go-http_out=paths=source_relative:./api metadata/metadata.proto

	protoc --proto_path=./third_party --go_out=paths=source_relative:./ errors/errors.proto


	protoc --proto_path=./api --proto_path=./third_party --go_out=paths=source_relative:./api --go-grpc_out=paths=source_relative:./api --go-http_out=paths=source_relative:./api metadata/metadata.proto


	protoc --proto_path=. \
         --proto_path=./third_party \
         --go_out=paths=source_relative:. \
         --go-errors_out=paths=source_relative:. \
         api/basic/v1/error_reason.proto

	protoc --proto_path=. \
         --proto_path=./third_party \
         --go_out=paths=source_relative:. \
         --go-errors_out=paths=source_relative:. \
         --go-grpc_out=paths=source_relative:. \
         --go-http_out=paths=source_relative:. \
         api/gateway/v1/service.proto


    protoc --proto_path=. \
         --proto_path=./third_party \
         --go_out=paths=source_relative:. \
         --go-errors_out=paths=source_relative:. \
         app/xcache/internal/conf/conf.proto
     
    kratos proto server api/helloworld/v1/demo.proto -t internal/service

docker:
    docker run -itd --name redis -p 6379:6379 redis --requirepass pwd9527

    docker run --name mysql8 \
 		-e  MYSQL_ROOT_PASSWORD=pwd9527  -d -i -p 3306:3306 mysql:latest  --lower_case_table_names=1

 	docker run -d --name mongo -p 27017:27017\
        -e MONGO_INITDB_ROOT_USERNAME=root \
        -e MONGO_INITDB_ROOT_PASSWORD=pwd9527 \
        mongo


    docker network create etcd-net --driver bridge

    docker run -d --name etcd-server \
    --network etcd-net \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest

    docker run -it --rm \
    --network etcd-net \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    bitnami/etcd:latest etcdctl --endpoints http://etcd-server:2379 set /message hello

    docker network create kafka-net --driver bridge

    docker run -d --name kafka-server \
    --network kafka-net \
    --publish 9092:9092 \
    -e ALLOW_PLAINTEXT_LISTENER=yes \
    bitnami/kafka:latest

    docker run -it --rm \
    --network kafka-net \
    bitnami/kafka:latest kafka-topics.sh --list  --bootstrap-server kafka-server:9092

    docker run -it --rm \
    --network kafka-net \
    bitnami/kafka:latest kafka-topics.sh --create --partitions 1  --topic smscode --bootstrap-server kafka-server:9092


    docker run -it --rm \
    --network kafka-net \
    bitnami/kafka:latest kafka-topics.sh --describe  --topic smscode --bootstrap-server kafka-server:9092

    docker run -it --rm \
    --network kafka-net \
    bitnami/kafka:latest kafka-topics.sh --delete  --topic smscode --bootstrap-server kafka-server:9092


    docker run -it --rm --network kafka-net bitnami/kafka:latest kafka-console-producer.sh --topic smscode --bootstrap-server kafka-server:9092


    docker run -it --rm  --network kafka-net bitnami/kafka:latest kafka-console-consumer.sh --topic smscode --bootstrap-server kafka-server:9092


    docker run -d  -p 9000:9000 -p 9001:9001 \
  quay.io/minio/minio server /data --console-address ":9001"

  ANnKuRRWMbmEp0o8JdfX/quckuk-vosJor-9mapva