-- +goose Up         
-- +goose StatementBegin
CREATE TABLE `books` (
  `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
  `author_id` bigint(20),
  `publisher_id` bigint(20),
  `title` varchar(128) NOT NULL,
  `language` varchar(20),
  `image` varchar(255) NOT NULL,
  `description` text,
  `release_date` datetime NOT NULL,
  `rating_average` float DEFAULT 0,
  `total_page` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` datetime,
  `deleted_at` datetime
);

-- +goose StatementEnd
-- +goose StatementBegin

CREATE TABLE `authors` (
  `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `birth_year` int(11)
);

-- +goose StatementEnd
-- +goose StatementBegin

CREATE TABLE `categories` (
  `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(128) NOT NULL
);
-- +goose StatementEnd
-- +goose StatementBegin

CREATE TABLE `book_categories` (
  `id` bigint(20),
  `book_id` bigint(20),
  `category_id` bigint(20)
);

-- +goose StatementEnd
-- +goose StatementBegin

CREATE TABLE `publishers` (
  `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
  `logo_path` varchar(128) NOT NULL,
  `name` varchar(50) NOT NULL
);

-- +goose StatementEnd
-- +goose StatementBegin

CREATE TABLE `users` (
  `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(128) NOT NULL,
  `email` varchar(128) NOT NULL,
  `password` varchar(255) NOT NULL,
  `age` int(11) NOT NULL DEFAULT 0,
  `avatar` varchar(255),
  `role` varchar(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` datetime,
  `deleted_at` datetime
);
-- +goose StatementEnd
-- +goose StatementBegin

CREATE TABLE `view_histories` (
  `id` bigint(20),
  `user_id` bigint(20),
  `book_id` bigint(20),
  `last_view_page` int(11),
  `last_view_at` datetime
);
-- +goose StatementEnd
-- +goose StatementBegin

CREATE TABLE `book_ratings` (
  `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `book_id` bigint(20) NOT NULL,
  `rating` tinyint(4) NOT NULL COMMENT 'value from 0 -> 5'
);
-- +goose StatementEnd
-- +goose StatementBegin

CREATE TABLE `comments` (
  `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
  `parent_id` bigint(20) COMMENT 'parrent_id=1: comment gốc(cuốn sách), parent_i>1!: comment trả lời cmt khác(comment con)',
  `user_id` bigint(20) NOT NULL,
  `book_id` bigint(20) NOT NULL,
  `content` varchar(255),
  `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` datetime,
  `deleted_at` datetime
);
-- +goose StatementEnd
-- +goose StatementBegin

insert  into `books`(`id`,`author_id`,`publisher_id`,`title`,`language`,`image`,`description`,`release_date`,`rating_average`,`total_page`,`created_at`,`updated_at`,`deleted_at`) values (1,NULL,NULL,'Đắc Nhân Tâm','En','/images/bookImages13.jpeg','Doi xu voi nhau','2021-11-03 07:00:00',4,100,'2021-11-03 08:30:38','2021-11-03 09:01:10',NULL) 


-- +goose StatementEnd
-- +goose StatementBegin
insert  into `authors`(`id`,`name`,`birth_year`) values (1,'Heeca',1966), (2,'GongCha',1988);

-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `books` ADD FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `books` ADD FOREIGN KEY (`publisher_id`) REFERENCES `publishers` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `book_categories` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `book_categories` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `view_histories` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `view_histories` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `book_ratings` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `book_ratings` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `comments` ADD FOREIGN KEY (`parent_id`) REFERENCES `comments` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `comments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin

ALTER TABLE `comments` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);
-- +goose StatementEnd