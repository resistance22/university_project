CREATE TABLE "employee" (
  "id" uuid PRIMARY KEY,
  "createdAt" timestamp,
  "startedWorking" timestamp,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "father_name" varchar NOT NULL
);

CREATE TABLE "consumable" (
  "id" uuid PRIMARY KEY,
  "createdAt" timestamp NOT NULL,
  "title" varchar NOT NULL,
  "uom" varchar NOT NULL,
  "remaining" float NOT NULL
);

CREATE TABLE "consumption_report" (
  "id" uuid PRIMARY KEY,
  "date" date NOT NULL,
  "consumable" uuid NOT NULL
);

CREATE TABLE "production_report" (
  "id" uuid PRIMARY KEY,
  "date" date NOT NULL,
  "amount" float NOT NULL
);

CREATE TABLE "sort" (
  "id" uuid PRIMARY KEY,
  "code" varchar NOT NULL,
  "remaining" float NOT NULL
);

CREATE TABLE "sort_report" (
  "id" uuid PRIMARY KEY,
  "date" date NOT NULL,
  "amount" float NOT NULL,
  "sort" uuid NOT NULL
);

CREATE TABLE "customer" (
  "id" uuid PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "email" varchar
);

CREATE TABLE "purchase" (
  "id" uuid PRIMARY KEY,
  "price" float NOT NULL
);

CREATE TABLE "consumable_purchase" (
  "purchase_id" uuid PRIMARY KEY,
  "consumable" uuid NOT NULL,
  "amount" float NOT NULL
);

CREATE TABLE "other_purchase" (
  "title" uuid PRIMARY KEY
);

CREATE TABLE "sale" (
  "id" uuid PRIMARY KEY,
  "date" date NOT NULL,
  "customer" uuid NOT NULL,
  "sort" uuid,
  "unit_price" float NOT NULL,
  "amount" float NOT NULL,
  "discount_per_unit" float NOT NULL
);

CREATE TABLE "payment" (
  "id" uuid PRIMARY KEY,
  "date" date NOT NULL,
  "amount" float NOT NULL,
  "employee" uuid NOT NULL
);

CREATE TABLE "app_user" (
  "id" uuid PRIMARY KEY,
  "created_at" date NOT NULL,
  "first_name" varchar(50),
  "last_name" varchar(50),
  "user_name" varchar(50),
  "password" varchar
);

ALTER TABLE "consumption_report" ADD FOREIGN KEY ("consumable") REFERENCES "consumable" ("id");

ALTER TABLE "sort_report" ADD FOREIGN KEY ("sort") REFERENCES "sort" ("id");

ALTER TABLE "consumable_purchase" ADD FOREIGN KEY ("purchase_id") REFERENCES "purchase" ("id");

ALTER TABLE "consumable_purchase" ADD FOREIGN KEY ("consumable") REFERENCES "consumable" ("id");

ALTER TABLE "other_purchase" ADD FOREIGN KEY ("title") REFERENCES "purchase" ("id");

ALTER TABLE "sale" ADD FOREIGN KEY ("customer") REFERENCES "customer" ("id");

ALTER TABLE "sale" ADD FOREIGN KEY ("sort") REFERENCES "sort" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("employee") REFERENCES "employee" ("id");