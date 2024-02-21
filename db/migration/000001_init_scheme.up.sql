CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL DEFAULT 'regular',
  "created_at" timestamp NOT NULL DEFAULT (now())
);
