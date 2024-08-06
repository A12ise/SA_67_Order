BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "employees" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"employee_id"	integer,
	"first_name"	text,
	"last_name"	text,
	"gender_id"	text,
	"username"	text,
	"password"	text,
	"position_id"	text,
	"register_date"	datetime,
	PRIMARY KEY("id" AUTOINCREMENT),
	CONSTRAINT "fk_orders_employee" FOREIGN KEY("employee_id") REFERENCES "orders"("id")
);
CREATE TABLE IF NOT EXISTS "status_orders" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"status_id"	integer,
	"status_name"	text,
	PRIMARY KEY("id" AUTOINCREMENT),
	CONSTRAINT "fk_orders_status_order" FOREIGN KEY("status_id") REFERENCES "orders"("id")
);
CREATE TABLE IF NOT EXISTS "products" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"product_id"	integer,
	"product_name"	text,
	"category_id"	text,
	"employee_id"	text,
	PRIMARY KEY("id" AUTOINCREMENT),
	CONSTRAINT "fk_order_products_product" FOREIGN KEY("product_id") REFERENCES "order_products"("id")
);
CREATE TABLE IF NOT EXISTS "order_products" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"order_id"	integer,
	"product_id"	integer,
	"category_id"	text,
	"quantity"	integer,
	PRIMARY KEY("id" AUTOINCREMENT),
	CONSTRAINT "fk_products_order_product" FOREIGN KEY("product_id") REFERENCES "products"("id"),
	CONSTRAINT "fk_orders_order_product" FOREIGN KEY("order_id") REFERENCES "orders"("id")
);
CREATE TABLE IF NOT EXISTS "orders" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"order_id"	integer,
	"booking_id"	integer,
	"employee_id"	integer,
	"status_id"	integer,
	CONSTRAINT "fk_order_products_order" FOREIGN KEY("order_id") REFERENCES "order_products"("id"),
	CONSTRAINT "fk_bookings_order" FOREIGN KEY("booking_id") REFERENCES "bookings"("id"),
	CONSTRAINT "fk_employees_order" FOREIGN KEY("employee_id") REFERENCES "employees"("id"),
	CONSTRAINT "fk_status_orders_order" FOREIGN KEY("status_id") REFERENCES "status_orders"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "bookings" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"booking_id"	integer,
	"number_of_customers"	integer,
	"date"	datetime,
	"soup_id_1"	text,
	"soup_id_2"	text,
	"package"	text,
	"table_id"	text,
	"member_id"	text,
	"employee_id"	text,
	CONSTRAINT "fk_orders_booking" FOREIGN KEY("booking_id") REFERENCES "orders"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE INDEX IF NOT EXISTS "idx_employees_deleted_at" ON "employees" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_status_orders_deleted_at" ON "status_orders" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_products_deleted_at" ON "products" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_order_products_deleted_at" ON "order_products" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_orders_deleted_at" ON "orders" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_bookings_deleted_at" ON "bookings" (
	"deleted_at"
);
COMMIT;
