CREATE TABLE Bank (
                      id SERIAL PRIMARY KEY,
                      client_id INT NOT NULL,
                      first_name VARCHAR(255),
                      last_name VARCHAR(255),
                      account_balance DECIMAL(20, 2) NOT NULL CHECK (account_balance >= 0),
                      CONSTRAINT unique_bank_name UNIQUE (first_name, last_name)
);

DO $$
DECLARE
person_cursor CURSOR FOR SELECT id, first_name, last_name FROM Person;
person_record RECORD;
    random_balance DECIMAL;
BEGIN
FOR person_record IN person_cursor LOOP
        random_balance := 10000 + ((random() * 100)::INTEGER * 100)::DECIMAL;

INSERT INTO Bank(client_id, first_name, last_name, account_balance)
VALUES (person_record.id, person_record.first_name, person_record.last_name, random_balance);
END LOOP;
END $$;
