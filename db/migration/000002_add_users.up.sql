CREATE TABLE "users" (
                         "username" varchar PRIMARY KEY,
                         "password" varchar NOT NULL,
                         "name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "updated_at" timestamp NOT NULL DEFAULT '0001-01-01',
                         "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");