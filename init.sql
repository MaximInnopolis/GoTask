CREATE TABLE checks (
                        id SERIAL PRIMARY KEY,
                        store_name VARCHAR(255) NOT NULL,
                        total DECIMAL(10, 2) NOT NULL,
                        payment_method VARCHAR(255) NOT NULL,
                        tax DECIMAL(10, 2) DEFAULT 0.00
);

CREATE TABLE check_hash (
                            id SERIAL PRIMARY KEY,
                            check_id INTEGER NOT NULL REFERENCES checks(id),
                            hash VARCHAR(255) NOT NULL
);

DROP TABLE checks;
DROP TABLE check_hash;