CREATE TABLE favorites (
                           id SERIAL PRIMARY KEY,
                           user_id INTEGER NOT NULL,
                           advertisement_id INTEGER NOT NULL
);