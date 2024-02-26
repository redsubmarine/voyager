CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL DEFAULT 'regular',
  "created_at" timestamp NOT NULL DEFAULT (now())
);
