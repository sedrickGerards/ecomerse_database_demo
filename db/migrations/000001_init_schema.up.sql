
-- CREATE TABLE "orders" (
--   "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
--   "costumer_id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
--   -- "sender" VARCHAR(100) NOT NULL,
--   -- "content" TEXT NOT NULL,
--   "created_at" TIMESTAMP DEFAULT NOW()

-- );



-- CREATE TABLE "order_items" (
--   "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
--   "thread_id" VARCHAR(36) NOT NULL,
--   "sender" VARCHAR(100) NOT NULL,
--   "content" TEXT NOT NULL,
--   "created_at" TIMESTAMP DEFAULT now(),

--   -- setting foreign key costrain
--   FOREIGN KEY ("thread_id") REFERENCES "thread"("id")
-- );

-- Create the 'customer' table
CREATE TABLE "customer" (
    "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
    "customer_name" TEXT NOT NULL,
    "contact" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW()

);

-- Create the 'product' table
CREATE TABLE "product" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "product_name" TEXT NOT NULL,
  "price" DECIMAL(10, 2) NOT NULL,
  "created_at" TIMESTAMP DEFAULT NOW()

);

-- Create the 'orders' table
CREATE TABLE "orders" (
    "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
    "customer_id" VARCHAR(36) NOT NULL,
    "product_id" VARCHAR(36) NOT NULL,
    "order_status" VARCHAR(36) NOT NULL DEFAULT 'PENDING',
    "total_amount" DECIMAL (10, 2) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY ("customer_id") REFERENCES "customer" ("id"),
    FOREIGN KEY ("product_id") REFERENCES "product" ("id")
);



