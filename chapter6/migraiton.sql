DROP DATABASE IF EXISTS chapter6;
CREATE DATABASE chapter6;

USE chapter6;

DROP TABLE IF EXISTS i_user;
CREATE TABLE i_user (
  user_id BIGINT NOT NULL DEFAULT 0 COMMENT 'ユーザーID',
  os_type INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'OSタイプ',
  name VARCHAR(16) COMMENT 'ユーザー名',
  gamestart_datetime DATETIME NOT NULL DEFAULT "1970-01-01 00:00:00"  COMMENT 'ゲーム開始日時',
  latest_version BIGINT COMMENT '最終バージョン',
  created_at DATETIME NOT NULL DEFAULT "1970-01-01 00:00:00" COMMENT '作成日時',
  updated_at DATETIME NOT NULL DEFAULT "1970-01-01 00:00:00" COMMENT '更新日時',
  PRIMARY KEY (user_id)
) ENGINE = InnoDB COMMENT = 'ユーザー' DEFAULT CHARACTER SET utf8mb4;

INSERT INTO `i_user` (`user_id`, `os_type`, `name`, `gamestart_datetime`, `latest_version`, `created_at`, `updated_at`)
VALUES
	(1, 1, '★キリト★', '2020-03-09 05:54:18', 1, '2020-03-09 05:54:10', '2020-03-09 05:54:14'),
	(2, 2, '†アスナ†', '2020-03-09 05:55:17', 2, '2020-03-09 05:55:21', '2020-03-09 05:55:21');

DROP TABLE IF EXISTS i_user_status;
CREATE TABLE i_user_status (
  user_id BIGINT NOT NULL DEFAULT 0 COMMENT 'ユーザーID',
  level INT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'レベル',
  stumina INT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'スタミナ',
  exp INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '経験値',
  created_at DATETIME NOT NULL DEFAULT "1970-01-01 00:00:00" COMMENT '作成日時',
  updated_at DATETIME NOT NULL DEFAULT "1970-01-01 00:00:00" COMMENT '更新日時',
  PRIMARY KEY (user_id)
) ENGINE = InnoDB COMMENT = 'ユーザー' DEFAULT CHARACTER SET utf8mb4;

INSERT INTO `i_user_status` (`user_id`, `level`, `stumina`, `exp`, `created_at`, `updated_at`)
VALUES
	(1, 99999, 100, 10000, '2020-03-09 05:54:10', '2020-03-09 05:54:14'),
	(2, 30000, 80, 200, '2020-03-09 05:55:21', '2020-03-09 05:55:21');