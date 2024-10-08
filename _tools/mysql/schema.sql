-- todo.`user` definition

CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
  `name` varchar(20) NOT NULL COMMENT 'ユーザー名',
  `password` varchar(80) NOT NULL COMMENT 'パスワードハッシュ',
  `role` varchar(80) NOT NULL,
  `created` datetime(6) NOT NULL,
  `modified` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='ユーザー';

-- todo.task definition

CREATE TABLE `task` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'タスクの識別子',
  `user_id` bigint unsigned NOT NULL COMMENT 'タスクを作成したユーザの識別子',
  `title` varchar(128) NOT NULL COMMENT 'タスクのタイトル',
  `status` varchar(20) NOT NULL COMMENT 'タスクの状態',
  `created` datetime(6) NOT NULL COMMENT 'レコード作成日時	',
  `modified` datetime(6) NOT NULL COMMENT 'レコードの修正日時	',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id`
        FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='タスク';