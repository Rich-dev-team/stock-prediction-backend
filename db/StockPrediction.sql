CREATE TABLE "company" (
  "id" uuid PRIMARY KEY NOT NULL,
  "company_name" varchar NOT NULL,
  "stock_symbol" varchar NOT NULL,
  "created_at" date DEFAULT (now())
);
CREATE TABLE "stock_price" (
  "id" BIGSERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "price" int NOT NULL,
  "cur_date" date NOT NULL
);
CREATE TABLE "stock_policy" (
  "id" BIGSERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" date NOT NULL
);
CREATE TABLE "composite_income_sheet" (
  "id" BIGSERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" date NOT NULL
);
CREATE TABLE "balance_sheet" (
  "id" BIGSERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" date NOT NULL
);
CREATE TABLE "month_income_sheet" (
  "id" BIGSERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" date NOT NULL
);
CREATE TABLE "cash_flow" (
  "id" BIGSERIAL PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "cur_date" date NOT NULL
);
CREATE INDEX ON "company" ("company_name");
CREATE INDEX ON "company" ("stock_symbol");
CREATE INDEX ON "stock_price" ("company_id");
CREATE INDEX ON "stock_price" ("cur_date");
CREATE INDEX ON "stock_price" ("company_id", "cur_date");
CREATE INDEX ON "stock_policy" ("company_id");
CREATE INDEX ON "stock_policy" ("cur_date");
CREATE INDEX ON "stock_policy" ("company_id", "cur_date");
CREATE INDEX ON "composite_income_sheet" ("company_id");
CREATE INDEX ON "composite_income_sheet" ("cur_date");
CREATE INDEX ON "composite_income_sheet" ("company_id", "cur_date");
CREATE INDEX ON "balance_sheet" ("company_id");
CREATE INDEX ON "balance_sheet" ("cur_date");
CREATE INDEX ON "balance_sheet" ("company_id", "cur_date");
CREATE INDEX ON "month_income_sheet" ("company_id");
CREATE INDEX ON "month_income_sheet" ("cur_date");
CREATE INDEX ON "month_income_sheet" ("company_id", "cur_date");
CREATE INDEX ON "cash_flow" ("company_id");
CREATE INDEX ON "cash_flow" ("cur_date");
CREATE INDEX ON "cash_flow" ("company_id", "cur_date");
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