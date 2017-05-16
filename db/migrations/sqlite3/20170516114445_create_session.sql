
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `Sessions` (
	`Id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`SessionId`	TEXT NOT NULL,
	`UserId`	INTEGER NOT NULL,
	`SessionStart`	TEXT NOT NULL,
	`SessionUpdate`	TEXT NOT NULL,
	`SessionActive`	NUMERIC
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE Sessions