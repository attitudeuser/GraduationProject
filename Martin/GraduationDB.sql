/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2020/1/7 15:55:20                            */
/*==============================================================*/


drop table if exists issue;

drop table if exists message;

drop table if exists subject;

drop table if exists user;

/*==============================================================*/
/* Table: issue                                                 */
/*==============================================================*/
create table issue
(
   Id                   int not null,
   Title                varchar(64),
   Type                 varchar(16),
   "Require"            varchar(128),
   Cap                  int,
   Content              varchar(255),
   Tid                  int not null,
   primary key (Id)
);

/*==============================================================*/
/* Table: message                                               */
/*==============================================================*/
create table message
(
   Id                   int not null,
   Fid                  int not null,
   Tid                  int not null,
   Message              varchar(255),
   Status               bool not null,
   primary key (Id)
);

/*==============================================================*/
/* Table: subject                                               */
/*==============================================================*/
create table subject
(
   Id                   int not null auto_increment,
   Uid                  int not null,
   Iid                  int not null,
   primary key (Id)
);

/*==============================================================*/
/* Table: user                                                  */
/*==============================================================*/
create table user
(
   Id                   int not null auto_increment,
   Username             varchar(32) not null,
   Password             varchar(64) not null,
   Email                varchar(128) not null,
   Phone                varchar(11),
   Type                 varchar(8),
   CreateTime           time,
   Age                  int,
   Sid                  varchar(11),
   College              varchar(16),
   Grade                char(10),
   primary key (Id)
);

alter table issue add constraint FK_Reference_5 foreign key (Tid)
      references user (Id) on delete restrict on update restrict;

alter table message add constraint FK_Reference_1 foreign key (Fid)
      references user (Id) on delete restrict on update restrict;

alter table message add constraint FK_Reference_4 foreign key (Tid)
      references user (Id) on delete restrict on update restrict;

alter table subject add constraint FK_Reference_2 foreign key (Uid)
      references user (Id) on delete restrict on update restrict;

alter table subject add constraint FK_Reference_3 foreign key (Iid)
      references issue (Id) on delete restrict on update restrict;

