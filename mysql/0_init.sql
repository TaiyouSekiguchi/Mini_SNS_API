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
  (0, 'test0'),
  (1, 'test1'),
  (2, 'test2'),
  (3, 'test3'),
  (4, 'test4'),
  (5, 'test5'),
  (6, 'test6'),
  (7, 'test7'),
  (8, 'test8'),
  (9, 'test9'),
  (10, 'test10'),
  (11, 'test11'),
  (12, 'test12'),
  (13, 'test13'),
  (14, 'test14'),
  (15, 'test15'),
  (16, 'test16'),
  (17, 'test17'),
  (18, 'test18'),
  (19, 'test19'),
  (20, 'test20'),
  (21, 'test21'),
  (22, 'test22'),
  (23, 'test23'),
  (24, 'test24'),
  (25, 'test25'),
  (26, 'test26'),
  (27, 'test27'),
  (28, 'test28'),
  (29, 'test29'),
  (30, 'test30'),
  (31, 'test31'),
  (32, 'test32'),
  (33, 'test33'),
  (34, 'test34'),
  (35, 'test35'),
  (36, 'test36'),
  (37, 'test37'),
  (38, 'test38'),
  (39, 'test39'),
  (40, 'test40'),
  (41, 'test41'),
  (42, 'test42'),
  (43, 'test43'),
  (44, 'test44'),
  (45, 'test45'),
  (46, 'test46'),
  (47, 'test47'),
  (48, 'test48'),
  (49, 'test49'),
  (50, 'test50');


INSERT INTO `friend_link`
  (`user1_id`, `user2_id`)
VALUE
  (1, 2),
  (1, 3),
  (1, 4),
  (5, 1),
  (6, 1),
  (6, 2),
  (6, 3),
  (7, 1),
  (8, 1);

INSERT INTO `block_list`
  (`user1_id`, `user2_id`)
VALUE
  (7, 1),
  (8, 7);
