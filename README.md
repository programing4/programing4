# programing4


# database
- database:mariadb
- driver:mysql
- tabatase name:BBS

# migration
- 変更するたびに
<pre>
goose down
goose up
</pre>

# 環境変数の設定
~.zshrc に追加

    export MYSQL_USERNAME="your_username"
    export MYSQL_PASSWORD="your_password"
