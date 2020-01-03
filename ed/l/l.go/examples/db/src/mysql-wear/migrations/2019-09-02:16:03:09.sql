create table test_mysql (id int key, code int);
insert into test_mysql values (1, 200);

CREATE TABLE `user2020`(
`id` VARCHAR(255) NOT NULL DEFAULT '' KEY,
`password_hash` VARCHAR(255) NOT NULL DEFAULT '',
`token_version` INT NOT NULL DEFAULT 0,
`stripe_customer_id` VARCHAR(255) NOT NULL DEFAULT '',
`status` VARCHAR(255) NOT NULL DEFAULT '',
`created` timestamp NOT NULL,
`updated` timestamp NOT NULL
);
insert into user2020 values ('u1', 'pwd', 1, 'stripe', 'ok', now(), now());

CREATE TABLE `user_email2020`(
`sha1` VARCHAR(255) NOT NULL DEFAULT '' KEY,
`user_id` VARCHAR(255) NOT NULL DEFAULT '',
`email` VARCHAR(255) NOT NULL DEFAULT '',
`primary` tinyint(1) NOT NULL DEFAULT 0,
`created` timestamp NOT NULL,
`updated` timestamp NOT NULL,
`verified_on` timestamp NOT NULL
);

insert into user_email2020 values ('u1sha1', 'u1', 'user1@email.com', 1, now(), now(), default);
