CREATE TABLE `Todo` (
  `Id` int(32) NOT NULL,
  `Title` varchar(200) NOT NULL,
  `Description` varchar(1024) DEFAULT NULL,
  `CreateAt` timestamp NOT NULL,
  `UpdateAt` timestamp NOT NULL,
  `Status` int(2) NOT NULL,
  `Order` int(32) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`Id`),
  INDEX(`Order`)
);

CREATE TABLE `Tag` (
  `Tag` varchar(100) NOT NULl,
  PRIMARY KEY (`Tag`),
  UNIQUE KEY `Tag` (`Tag`)
);

CREATE TABLE `TagTodo` (
  `Tag` varchar(100) NOT NULl,
  `Id` int(32) NOT NULL,
  FOREIGN KEY (`Tag`) REFERENCES `Tag`(`Tag`) ON UPDATE CASCADE ON DELETE CASCADE, 
  FOREIGN KEY (`Id`) REFERENCES `Todo`(`Id`) ON UPDATE CASCADE ON DELETE CASCADE,
  PRIMARY KEY (`Id`, `Tag`),
  UNIQUE KEY `TagId` (`Tag`, `Id`)
);



DELIMITER //

CREATE PROCEDURE `AddTagIfNotExist`(IN tag varchar(100) )
BEGIN
  INSERT IGNORE INTO `Tag`(`Tag`)
  VALUES (tag);
END //

DELIMITER ;


DELIMITER //

CREATE PROCEDURE `AddTagTodoIfNotExist`(IN tag varchar(100), IN id int(32) )
BEGIN
  INSERT IGNORE INTO `TagTodo`(`Tag`, `Id`)
  VALUES (tag, id);
END //

DELIMITER ;



DELIMITER //

CREATE PROCEDURE `ReOrderTodo`(IN id int(32), IN pos int(32) )
BEGIN
  -- SELECT @s= (SELECT `Order`  FROM `Todo` WHERE `Id`= id LIMIT 1);

  IF (id > pos) THEN

    UPDATE `Todo`
    SET `Order`=0 
    WHERE `Order`=id;

   
    UPDATE `Todo`
    SET `Order`=`Order`+1
    WHERE  `Order`>=pos AND  `Order` <  id;

    UPDATE `Todo`
    SET `ORder`=pos 
    WHERE `Order`=0;

  ELSEIF (id < pos) THEN
    UPDATE `Todo`
    SET `Order`=0 
    WHERE `Order`=id;

    
    UPDATE `Todo`
    SET `Order`=`Order`- 1
    WHERE  `Order`<=pos AND  `Order` > id;

    UPDATE `Todo`
    SET `Order`=pos 
    WHERE `Order`=0;
  END IF;

END //

DELIMITER ;




