DROP SCHEMA IF EXISTS `attendance`;
CREATE SCHEMA IF NOT EXISTS `attendance` DEFAULT CHARACTER SET utf8mb4;
USE `attendance`;

SET CHARSET utf8mb4;

CREATE TABLE if NOT EXISTS `attendance` (
    `id` VARCHAR(64) NOT NULL COMMENT 'ユーザーID',
    `time`  DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '登録時刻',
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `user` (
    `id` VARCHAR(64) NOT NULL COMMENT 'ユーザーID',
    `name` VARCHAR(128) NOT NULL COMMENT 'ユーザーネーム',
    `affiliation` VARCHAR(64) NOT NULL COMMENT '所属',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成時',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新時',
    PRIMARY KEY (`id`)
);

INSERT INTO `user` (`id`,`name`,`affiliation`) VALUES ("1","test1","test");
INSERT INTO `user` (`id`,`name`,`affiliation`) VALUES ("2","test2","test");
