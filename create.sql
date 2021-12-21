CREATE TABLE `Todo` (
  `Id` int(6) NOT NULL AUTO_INCREMENT,
  `Title` varchar(200) NOT NULL,
  `Description` varchar(1024) DEFAULT NULL,
  `CreateAt` timestamp NOT NULL,
  `UpdateAt` timestamp NOT NULL,
  `IsDone` boolean NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `ID_UNIQUE` (`Id`)
);

