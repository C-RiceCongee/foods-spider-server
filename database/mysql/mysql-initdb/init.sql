-- auto-generated definition
SET GLOBAL time_zone = '+8:00';
CREATE DATABASE IF NOT EXISTS cldquicker;
use cldquicker;
DROP TABLE IF EXISTS `users`;

create table users
(
    id          int auto_increment comment '自增主键'
        primary key comment '主键为自增的id',
    userid      bigint                               not null comment '用户ID',
    username    varchar(50)                          not null comment '用户名',
    avatar      varchar(255)                         null comment '用户头像',
    openid      varchar(50)                          not null comment '微信小程序 openid',
    sessionkey  varchar(255)                         not null comment '小程序登录凭证',
    userpoint   int        default 0                 null comment '用户积分',
    userstatus  tinyint(1) default 0                 null comment '用户状态',
    create_time timestamp  default CURRENT_TIMESTAMP null comment '数据创建时间',
    update_time timestamp  default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '数据更新时间',
    constraint openid
        unique (openid),
    constraint uniq_openid
        unique (openid),
    constraint userid
        unique (userid)
)
    comment '用户表，记录系统中的所有用户信息' charset = utf8mb4;

create index idx_username
    on users (username)
    comment '对用户名加索引';

-- auto-generated definition
DROP TABLE IF EXISTS `files`;
create table files
(
    id          bigint unsigned auto_increment,
    file_id     bigint unsigned                     not null,
    user_id     bigint unsigned                     not null,
    upload_type varchar(10)                         not null comment '上传类型[七牛云，本地..]',
    scene_type  varchar(100)                        not null comment '场景类别[风景，人物...]',
    base_url    varchar(100)                        not null comment 'url(本地路径就是/，非本地就存远程URL)',
    path        varchar(100)                        not null comment '本地路径 /home/xxx 或者是远程的域名',
    name        varchar(100)                        not null comment '文件原始名',
    extension   varchar(50)                         not null comment '扩展名',
    size        int       default 0                 not null comment '长度',
    year        int       default 0                 not null comment '年份',
    month       int       default 0                 not null comment '月份',
    day         int       default 0                 not null comment '日',
    upload_ip   varchar(16)                         not null comment '上传者ip',
    status      tinyint   default 1                 not null comment '状态[-1:删除;0:禁用;1启用]',
    create_time timestamp default CURRENT_TIMESTAMP null comment '数据创建时间',
    update_time timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '数据更新时间',
    primary key (id, file_id)
)
    comment '文件上传管理表';

create index file_id
    on files (file_id);


