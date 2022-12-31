#バージョンは最新を使用する。
FROM mysql:latest

#操作ログとエラーログの保管ファイルを作成//
RUN touch /var/log/mysql/mysql.log
RUN touch /var/log/mysql/mysql_error.log

#設定ファイルをコピーする。
COPY ./settings/my.cnf/etc/mysql/conf.d