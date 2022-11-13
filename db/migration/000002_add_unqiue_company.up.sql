ALTER TABLE "company" ADD CONSTRAINT "company_name_unique_key" UNIQUE ("company_name");

ALTER TABLE "company" ADD CONSTRAINT "stock_symobl_unique_key" UNIQUE ("stock_symbol");
