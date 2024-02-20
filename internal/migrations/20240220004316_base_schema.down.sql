-- reverse: create index "users_email" to table: "users"
DROP INDEX `users_email`;
-- reverse: create index "users_username" to table: "users"
DROP INDEX `users_username`;
-- reverse: create "users" table
DROP TABLE `users`;
-- reverse: create index "posts_thread_id_post_number" to table: "posts"
DROP INDEX `posts_thread_id_post_number`;
-- reverse: create index "posts_thread_id_id" to table: "posts"
DROP INDEX `posts_thread_id_id`;
-- reverse: create "posts" table
DROP TABLE `posts`;
-- reverse: create index "threads_board_id_id" to table: "threads"
DROP INDEX `threads_board_id_id`;
-- reverse: create "threads" table
DROP TABLE `threads`;
-- reverse: create index "boards_category_id_id" to table: "boards"
DROP INDEX `boards_category_id_id`;
-- reverse: create "boards" table
DROP TABLE `boards`;
-- reverse: create "categories" table
DROP TABLE `categories`;
