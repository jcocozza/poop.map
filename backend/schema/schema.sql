CREATE TABLE poop_location (
    uuid TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    latitude REAL,
    longitude REAL,
    first_created TEXT,
    last_modified TEXT,
    seasonal BOOLEAN,
    seasons_mask INTEGER,
    accessible BOOLEAN,
    upvotes INTEGER,
    downvotes INTEGER
);

CREATE TABLE review (
    uuid TEXT PRIMARY KEY NOT NULL,
    poop_location_uuid TEXT NOT NULL,
    rating INTEGER,
    comment TEXT,
    time TEXT,
    upvotes INTEGER,
    downvotes INTEGER,

    FOREIGN KEY (poop_location_uuid) REFERENCES poop_location (uuid)
);
