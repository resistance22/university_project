CREATE TABLE IF NOT EXISTS "product" (
  "id" serial PRIMARY KEY,
  "created_by" serial NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar
);

CREATE TABLE IF NOT EXISTS "product_image" (
  "product_id" serial,
  "file_id" serial,
  PRIMARY KEY ("product_id", "file_id")
);

CREATE TABLE IF NOT EXISTS "file" (
  "id" serial PRIMARY KEY,
  "uploaded_by" serial NOT NULL,
  "location" varchar NOT NULL,
  "file_name" varchar NOT NULL,
  "mime_type" varchar(127) NOT NULL,
  "ext" varchar(10) NOT NULL
);

ALTER TABLE "product" ADD FOREIGN KEY ("created_by") REFERENCES "app_user" ("id");

ALTER TABLE "product_image" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_image" ADD FOREIGN KEY ("file_id") REFERENCES "file" ("id");

ALTER TABLE "file" ADD FOREIGN KEY ("uploaded_by") REFERENCES "app_user" ("id");