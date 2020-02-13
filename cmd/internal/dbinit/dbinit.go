package dbinit

import (
	"database/sql"
)

const managersDDl = `CREATE TABLE IF NOT EXISTS managers
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    name    TEXT    NOT NULL,
    login   TEXT    NOT NULL UNIQUE,
	password TEXT NOT NULL,
    salary  INTEGER NOT NULL CHECK ( salary > 0 ),
    plan    INTEGER NOT NULL CHECK ( plan >= 0 ),
    unit text,
    boss_id INTEGER REFERENCES managers
);
`

func Init(db *sql.DB) (err error) {
	_, err = db.Exec(managersDDl)
	if err != nil {
		return err
	}

	// adding managers
	_, err = db.Exec(`INSERT INTO managers
VALUES (1, 'Vasya', 'vasya', 'secret', 100000, 0, NULL , NULL), -- Ctrl + D
       (2, 'Petya', 'petya', 'secret', 90000, 90000, 'boys', 1),
       (3, 'Vanya', 'vanya', 'secret', 80000, 80000, 'boys', 2),
       (4, 'Masha', 'masha', 'secret', 80000, 80000,'girls', 1),
       (5, 'Dasha', 'dasha', 'secret', 60000, 60000,'girls', 4),
       (6, 'Sasha', 'sasha', 'secret', 40000, 40000,'girls', 5) ON CONFLICT DO NOTHING;`)
	if err != nil {
		return err
	}

// create products
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products
(
    id    INTEGER PRIMARY KEY AUTOINCREMENT,
    name  TEXT NOT NULL UNIQUE,
    price INTEGER NOT NULL CHECK ( price > 0 ),
    qty INTEGER NOT NULL CHECK ( qty > 0 )
);`)
	if err != nil {
		return err
	}

	// adding products
_, err = db.Exec(`INSERT INTO products(name, price, qty)
VALUES ('Big Mac', 200, 10),
       ('Chicken Mac', 150, 15),
       ('Cheese Burger', 100, 20),
       ('Tea', 50, 10),
       ('Coffee', 80, 10),
       ('Cola', 100, 20) ON CONFLICT DO NOTHING;`)
if err != nil {
	return err
}

	// crete sales
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS sales (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    manager_id INTEGER NOT NULL REFERENCES managers,
    product_id INTEGER NOT NULL REFERENCES products,
    qty INTEGER NOT NULL CHECK ( qty > 0 ),
    price INTEGER NOT NULL CHECK ( price > 0 )
);`)
	if err != nil {
		return err
	}
	// adding sales
	_, err = db.Exec(`INSERT INTO sales(manager_id, product_id, price, qty)
VALUES (1, 1, 150, 10),
       (2, 2, 150, 5),
       (3, 3, 100, 5),
       (4, 1, 250, 5),
       (4, 4, 100, 5),
       (5, 5, 100, 5),
       (5, 6, 120, 10);`)

	return nil
}

