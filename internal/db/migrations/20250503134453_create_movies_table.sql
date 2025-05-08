-- +goose Up
-- +goose StatementBegin

create table if not exists seats (
    id SERIAL PRIMARY KEY,
    seat_number VARCHAR(10) NOT NULL,
    seat_type VARCHAR(50) NOT NULL,
    is_booked BOOLEAN DEFAULT FALSE,
    theatre_id INT NOT NULL,
    FOREIGN KEY (theatre_id) REFERENCES theatres(id)
);

create table if not EXISTS theatres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    total_seats INT NOT NULL
);

create table if not EXISTS shows (
    id SERIAL PRIMARY KEY,
   show_time date NOT NULL,
   movie_id INT NOT NULL,
   theatre_id INT NOT NULL,
   FOREIGN KEY (movie_id) REFERENCES movies(id),
   FOREIGN KEY (theatre_id) REFERENCES theatres(id)
);

CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    price DECIMAL(10, 2) NOT NULL,
    seat_id INT NOT NULL,
    show_id INT NOT NULL,
    customer_id INT NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
    FOREIGN KEY (show_id) REFERENCES shows(id)
    FOREIGN KEY (seat_id) REFERENCES seats(id)
);

Create Table if not EXISTS customers (
id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
phone_number VARCHAR(15) NOT NULL,
)

CREATE TABLE if not EXISTS movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    genre VARCHAR(100),
    duration_minutes INT,
    poster_image VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS shows;
DROP TABLE IF EXISTS theatres;
DROP TABLE IF EXISTS seats;
DROP TABLE IF EXISTS tickets;
-- +goose StatementEnd


