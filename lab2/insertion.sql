BEGIN;

set schema 'public';

insert into "user"(id, name, email, password) values(1, 'User1', 'email1@gmail.com', sha256('pass1')),
(2, 'User2', 'email2@gmail.com', sha256('pass2')),
(3, 'User3', 'email3@gmail.com', sha256('pass3')),
(4, 'User4', 'email4@gmail.com', sha256('pass5')),
(5, 'User5', 'email5@gmail.com', sha256('pass1')),
(6, 'User6', 'email6@gmail.com', sha256('pass1')),
(7, 'User7', 'email7@gmail.com', sha256('pass1')),
(8, 'User8', 'email8@gmail.com', sha256('pass1')),
(9, 'User9', 'email9@gmail.com', sha256('pass1')),
(10, 'User10', 'email10@gmail.com', sha256('pass10'));

insert into friends(user_id, friend_id, tag) values (1, 2, 'school_friend'),
(1, 3, 'relative'),
(2, 5, 'best_friend'),
(7, 9, 'colleague'),
(3, 4, 'university_friend'),
(6, 10, 'relative'),
(3, 8, 'best_friend');

insert into channel(id, name) values (1, 'Cool channel'),
(2, 'Bad channel'),
(3, 'Great channel'),
(4, 'Perfect channel'),
(5, 'Work channel'),
(6, 'Study channel');

insert into channel_member(channel_id, user_id, nickname, banned) values (1, 1, 'the worst', true), 
(4, 2, 'the best', true), 
(6, 1, 'the best', true),
(1, 5, 'mem1', true);
insert into channel_member(channel_id, user_id) values (1, 2),
(3, 2),
(5, 2),
(6, 3),
(5, 4),
(2, 7),
(2, 9), 
(6, 5),
(6, 8);
insert into public.channel_member(channel_id, user_id, nickname) values (2, 1, 'the best'),
(1, 3, 'mem1'),
(3, 1, 'the best'),
(2, 5, 'the best'),
(4, 7, 'the best'),
(5, 8, 'm1'),
(3, 10, 'number 1');

insert into "role"(id, name, channel_id) values (1, 'admin', 1),
(2, 'admin', 2),
(3, 'admin', 3),
(4, 'admin', 4),
(5, 'admin', 5),
(6, 'admin', 6),
(7, 'bad user', 2),
(8, 'cool user', 1),
(9, 'Great user', 3),
(10, 'developer', 5),
(11, 'teacher', 6),
(12, 'student', 6),
(13, 'tester', 5);

insert into "permission"(id, name, description) values (1, 'ban.member', 'Забанить пользователя на канале'),
(2, 'mute.member', 'Замьютить пользователя');
insert into "permission"(id, name) values (3, 'create.role');
insert into "permission"(id, name, description) values (4, 'share.screen', 'Демонстрировать экран');
insert into "permission"(id, name) values (5, 'use.chat');

insert into role_permission(role_id, permission_id) values (1, 1),
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(2, 1),
(2, 2),
(2, 3),
(3, 1),
(3, 2),
(3, 3),
(4, 1),
(4, 2),
(4, 3),
(5, 1),
(5, 2),
(5, 3),
(6, 1),
(6, 2),
(6, 3),
(7, 5),
(8, 5),
(8, 4),
(9, 5),
(9, 4),
(9, 3),
(10, 5),
(10, 4),
(10, 3),
(11, 5),
(11, 4),
(11, 3),
(11, 2),
(12, 5),
(12, 4),
(13, 4),
(13, 5);

insert into chat(id, name, channel_id) values (1, 'cool chat', 1),
(2, 'ok chat', 1),
(3, 'best chat', 3),
(4, 'worst chat', 2),
(5, 'bad chat', 2),
(6, 'perfect chat', 4),
(7, 'working', 5),
(8, '1st group', 6),
(9, '2nd group', 6),
(10, 'chill chat', 5);
insert into chat(id, name) values(11, 'friends chat'),
(12, 'our chat');


insert into chat_user(chat_id, user_id) values (3, 1),
(9, 1),
(1, 3),
(12, 7),
(12, 9),
(11, 2),
(11, 5),
(2, 1),
(2, 3),
(1, 1),
(4, 5),
(5, 5),
(6, 7),
(7, 2),
(10, 2);

insert into "call"(id, chat_id, duration, date) values (1, 11, '00:25:10', current_timestamp),
(2, 1, '00:55:35', current_timestamp),
(3, 12,'00:15:35', current_timestamp),
(4, 2,'00:15:35', '2023-11-04 15:00:00'),
(5, 11,'01:12:35', '2023-11-05 19:00:00'),
(6, 1,'00:12:35', '2023-10-31 17:00:00');

insert into "message"(id, text, chat_id, author, date) values (1, 'hello', 1, 3, current_timestamp - interval '1 hour'),
(2, 'hi', 1, 1, current_timestamp - interval '30 minute'),
(3, 'how are you', 12, 7, current_timestamp - interval '3 minute'),
(4, 'lets talk', 12, 9, current_timestamp),
(5, 'hello world', 2, 1, current_timestamp),
(6, 'i am alone', 4, 5, current_timestamp);
insert into "message"(id, file, chat_id, author, date) values (7, '\x123456'::bytea, 10, 2, current_timestamp);

insert into channel_member_role(channel_member_user_id, role_id) values (1, 8),
(1, 12),
(3, 1),
(3, 11),
(7, 4),
(2, 10),
(5, 8),
(8, 5),
(10, 9),
(4, 13),
(7, 7),
(2, 3),
(9, 2),
(5, 6),
(8, 11);

END;
