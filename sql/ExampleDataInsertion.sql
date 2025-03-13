INSERT INTO cats (name, breed, salary, years_of_experience) VALUES
('Tom', 'Abyssinian', 3000, 5),
('Bob', 'American Bobtail', 2500, 3),
('Jerry', 'American Wirehair', 4000, 7),
('Jack', 'British Longhair', 3500, 4),
('Mike', 'Chartreux', 2800, 2);

INSERT INTO targets (name, country, note, is_complete) VALUES
('Target1', 'Poland', 'Target1', false),
('Target2', 'USA', 'Target2', false),
('Target3', 'Germany', 'Target3', false),
('Target4', 'France', 'Target4', false),
('Target5', 'Italy', 'Target5', false);

INSERT INTO missions (assignee_id, target_ids, note, is_complete) VALUES
(1, ARRAY[1, 2], 'Note1', false),
(2, ARRAY[3], 'Note2', false),
(3, ARRAY[4, 5], 'Note3', false);