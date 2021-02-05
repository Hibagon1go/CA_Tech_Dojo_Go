FROM golang:latest
# このファイルからビルドされるDockerImageのベースとして、DockerHub上にある最新のgolangのDockerImageを指定

COPY src/api /go/src/api 
# このコンテナ内で走らせたいファイルのあるディレクトリを、Dockerコンテナ内にコピー
# これにより走らせたいファイルをDockerImageに組み込み、コンテナ内で使用可能に 

WORKDIR /go/src/api/
# docker ** という命令を行うディレクトリを指定

RUN go env -w GO111MODULE=on \
  && go env -w GOSUMDB=off 
# $GOPATHの外で開発できるように設定
# $GOSUMDBをoffにし、golang.orgに辿り着けるように設定

RUN go get -u github.com/gin-gonic/gin \ 
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql \
  && apt-get update && apt-get install -y mariadb-client
# DockerImageをビルドする際にコンテナ内で実行されるコマンド
# 必要なモノを諸々コンテナ内に取り込む

CMD ["sh", "./start_app.sh"]
# このコンテナが起動する際に1度実行されるコマンド