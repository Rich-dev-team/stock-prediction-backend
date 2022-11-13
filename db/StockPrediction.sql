CREATE TABLE "company" (
  "id" bigserial PRIMARY KEY,
  "company_name" varchar UNIQUE NOT NULL,
  "stock_symbol" varchar UNIQUE NOT NULL,
  "created_at" date DEFAULT (now())
);
CREATE TABLE "stock_price" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "price" int NOT NULL,
  "created_at" timestamptz NOT NULL
);
CREATE TABLE "stock_policy" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "date" date NOT NULL
);
CREATE TABLE "composite_income_sheet" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "date" date NOT NULL
);
CREATE TABLE "balance_sheet" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "date" date NOT NULL
);
CREATE TABLE "month_income_sheet" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "date" date NOT NULL
);
CREATE TABLE "cash_flow" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "date" date NOT NULL
);
CREATE INDEX ON "company" ("id");
CREATE INDEX ON "company" ("stock_symbol");
CREATE INDEX ON "stock_price" ("company_id");
CREATE INDEX ON "stock_price" ("created_at");
CREATE INDEX ON "stock_price" ("company_id", "created_at");
CREATE INDEX ON "stock_policy" ("company_id");
CREATE INDEX ON "stock_policy" ("date");
CREATE INDEX ON "stock_policy" ("company_id", "date");
CREATE INDEX ON "composite_income_sheet" ("company_id");
CREATE INDEX ON "composite_income_sheet" ("date");
CREATE INDEX ON "composite_income_sheet" ("company_id", "date");
CREATE INDEX ON "balance_sheet" ("company_id");
CREATE INDEX ON "balance_sheet" ("date");
CREATE INDEX ON "balance_sheet" ("company_id", "date");
CREATE INDEX ON "month_income_sheet" ("company_id");
CREATE INDEX ON "month_income_sheet" ("date");
CREATE INDEX ON "month_income_sheet" ("company_id", "date");
CREATE INDEX ON "cash_flow" ("company_id");
CREATE INDEX ON "cash_flow" ("date");
CREATE INDEX ON "cash_flow" ("company_id", "date");
ALTER TABLE "stock_price"
ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");
ALTER TABLE "balance_sheet"
ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");
ALTER TABLE "cash_flow"
ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");
ALTER TABLE "stock_policy"
ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");
ALTER TABLE "composite_income_sheet"
ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");
ALTER TABLE "month_income_sheet"
ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");