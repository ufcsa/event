DROP TABLE IF EXISTS `result`;

CREATE TABLE `result` (
  `id` int(5) unsigned NOT NULL AUTO_INCREMENT,
  `prize_id` int(5) unsigned NOT NULL DEFAULT '0' COMMENT 'Prize id',
  `prize_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'prize name',
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'user uuid',
  `left_num` int(5) NOT NULL DEFAULT '0' COMMENT 'number of prize left',
  `verifycode` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '4 digit random number as verification code',
  `client_ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'user\'s ip',
  `created` datetime not null default CURRENT_TIMESTAMP COMMENT 'create time',
  `updated` datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`prize_id`) REFERENCES prize(`id`),
  FOREIGN KEY (`user_id`) REFERENCES user(`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;