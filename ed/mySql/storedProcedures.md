Stored procedures
-

````sql
create table countries (id int auto_increment primary key, name varchar(100));
insert into countries values
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

call getCountry('ua');
+----+------+
| id | name |
+----+------+
|  3 | ua   |
+----+------+
````
