CREATE TABLE countries (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100));
INSERT INTO countries VALUES
 (null, 'usa')
,(null, 'gb')
,(null, 'ua')
,(null, 'canada')
,(null, 'brazil')
;

DELIMITER //
CREATE PROCEDURE getCountry (IN countryName VARCHAR(255))
    BEGIN
    SELECT * 
    FROM countries
    WHERE name = countryName;
    END //
DELIMITER ;

SHOW CREATE PROCEDURE getCountry;

CALL getCountry('ua');
+----+------+
| id | name |
+----+------+
|  3 | ua   |
+----+------+

-- Example with loop
DROP PROCEDURE IF EXISTS getHi;

DELIMITER //
CREATE PROCEDURE getHi (IN count INT)
    BEGIN
    DECLARE i INT;
    SET i = 1;
    WHILE i  <= count DO
      SELECT i;
      SET  i = i + 1; 
    END WHILE;
    END //
DELIMITER ;

CALL getHi(3);
