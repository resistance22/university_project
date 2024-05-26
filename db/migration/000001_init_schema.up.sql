CREATE TABLE "employee" (
  "id" SERIAL PRIMARY KEY,
  "createdAt" timestamp,
  "startedWorking" timestamp,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "father_name" varchar NOT NULL
);

CREATE TABLE "consumable" (
  "id" SERIAL PRIMARY KEY,
  "createdAt" timestamp NOT NULL,
  "title" varchar NOT NULL,
  "uom" varchar NOT NULL,
  "remaining" float NOT NULL
);

CREATE TABLE "consumption_report" (
  "id" SERIAL PRIMARY KEY,
  "date" date NOT NULL,
  "consumable" serial NOT NULL
);

CREATE TABLE "production_report" (
  "id" SERIAL PRIMARY KEY,
  "date" date NOT NULL,
  "amount" float NOT NULL
);

CREATE TABLE "sort" (
  "id" SERIAL PRIMARY KEY,
  "code" varchar NOT NULL,
  "remaining" float NOT NULL
);

CREATE TABLE "sort_report" (
  "id" SERIAL PRIMARY KEY,
  "date" date NOT NULL,
  "amount" float NOT NULL,
  "sort" serial NOT NULL
);

CREATE TABLE "customer" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "email" varchar
);

CREATE TABLE "purchase" (
  "id" SERIAL PRIMARY KEY,
  "price" float NOT NULL
);

CREATE TABLE "consumable_purchase" (
  "purchase_id" SERIAL PRIMARY KEY,
  "consumable" serial NOT NULL,
  "amount" float NOT NULL
);

CREATE TABLE "other_purchse" (
  "title" SERIAL PRIMARY KEY
);

CREATE TABLE "sale" (
  "id" SERIAL PRIMARY KEY,
  "date" date NOT NULL,
  "customer" serial NOT NULL,
  "sort" serial,
  "unit_price" float NOT NULL,
  "amount" float NOT NULL,
  "discount_per_unit" float NOT NULL
);

CREATE TABLE "payment" (
  "id" SERIAL PRIMARY KEY,
  "date" date NOT NULL,
  "amount" float NOT NULL,
  "employee" serial NOT NULL
);

ALTER TABLE "consumption_report" ADD FOREIGN KEY ("consumable") REFERENCES "consumable" ("id");

ALTER TABLE "sort_report" ADD FOREIGN KEY ("sort") REFERENCES "sort" ("id");

ALTER TABLE "consumable_purchase" ADD FOREIGN KEY ("purchase_id") REFERENCES "purchase" ("id");

ALTER TABLE "consumable_purchase" ADD FOREIGN KEY ("consumable") REFERENCES "consumable" ("id");

ALTER TABLE "other_purchse" ADD FOREIGN KEY ("title") REFERENCES "purchase" ("id");

ALTER TABLE "sale" ADD FOREIGN KEY ("customer") REFERENCES "customer" ("id");

ALTER TABLE "sale" ADD FOREIGN KEY ("sort") REFERENCES "sort" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("employee") REFERENCES "employee" ("id");