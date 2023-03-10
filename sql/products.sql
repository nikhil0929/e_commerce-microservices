CREATE TABLE "products" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"name" text,"description" text,"price" decimal,"image" text,"inventory" bigint,"collection" text,PRIMARY KEY ("id"))

SELECT * from products