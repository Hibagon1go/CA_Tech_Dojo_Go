# GO言語を用いたガチャAPI作成
## 実際に使用した技術
- GO(gin+gorm)
- MySQL
- Docker
- (Swagger)

## ガチャAPIの仕様
まずdocker-compose up -d などでコンテナ達を起動。
### ユーザー関連
#### ユーザー情報作成(POST)
以下のようにHTTPヘッダにjson形式でユーザー名を入れ渡すと、登録できる。<br>
例) curl -X POST "http://localhost:8080/user/create" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{  \"name\": \"Hibagon1go\"}" <br>
この時、以下のjson形式でトークンを返す。<br>
例) { "token" : "..." }
#### ユーザー情報取得(GET)
以下のようにHTTPヘッダに以下の形でトークンを入れ渡すと、そのユーザー名を取得できる。<br>
例) curl -X GET "http://localhost:8080/user/get" -H  "accept: application/json" -H  "x-token: ... " <br>
この時、以下のjson形式でユーザー情報を返す。<br>
例) { "name" : "Hibagon1go" }
#### ユーザー情報更新(PUT)
以下のようにHTTPヘッダに以下の形でトークンと、json形式で更新したいユーザー名を入れ渡すと、そのユーザー名を更新できる。<br>
例) curl -X PUT "http://localhost:8080/user/update" -H  "accept: application/json" -H  "x-token: ... " -H  "Content-Type: application/json" -d "{  \"name\": \"Hibagon2go\"}"

### ガチャ関連
#### ガチャ実行(POST)
以下のようにHTTPヘッダに以下の形でトークンとガチャ実行回数を入れ渡すと、ガチャを指定回数実行し、ユーザーの所持キャラクターに追加できる。<br>
例) curl -X POST "http://localhost:8080/gacha/draw" -H  "accept: application/json" -H  "x-token: ... " -H  "Content-Type: application/json" -d "{  \"times\": 2}" <br>
この時、以下のjson形式でユーザー情報を返す。<br>
例) {"results" : [{"characterID":"...","name":"violin"},{"characterID":"...","name":"contrabass"}]}

### キャラクター関連
#### ユーザー所持キャラクター一覧取得(GET)
以下のようにHTTPヘッダに以下の形でトークンを入れ渡すと、ユーザーの所持キャラクター一覧を取得できる。<br>
例) curl -X GET "http://localhost:8080/character/list" -H  "accept: application/json" -H  "x-token: ... " <br>
この時、以下のjson形式でユーザー情報を返す。<br>
例) {"characters" : [{"userCharacterID":"...","characterID":"...","name":"violin"},{"userCharacterID":"...","characterID":"...","name":"contrabass"}]}

## 展望
- Dockerを用いたSwaggerをうまく起動できていない
- ガチャを実行する際のアルゴリズム
- ガチャ実行時に引いたキャラクター(今回は楽器)の強さ(今回は楽器の価値)を付与していて、それを使ってさらなる機能追加
- トークンの作成方法の改善
- 今回の実装方法だとuserCharacterIDとcharacterIDのどちらかは不要

