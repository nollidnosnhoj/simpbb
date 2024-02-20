-- create "categories" table
CREATE TABLE `categories` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL,
  `slug` text NOT NULL,
  `position` integer NOT NULL,
  `created_at` text NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CHECK (length(name) <= 255)
);
-- create "boards" table
CREATE TABLE `boards` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL,
  `description` text NULL,
  `position` integer NOT NULL,
  `hidden` integer NOT NULL DEFAULT 0,
  `locked` integer NOT NULL DEFAULT 0,
  `redirect` text NULL,
  `created_at` text NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  `category_id` integer NOT NULL,
  `parent_board_id` integer NULL,
  CONSTRAINT `0` FOREIGN KEY (`parent_board_id`) REFERENCES `boards` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CHECK (length(name) <= 255),
  CHECK (length(description) <= 255)
);
-- create index "boards_category_id_id" to table: "boards"
CREATE UNIQUE INDEX `boards_category_id_id` ON `boards` (`category_id`, `id`);
-- create "threads" table
CREATE TABLE `threads` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `subject` text NOT NULL,
  `slug` text NOT NULL,
  `body` text NOT NULL,
  `pinned` integer NOT NULL DEFAULT 0,
  `locked` integer NOT NULL DEFAULT 0,
  `created_at` text NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` text NULL,
  `updated_reason` text NULL,
  `board_id` integer NOT NULL,
  `user_id` integer NOT NULL,
  CONSTRAINT `0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `1` FOREIGN KEY (`board_id`) REFERENCES `boards` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CHECK (length(subject) <= 255)
);
-- create index "threads_board_id_id" to table: "threads"
CREATE UNIQUE INDEX `threads_board_id_id` ON `threads` (`board_id`, `id`);
-- create "posts" table
CREATE TABLE `posts` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `post_number` integer NOT NULL,
  `body` text NOT NULL,
  `created_at` text NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` text NULL,
  `updated_reason` text NULL,
  `thread_id` integer NOT NULL,
  `user_id` integer NOT NULL,
  `reply_to_post_id` integer NULL,
  CONSTRAINT `0` FOREIGN KEY (`reply_to_post_id`) REFERENCES `posts` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `2` FOREIGN KEY (`thread_id`) REFERENCES `threads` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "posts_thread_id_id" to table: "posts"
CREATE UNIQUE INDEX `posts_thread_id_id` ON `posts` (`thread_id`, `id`);
-- create index "posts_thread_id_post_number" to table: "posts"
CREATE INDEX `posts_thread_id_post_number` ON `posts` (`thread_id`, `post_number`);
-- create "users" table
CREATE TABLE `users` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `username` text NOT NULL,
  `email` text NOT NULL,
  `password` text NOT NULL,
  `created_at` text NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CHECK (length(username) <= 50)
);
-- create index "users_username" to table: "users"
CREATE UNIQUE INDEX `users_username` ON `users` (`username`);
-- create index "users_email" to table: "users"
CREATE UNIQUE INDEX `users_email` ON `users` (`email`);
