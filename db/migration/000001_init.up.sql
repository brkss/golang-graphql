CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "userinfo" (
  "id" varchar PRIMARY KEY,
  "weight" int NOT NULL,
  "height" int NOT NULL,
  "birth" timestamp NOT NULL,
  "user_id" varchar NOT NULL
);

CREATE TABLE "recipe_categories" (
  "id" varchar PRIMARY KEY,
  "title" varchar NOT NULL,
  "image" varchar NOT NULL,
  "active" bool DEFAULT true
);

CREATE TABLE "recipes" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "image" varchar DEFAULT 'default.png',
  "active" bool DEFAULT true,
  "time" varchar NOT NULL,
  "url" varchar,
  "servings" int,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipe_category_Recipe" (
  "id" varchar PRIMARY KEY,
  "recipe_id" varchar,
  "recipe_category_id" varchar
);

ALTER TABLE "userinfo" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "recipe_category_Recipe" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");

ALTER TABLE "recipe_category_Recipe" ADD FOREIGN KEY ("recipe_category_id") REFERENCES "recipe_categories" ("id");

