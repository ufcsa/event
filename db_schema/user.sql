DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT 'full name',
  `created` datetime not null default CURRENT_TIMESTAMP COMMENT 'create time',
  `updated` datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  PRIMARY KEY (`id`),
  UNIQUE KEY user_key (`id`, `name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;