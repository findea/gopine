CREATE DATABASE goweb;

USE goweb;

CREATE TABLE IF NOT EXISTS `user`
(
    `user_id`       int(10) unsigned NOT NULL AUTO_INCREMENT,            -- 用户ID
    `nickname`      varchar(255)     NOT NULL DEFAULT '',                -- 昵称
    `email`         VARCHAR(255)              DEFAULT NULL,              -- 邮箱
    `password`      VARCHAR(255)              DEFAULT NULL,              -- 密码
    `last_login_ip` varchar(64)      NULL,                               -- 最后登录的IP
    `created_time`  timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_time`  timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 最后修改时间
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `email` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  AUTO_INCREMENT = 100000;