CREATE TABLE properties
(
    "id"                    serial PRIMARY KEY,
    "name"                  varchar(255) NOT NULL,
    "zip"                   int          NOT NULL,
    "address"               varchar(255) NOT NULL,
    "current_price_per_day" numeric      NOT NULL,
    "created_at"            timestamp    NOT NULL,
    "updated_at"            timestamp    NOT NULL,
    "deleted_at"            timestamp DEFAULT NULL
);

CREATE TABLE bookings
(
    "id"            serial PRIMARY KEY,
    "user_id"       int       NOT NULL,
    "property_id"   int       NOT NULL,
    "start_date"    timestamp NOT NULL,
    "end_date"      timestamp NOT NULL,
    "price_per_day" numeric   NOT NULL,
    "paid"          boolean   NOT NULL,
    "created_at"    timestamp,
    "updated_at"    timestamp,
    "deleted_at"    timestamp DEFAULT NULL
);

INSERT INTO properties(id, name, zip, address, current_price_per_day, created_at, updated_at)
VALUES (1, 'Medos Hotel Budapest', 1061, 'Jókai tér 9', 18204, current_timestamp, current_timestamp);
INSERT INTO properties(id, name, zip, address, current_price_per_day, created_at, updated_at)
VALUES (2, 'Radisson Blu Beke Hotel', 1067, 'Teréz krt. 43', 37546, current_timestamp, current_timestamp);
INSERT INTO properties(id, name, zip, address, current_price_per_day, created_at, updated_at)
VALUES (3, 'Up Hotel Budapest', 1067, 'Csengery u. 31', 26925, current_timestamp, current_timestamp);
INSERT INTO properties(id, name, zip, address, current_price_per_day, created_at, updated_at)
VALUES (4, 'Prestige Hotel Budapest', 1051, 'Vigyázó Ferenc u. 5', 43499, current_timestamp, current_timestamp);
INSERT INTO properties(id, name, zip, address, current_price_per_day, created_at, updated_at)
VALUES (5, 'B&B Hotel Budapest City', 1094, 'Angyal u. 1-3', 17295, current_timestamp, current_timestamp);

INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (1, 1, 1, '2023-02-01', '2023-02-07', 17465, true, current_timestamp, current_timestamp);
INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (2, 1, 2, '2023-03-10', '2023-03-19', 39220, true, current_timestamp, current_timestamp);
INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (3, 1, 3, '2023-06-11', '2023-06-12', 24566, true, current_timestamp, current_timestamp);
INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (4, 1, 4, '2023-08-08', '2023-09-02', 43499, true, current_timestamp, current_timestamp);
INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (5, 1, 5, '2023-10-03', '2023-10-06', 18203, false, current_timestamp, current_timestamp);
INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (6, 2, 1, '2024-01-01', '2024-01-07', 18204, true, current_timestamp, current_timestamp);
INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (7, 3, 2, '2024-01-10', '2024-01-19', 37546, true, current_timestamp, current_timestamp);
INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (8, 4, 3, '2024-01-11', '2024-01-12', 26925, true, current_timestamp, current_timestamp);
INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (9, 5, 4, '2024-01-08', '2024-02-02', 43499, true, current_timestamp, current_timestamp);
INSERT INTO bookings(id, user_id, property_id, start_date, end_date, price_per_day, paid, created_at, updated_at)
VALUES (10, 6, 5, '2024-01-03', '2024-01-06', 17295, true, current_timestamp, current_timestamp);