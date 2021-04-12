CREATE TABLE `user` (
	`id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`username` varchar(20) NOT NULL DEFAULT '' COMMENT 'username',
	`password` varchar(100) NOT NULL DEFAULT '' COMMENT 'password',
	`email` varchar(100) NOT NULL DEFAULT '' COMMENT 'email',
	PRIMARY KEY (`id`),
	UNIQUE KEY(`username`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8

INSERT INTO `user` (`id`,`username`,`password`,`email`) VALUES(?,?,?,?)

SELECT * FROM `user` WHERE `username` = ?

UPDATE `user` SET `password` = ? , `email` = ? WHERE `username` = ?

CREATE TABLE `file` (
	`id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`username` varchar(20) NOT NULL DEFAULT '' COMMENT 'username',
	`filename` varchar(100) NOT NULL DEFAULT '' COMMENT 'filename',
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	UNIQUE KEY `per_file` (`username`,`filename`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8

INSERT INTO `file` (`id`,`username`,`filename`,`update_time`) VALUES(0,?,?,?)

UPDATE `file` SET `filename` = ?,`update_time` = ? WHERE `username` = ?

DELETE from `file` WHERE `username` = ? AND `filename` = ?

SELECT * FROM `file` WHERE `username` = ?

CREATE TABLE `kv` (
	`id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`key` varchar(255) NOT NULL DEFAULT '' COMMENT 'key',
	`value` varchar(255) NOT NULL DEFAULT '' COMMENT 'value',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8
