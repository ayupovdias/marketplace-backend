CREATE TABLE reviews(
                        id SERIAL PRIMARY KEY,
                        reviewer_id INT NOT NULL,
                        reviewed_user_id INT NOT NULL,
                        rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
                        created_at TIMESTAMP DEFAULT now(),
                        FOREIGN KEY (reviewer_id) REFERENCES users(id) ON DELETE CASCADE,
                        FOREIGN KEY (reviewed_user_id) REFERENCES users(id) ON DELETE CASCADE,
                        CHECK (reviewer_id <> reviewed_user_id)
);

