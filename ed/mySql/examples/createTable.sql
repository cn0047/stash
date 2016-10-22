create table moderatorComment (
    id int unsigned not null auto_increment key,
    moderatorId int unsigned not null default 0,
    forUserId int unsigned not null default 0,
    message text not null default '',
    createdAt timestamp not null default current_timestamp
    tid tinyint unsigned not null default 0,
    name varchar(25) not null default '' unique,
);
