create table if not exists {{{table_name}}}
(
    autoid  bigint(20) unsigned NOT NULL AUTO_INCREMENT
        primary key,
    recdate     datetime                   not null,
    clientip   varchar(15)                 not null,
    urlpath    varchar(255)                not null,
    userid     varchar(50)                 not null
);

