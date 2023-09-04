#!/bin/bash


# MySQL接続情報
DB_USER="user"
DB_PASSWORD="password"
DB_NAME="db"

# SQLファイル名
DB_TABLE_NAME=$1
if [ -z "$DB_TABLE_NAME" ]; then
    echo "Table name not found"
    exit 1
fi

# MySQLクエリの実行
mysql -h localhost --port 3306 --protocol tcp -u $DB_USER -p$DB_PASSWORD $DB_NAME -e "

CREATE TABLE $DB_TABLE_NAME (
  id         INT AUTO_INCREMENT NOT NULL,
  user_id    VARCHAR(128) NOT NULL,
  message    VARCHAR(255) NOT NULL,
  PRIMARY KEY (\`id\`)
);

INSERT INTO $DB_TABLE_NAME
  (user_id, message)
VALUES
  ('user1', 'hello world'),
  ('user2', 'nice to meet you'),
  ('user1', 'how are you'),
  ('user3', 'fine thank you');
"
