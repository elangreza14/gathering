BEGIN;

DROP INDEX IF EXISTS "invitations_id_index";

DROP TABLE IF EXISTS "invitations" CASCADE;

COMMIT;