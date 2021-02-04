# MySQLサーバーが起動するまで待機する
until mysqladmin ping -h mysql -P 3306 --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo "app is starting...!"
#exec go mod init modファイルがない場合のみ実行
exec go run main.go