BEGIN;

CREATE TABLE IF NOT EXISTS "gatherings" (
  "id" BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,

  "creator" VARCHAR(50) NOT NULL,
  "type" VARCHAR(50) NOT NULL,
  "schedule_at" TIMESTAMPTZ NOT NULL,
  "name" VARCHAR(50) NOT NULL,
  "location" VARCHAR(50) NOT NULL,

  UNIQUE ("id")
);

-- index
CREATE INDEX "gatherings_id_index" ON "gatherings" ("id");

COMMIT;