CREATE TABLE registrations(
    "brand"           TEXT,
    "capacity"        INT,
    "color"           TEXT NOT NULL,
    "d_first_reg"     DATE,
    "d_reg"           DATE,
    "fuel"            TEXT,
    "kind"            TEXT,
    "make_year"       INT  NOT NULL,
    "model"           TEXT,
    "n_doc"           TEXT NOT NULL,
    "s_doc"           TEXT NOT NULL,
    "n_reg_new"       TEXT NOT NULL,
    "n_seating"       INT,
    "n_standing"      INT,
    "own_weight"      INT,
    "rank_category"   TEXT,
    "total_weight"    INT,
    "vin"             TEXT,
    PRIMARY KEY ("s_doc", "n_doc")
);

CREATE INDEX registrations_n_reg_new_idx ON registrations("n_reg_new");
CREATE INDEX registrations_vin_idx       ON registrations("vin");
