BEGIN;

ALTER TABLE registrations DROP COLUMN "body";

DROP INDEX kind_idx;

COMMIT;