-- user
CREATE TABLE `user` (
    `id` bigint (20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint (20) NOT NULL,
    `username` varchar (64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar (64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar (64) COLLATE utf8mb4_general_ci,
    `gender` tinyint (4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON
UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
-- community
CREATE TABLE `community` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `community_id` INT(10) UNSIGNED NOT NULL,
    `community_name` VARCHAR(128) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction` VARCHAR(128) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`community_id`),
    UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `post_id` BIGINT(20) NOT NULL COMMENT '帖子id',
    `title` VARCHAR(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
    `content` VARCHAR(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `author_id` BIGINT(20) NOT NULL COMMENT '作者的用户id',
    `community_id` BEGIN(20) NOT NULL COMMENT '所属社区',
    `stastus` TINYINT(4) NOT NULL DPEFAULT '1' COMMENT '帖子状态',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;