CREATE TABLE `Calendar` (
	`Id` int NOT NULL AUTO_INCREMENT,
	`Name` varchar(100) NOT NULL,
	`Year` int NOT NULL,
	PRIMARY KEY (`Id`)
);

CREATE TABLE `Month` (
	`Id` int NOT NULL AUTO_INCREMENT,
	`Name` varchar(15) NOT NULL,
	`CalendarId` int NOT NULL,
	PRIMARY KEY (`Id`)
);

CREATE TABLE `Event` (
	`Id` int NOT NULL AUTO_INCREMENT,
	`Title` varchar(100) NOT NULL,
	`Description` TEXT NOT NULL,
	`MonthId` int NOT NULL,
	`Day` tinyint NOT NULL,
	PRIMARY KEY (`Id`)
);

ALTER TABLE `Month` ADD CONSTRAINT `Month_fk0` FOREIGN KEY (`CalendarId`) REFERENCES `Calendar`(`Id`);

ALTER TABLE `Event` ADD CONSTRAINT `Event_fk0` FOREIGN KEY (`MonthId`) REFERENCES `Month`(`Id`);

CREATE INDEX idx_calendar_created ON Calendar(Id);
CREATE INDEX idx_month_created ON Month(Id);
CREATE INDEX idx_event_created ON Event(Id);