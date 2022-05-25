CREATE TABLE IF NOT EXISTS `users` (
id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID' PRIMARY KEY,
mail varchar(255) NOT NULL COMMENT 'メールアドレス',
nickname varchar(20) NOT NULL COMMENT 'ユーザー名',
created_at datetime NOT NULL COMMENT '登録日時',
updated_at datetime NOT NULL COMMENT '更新日時',
deleted_at datetime COMMENT '削除日時'
) ENGINE= InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `m_todo_statusa` (
id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ステータスID' PRIMARY KEY,
name varchar(20) NOT NULL COMMENT 'ステータス名',
`rank` int(11) NOT NULL COMMENT 'ランク',
created_at datetime NOT NULL COMMENT '登録日時',
updated_at datetime COMMENT '更新日時',
INDEX `rank_index` (`rank`)
) ENGINE= InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `todos` (
id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
user_id int(11) unsigned COMMENT 'ユーザーID',
title varchar(20) NOT NULL COMMENT 'タイトル',
description text COMMENT '詳細',
status tinyint(4) NOT NULL COMMENT 'ステータス',
start_date datetime COMMENT '開始日',
due_date datetime COMMENT '期限日',
created_at datetime NOT NULL COMMENT '登録日時',
updated_at datetime NOT NULL COMMENT '更新日時',
INDEX user_id_index (user_id),
INDEX status_index (status),
INDEX start_date_index (start_date),
INDEX due_date_index (due_date),
FOREIGN KEY (user_id)
 REFERENCES users(id)
) ENGINE= InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `tags` (
id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'タグID' PRIMARY KEY,
user_id int(11) unsigned NOT NULL COMMENT 'ユーザーID',
name varchar(20) NOT NULL COMMENT 'タグ名',
created_at timestamp NOT NULL COMMENT '登録日時',
updated_at datetime COMMENT '更新日時',
INDEX user_id_index (user_id),
FOREIGN KEY (user_id)
 REFERENCES users(id)
) ENGINE= InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `todo_tags` (
id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
todo_id int(11) unsigned NOT NULL COMMENT 'TODO ID',
tag_id int(11) unsigned NOT NULL COMMENT 'タグID',
INDEX todo_id_index (todo_id),
INDEX tag_id_index (tag_id),
FOREIGN KEY (todo_id)
 REFERENCES todos(id),
FOREIGN KEY (tag_id)
 REFERENCES tags(id)
) ENGINE= InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `m_log_type` (
id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ログタイプID' PRIMARY KEY,
name varchar(20) NOT NULL COMMENT 'ログタイプ名',
description varchar(255) COMMENT '説明',
created_at datetime NOT NULL COMMENT '登録日時',
updated_at datetime COMMENT '更新日時'
) ENGINE= InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `logs` (
id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ログID' PRIMARY KEY,
todo_id int(11) unsigned NOT NULL COMMENT 'TODOID',
type_id int(11) unsigned NOT NULL COMMENT 'タイプID',
created_at datetime NOT NULL COMMENT '登録日時',
FOREIGN KEY (todo_id)
 REFERENCES todos(id),
FOREIGN KEY (type_id)
 REFERENCES m_log_type(id)
) ENGINE= InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4;
