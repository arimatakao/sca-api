CREATE TABLE cats
(
    id serial NOT NULL PRIMARY KEY,
    name character varying,
    breed character varying,
    salary integer,
    years_of_experience integer
);

CREATE TABLE targets
(
    id serial NOT NULL PRIMARY KEY,
    name character varying,
    country character varying,
    note character varying,
    is_complete boolean
);

CREATE TABLE missions
(
    id serial NOT NULL PRIMARY KEY,
    assignee_id integer,
    target_ids integer[],
    note text,
    is_complete boolean,
    CONSTRAINT fk_assignee_id FOREIGN KEY (assignee_id)
        REFERENCES cats (id) ON DELETE CASCADE,
);
