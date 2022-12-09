BEGIN;

ALTER TABLE registrations ADD COLUMN "body" TEXT;

CREATE INDEX kind_idx ON registrations(kind);

COMMIT;