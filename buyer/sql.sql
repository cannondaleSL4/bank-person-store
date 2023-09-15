CREATE TABLE Person (
                        id SERIAL PRIMARY KEY,
                        first_name VARCHAR(255) NOT NULL,
                        last_name VARCHAR(255) NOT NULL,
                        salary DECIMAL(20, 2) NOT NULL,
                        CONSTRAINT unique_person_name UNIQUE (first_name, last_name)
);

DO $$
DECLARE
first_names TEXT[] := ARRAY['Liam', 'Noah', 'Oliver', 'Ethan', 'Aiden', 'Lucas', 'Jackson', 'Elijah', 'Benjamin', 'James',
'Logan', 'Mason', 'Alexander', 'Michael', 'Daniel', 'Jacob', 'William', 'Matthew', 'Henry', 'Joseph',
'Samuel', 'David', 'John', 'Sebastian', 'Owen', 'Jack', 'Theodore', 'Joshua', 'Carter', 'Dylan',
'Luke', 'Gabriel', 'Anthony', 'Isaac', 'Grayson', 'Jayden', 'Caleb', 'Christopher', 'Andrew', 'Ryan',
'Nathan', 'Christian', 'Hunter', 'Brayden', 'Jaxon', 'Nicholas', 'Lincoln', 'Jonathan', 'Adam', 'Brandon',
'Jordan', 'Nolan', 'Josiah', 'Tyler', 'Ezra', 'Aaron', 'Charles', 'Robert', 'Angel', 'Connor',
'Adrian', 'Julian', 'Austin', 'Eli', 'Kevin', 'Ian', 'Cameron', 'Thomas', 'Landon', 'Zachary',
'Jose', 'Dominic', 'Asher', 'Mateo', 'Jason', 'Santiago', 'Easton', 'Leo', 'Jeremiah', 'Xavier',
'Ryker', 'Miguel', 'Kayden', 'Jace', 'Ayden', 'Hudson', 'Bryson', 'Carlos', 'Maxwell', 'Damian',
'Juan', 'Diego', 'Alex', 'Wesley', 'Jesus', 'Chase', 'Blake', 'Micah', 'Brody', 'Ezekiel'
];
    last_names TEXT[] := ARRAY['Smith', 'Johnson', 'Williams', 'Brown', 'Jones', 'Garcia', 'Miller', 'Davis', 'Rodriguez', 'Martinez',
'Perez', 'Wilson', 'Anderson', 'Taylor', 'Thomas', 'Hernandez', 'Moore', 'Martin', 'Jackson', 'Thompson',
'White', 'Lopez', 'Lee', 'Gonzalez', 'Harris', 'Clark', 'Lewis', 'Robinson', 'Walker', 'Hall',
'Allen', 'Young', 'Hernandez', 'King', 'Wright', 'Lopez', 'Hill', 'Scott', 'Green', 'Adams',
'Baker', 'Gonzalez', 'Nelson', 'Carter', 'Mitchell', 'Perez', 'Roberts', 'Turner', 'Phillips', 'Campbell',
'Parker', 'Evans', 'Edwards', 'Collins', 'Stewart', 'Sanchez', 'Morris', 'Rogers', 'Reed', 'Cook',
'Morgan', 'Bell', 'Murphy', 'Bailey', 'Rivera', 'Cooper', 'Richardson', 'Cox', 'Howard', 'Ward',
'Torres', 'Peterson', 'Gray', 'Ramirez', 'James', 'Watson', 'Brooks', 'Kelly', 'Sanders', 'Price',
'Bennett', 'Wood', 'Barnes', 'Ross', 'Henderson', 'Coleman', 'Jenkins', 'Perry', 'Powell', 'Long',
'Patterson', 'Hughes', 'Flores', 'Washington', 'Butler', 'Simmons', 'Foster', 'Gonzales', 'Bryant', 'Alexander'
];
    random_first_name TEXT;
    random_last_name TEXT;
    random_salary DECIMAL;
BEGIN
FOR i IN 1..1000 LOOP
        random_first_name := first_names[1 + (random() * (array_length(first_names, 1) - 1))::INTEGER];
	    random_last_name := last_names[1 + (random() * (array_length(last_names, 1) - 1))::INTEGER];
        random_salary := 10000 + ((random() * 100)::INTEGER * 100)::DECIMAL;

INSERT INTO Person(first_name, last_name, salary)
VALUES (random_first_name, random_last_name, random_salary)
    ON CONFLICT (first_name, last_name) DO NOTHING;
END LOOP;
END $$;