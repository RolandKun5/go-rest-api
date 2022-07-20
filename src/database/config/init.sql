create table permission_levels (
  "id" uuid not null unique,
  "permission_level" integer,
  primary key(id)
);

create table users (
  "id" uuid not null unique,
  "user_name" text not null,
  "permission_level_id" uuid not null,
  "first_name" text,
  "last_name" text,
  "email" text not null,
  "city" text,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  primary key(id),
  CONSTRAINT "permission_level_fk" FOREIGN KEY ("permission_level_id") REFERENCES permission_levels("id")
);


insert into "permission_levels"("id","permission_level")
values
    ('74ced41d-a6b2-43cc-a6b8-4c759610de62',0),
    ('ae99e937-a79c-46d8-bfbe-3dabd18b7260', 1);
insert into "users"("id","user_name","permission_level_id","first_name","last_name","email","city")
values
    ('da0b9b2f-1b92-4054-8e20-02e66056445d', 'admin_user','74ced41d-a6b2-43cc-a6b8-4c759610de62','admin','user','admin@user.com','Budapest'),
    ('b6f1d94f-f527-4992-b539-3368771260c6', 'normal_user','ae99e937-a79c-46d8-bfbe-3dabd18b7260','normal','user','normal@user.com','Budapest');