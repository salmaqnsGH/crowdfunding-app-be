CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "occupation" varchar,
  "email" varchar,
  "password_hash" varchar,
  "avatar_file_name" varchar,
  "role" varchar,
  "token" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "campaigns" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer,
  "name" varchar,
  "short_description" varchar,
  "description" text,
  "goal_amount" integer,
  "current_amount" integer,
  "perks" text,
  "backer_count" integer,
  "slug" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "campaign_images" (
  "id" SERIAL PRIMARY KEY,
  "campaign_id" integer,
  "file_name" varchar,
  "is_primary" integer,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" SERIAL PRIMARY KEY,
  "campaign_id" integer,
  "user_id" integer,
  "amount" integer,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "campaigns" ("user_id");

CREATE INDEX ON "campaign_images" ("campaign_id");

CREATE INDEX ON "transactions" ("campaign_id");

CREATE INDEX ON "transactions" ("user_id");


ALTER TABLE "campaigns" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "campaign_images" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaigns" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaigns" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

INSERT INTO users( id ,name ,occupation ,email ,password_hash ,avatar_file_name ,role ,token ,created_at ,updated_at ) VALUES('nana','-','-','-','-','-','-',now(),now());
INSERT INTO campaigns( id,user_id ,name ,short_description ,description ,goal_amount ,current_amount ,backer_count ,slug ,created_at ,updated_at ) VALUES(1,1,'nana2','-','-',0,0,0,'-',now(),now());
INSERT INTO campaign_images( id,campaign_id ,file_name ,is_primary ,created_at ,updated_at ) VALUES(3,1,'tiga.png',0,now(),now());

ALTER TABLE campaigns ADD perks text;