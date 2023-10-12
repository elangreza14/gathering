BEGIN;

DROP INDEX IF EXISTS "gatherings_id_index";

DROP TABLE IF EXISTS "gatherings" CASCADE;

COMMIT;