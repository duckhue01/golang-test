CREATE TABLE `Todo` (
  `ID` int(6) NOT NULL AUTO_INCREMENT,
  `Title` varchar(200) NOT NULL,
  `Description` varchar(1024) DEFAULT NULL,
  `InsertAt` timestamp NOT NULL,
  `UpdateAt` timestamp NOT NULL,
  `IsDone` boolean NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID_UNIQUE` (`ID`)
);

