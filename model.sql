create table users (
  id serial,
  username varchar not null,
  password varchar not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  primary key (id)
);
create unique index users_username_key on users (username);
create index users_created_at_idx on users (created_at desc);

create table forums (
  id serial,
  title varchar not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  primary key (id)
);
create index forums_created_at_idx on forums (created_at desc);

create table topics (
  id serial,
  title varchar not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  primary key (id)
);
create index topics_updated_at_idx on topics (updated_at desc);

create table comments (
  id serial,
  content varchar not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  primary key (id)
);
create index comments_created_at_idx on comments (created_at);

create table forum_topics (
  forum_id int,
  topic_id int,
  primary key (topic_id),
  foreign key (forum_id) references forums (id),
  foreign key (topic_id) references topics (id)
);

create table user_topics (
  user_id int,
  topic_id int,
  primary key (topic_id),
  foreign key (user_id) references users (id),
  foreign key (topic_id) references topics (id)
);

create table user_comments (
  user_id int,
  comment_id int,
  primary key (comment_id),
  foreign key (user_id) references users (id),
  foreign key (comment_id) references comments (id)
);

create table topic_comments (
  topic_id int,
  comment_id int,
  primary key (comment_id),
  foreign key (topic_id) references topics (id),
  foreign key (comment_id) references comments (id)
);

