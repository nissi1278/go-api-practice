create table if not exists articles (
    article_id integer unsigned AUTO_INCREMENT primary key,
    title varchar(100) not null,
    contents text not null,
    username varchar(100) not null,
    nice integer not null,
    created_at datetime
);

create table if not exists comments (
    comment_id integer unsigned AUTO_INCREMENT primary key,
    article_id integer unsigned not null,
    message text not null,
    created_at datetime,
    foreign key(article_id) references articles(article_id)
);

INSERT INTO articles (title, contents, username, nice, created_at) VALUES
    ('firstPost', 'This is my first blog', 'nissi', 2, now());

INSERT INTO articles (title, contents, username, nice) VALUES
    ('second Post', 'This is my second blog', 'nissi', 4);

INSERT INTO comments (article_id, message, created_at) VALUES
    (1,'first comment', now());

INSERT INTO comments (article_id, message) VALUES
    (1, 'second comment');