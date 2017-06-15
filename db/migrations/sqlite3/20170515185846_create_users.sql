-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `Users` (
	`Id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`Name`	TEXT NOT NULL,
	`Guid`	TEXT NULL,
	`Email`	TEXT NULL,
	`Password`	TEXT NOT NULL,
	`Salt`	TEXT  NULL,
	`JoinedDate` TEXT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Users;