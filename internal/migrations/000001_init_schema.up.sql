create table teacher
(
    id bigserial primary key,
    username varchar(20) not null,
    password_hash varchar(255) not null,
    email varchar(20) not null,
    token_email text,
    is_email_verified boolean not null default false,
    refresh_token text,
    refresh_token_exxpiry bigint not null default 0
);

create table homeworks
(
    id bigserial primary key,
    title varchar(50),
    descript text,
    mark int not null default 0
);
create table lesson
(
    id bigserial primary key,
    title varchar(50),
    filling text,
    file_name varchar(100)
);

create table lesson_homeworks
(
    id bigserial primary key,
    lesson_id bigint references lesson(id) on delete cascade,
    homework_id bigint references homeworks(id) on delete cascade
);

create table lesson_teacher
(
    id bigserial primary key,
    lesson_id bigint not null references lesson(id) on delete cascade,
    teacher_id bigint not null references teacher(id) on delete cascade
);

create table student
(
    id bigserial primary key,
    username varchar(20),
    password_hash varchar(255),
    email varchar(20),
    token_email text,
    is_email_verified boolean not null default false,
    refresh_token text,
    refresh_token_exxpiry bigint not null default 0,
    hearts int default 5
);

create table lesson_user
(
    id bigserial primary key,
    lesson_id bigint not null references lesson(id) on delete cascade,
    user_id bigint not null references student(id) on delete cascade
);

create table homeworks_user
(
    id bigserial primary key,
    homework_id bigint not null references homeworks(id) on delete cascade,
    user_id bigint not null references student(id) on delete cascade
)