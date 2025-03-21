-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2025-03-21T19:38:29.043Z

CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "profile_picture" varchar NOT NULL,
  "bio" text NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "post_id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "content_url" varchar NOT NULL,
  "caption" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comments" (
  "comment_id" bigserial PRIMARY KEY,
  "post_id" bigserial NOT NULL,
  "user_id" bigserial NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "likes" (
  "like_id" bigserial PRIMARY KEY,
  "post_id" bigserial NOT NULL,
  "user_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "follows" (
  "follower_id" bigserial NOT NULL,
  "following_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("follower_id", "following_id")
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "likes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id");

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "follows" ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("user_id");

ALTER TABLE "follows" ADD FOREIGN KEY ("following_id") REFERENCES "users" ("user_id");
