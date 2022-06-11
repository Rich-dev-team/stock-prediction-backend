CREATE TABLE "company" (
  "id" uuid PRIMARY KEY NOT NULL,
  "company_name" varchar NOT NULL,
  "stockSymbol" varchar NOT NULL,
  "created_at" datetime DEFAULT (now())
);

CREATE TABLE "stockPrice" (
  "id" SERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" datetime NOT NULL
);

CREATE TABLE "stockPolicy" (
  "id" SERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" datetime NOT NULL
);

CREATE TABLE "compositeIncomeSheet" (
  "id" SERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" datetime NOT NULL
);

CREATE TABLE "balanceSheet" (
  "id" SERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" datetime NOT NULL
);

CREATE TABLE "monthIncomeSheet" (
  "id" SERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" datetime NOT NULL
);

CREATE TABLE "cashFlow" (
  "id" SERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" datetime NOT NULL
);

CREATE INDEX ON "company" ("company_name");

CREATE INDEX ON "company" ("stockSymbol");

CREATE INDEX ON "stockPrice" ("company_id");

CREATE INDEX ON "stockPrice" ("cur_date");

CREATE INDEX ON "stockPrice" ("company_id", "cur_date");

CREATE INDEX ON "stockPolicy" ("company_id");

CREATE INDEX ON "stockPolicy" ("cur_date");

CREATE INDEX ON "stockPolicy" ("company_id", "cur_date");

CREATE INDEX ON "compositeIncomeSheet" ("company_id");

CREATE INDEX ON "compositeIncomeSheet" ("cur_date");

CREATE INDEX ON "compositeIncomeSheet" ("company_id", "cur_date");

CREATE INDEX ON "balanceSheet" ("company_id");

CREATE INDEX ON "balanceSheet" ("cur_date");

CREATE INDEX ON "balanceSheet" ("company_id", "cur_date");

CREATE INDEX ON "monthIncomeSheet" ("company_id");

CREATE INDEX ON "monthIncomeSheet" ("cur_date");

CREATE INDEX ON "monthIncomeSheet" ("company_id", "cur_date");

CREATE INDEX ON "cashFlow" ("company_id");

CREATE INDEX ON "cashFlow" ("cur_date");

CREATE INDEX ON "cashFlow" ("company_id", "cur_date");

ALTER TABLE "stockPrice" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "balanceSheet" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "cashFlow" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "stockPolicy" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "compositeIncomeSheet" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "monthIncomeSheet" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");
