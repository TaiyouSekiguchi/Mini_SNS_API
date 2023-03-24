CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `name` varchar(64) DEFAULT '' NOT NULL,
  PRIMARY KEY (`id`)
);
-- user1 user2
CREATE TABLE `friend_link` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user1_id` int(11) NOT NULL,
  `user2_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
);
-- user1 user2 block
CREATE TABLE `block_list` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user1_id` int(11) NOT NULL,
  `user2_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
);

-- seed data
INSERT INTO `users`
  (`user_id`, `name`)
VALUE
  (1, 'test1'),
  (2, 'test2'),
  (3, 'test3'),
  (4, 'test4'),
  (5, 'test5'),
  (6, 'test6'),
  (7, 'test7'),
  (8, 'test8'),
  (9, 'test9');


INSERT INTO `friend_link`
  (`user1_id`, `user2_id`)
VALUE
  (1, 2),
  (1, 4),
  (1, 6),
  (2, 3),
  (2, 5),
  (3, 4),
  (4, 6),
  (6, 3),
  (7, 8);

INSERT INTO `block_list`
  (`user1_id`, `user2_id`)
VALUE
  (1, 4),
  (3, 4);
