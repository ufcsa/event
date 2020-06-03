DROP TABLE IF EXISTS `prize`;

CREATE TABLE `prize` (
  `id` int(5) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'name of the prize',
  `count` int(5) NOT NULL DEFAULT '-1' COMMENT 'number of this prize',
  `left_num` int(5) NOT NULL DEFAULT '0' COMMENT 'number of prize left',
  `begin` int(11) NOT NULL DEFAULT '0' COMMENT 'start time',
  `end` int(11) NOT NULL DEFAULT '0' COMMENT 'end time',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'status: 0: available, 1: deleted, 2: banned',
  `created` datetime not null default CURRENT_TIMESTAMP COMMENT 'create time',
  `updated` datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;