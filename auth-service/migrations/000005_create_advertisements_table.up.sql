CREATE TABLE advertisements (
                                id SERIAL PRIMARY KEY,
                                title TEXT NOT NULL,
                                description TEXT NOT NULL,
                                price NUMERIC NOT NULL,
                                city TEXT,
                                image_url TEXT,
                                user_id INTEGER NOT NULL,
                                category_id INTEGER,
                                created_at TIMESTAMP DEFAULT now()
);