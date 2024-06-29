CREATE TABLE IF NOT EXISTS cats (
                                    id SERIAL PRIMARY KEY,
                                    name VARCHAR(255) NOT NULL,
    years_of_experience INT NOT NULL,
    breed VARCHAR(255) NOT NULL,
    salary NUMERIC(10, 2) NOT NULL
    );

CREATE TABLE IF NOT EXISTS missions (
                                        id SERIAL PRIMARY KEY,
                                        cat_id INT,
                                        complete BOOLEAN NOT NULL DEFAULT FALSE,
                                        FOREIGN KEY (cat_id) REFERENCES cats (id)
    );

CREATE TABLE IF NOT EXISTS targets (
                                       id SERIAL PRIMARY KEY,
                                       mission_id INT,
                                       name VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    notes TEXT,
    complete BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (mission_id) REFERENCES missions (id)
    );
