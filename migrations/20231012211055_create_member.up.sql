BEGIN;

CREATE TABLE IF NOT EXISTS "members" (
  "id" BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "first_name" VARCHAR(50) NOT NULL,
  "last_name" VARCHAR(50) NOT NULL,
  "email" VARCHAR NOT NULL,
  
  UNIQUE ("id", "email")
);

-- index
CREATE INDEX "members_id_index" ON "members" ("id");
CREATE INDEX "members_email_index" ON "members" ("email");

COMMIT;