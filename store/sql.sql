CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          price NUMERIC(10,2) NOT NULL CHECK (price >= 0) -- assuming you want two decimal places for the price
);

INSERT INTO products (name, price) VALUES
                                       ('sugar', 1.12),
                                       ('tomatoes', 5.6),
                                       ('cucumbers', 4.9),
                                       ('oranges', 4.95),
                                       ('apples', 5.8),
                                       ('bananas', 4.9),
                                       ('coca-cola', 6.0),
                                       ('wine', 31),
                                       ('bread', 13),
                                       ('cabbage', 2.9),
                                       ('chicken', 19),
                                       ('meat', 45);
