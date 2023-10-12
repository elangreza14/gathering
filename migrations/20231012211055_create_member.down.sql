BEGIN;

DROP INDEX IF EXISTS "members_id_index";

DROP INDEX IF EXISTS "members_email_index";

DROP TABLE IF EXISTS "members" CASCADE;

COMMIT;