CREATE TABLE "group" (
	"id" serial8 PRIMARY KEY,
	"name" text NOT NULL,
	"parent_group" int8 REFERENCES "group" ON DELETE CASCADE
);

CREATE TABLE "user" (
	"id" serial8 PRIMARY KEY,
	"first_name" text NOT NULL,
	"last_name" text NOT NULL,
	"birth_year" int2 NOT NULL,
	"group" int8 NOT NULL REFERENCES "group" ON DELETE CASCADE
);
