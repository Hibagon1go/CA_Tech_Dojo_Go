# MySQLサーバーが起動するまで待機する
until mysqladmin ping -h mysql -P 3306 --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo "app is starting...!"
if [ ! -e "go.mod" ]; then
  exec go mod init # go.modファイルが無ければgo.modファイルを作成し、初期化
fi

exec go run main.go