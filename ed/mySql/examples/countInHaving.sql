CREATE TABLE author (
    id int,
    name_author text,
    PRIMARY KEY (id)
) ENGINE=InnoDB;

INSERT INTO author VALUES
(1 , 'david'),
(2 , 'kate'),
(3 , 'tom'),
(4 , 'mark');

CREATE TABLE books (
    id int,
    name_book text,
    PRIMARY KEY (id)
) ENGINE=InnoDB;

INSERT INTO books VALUES
(1 , 'book1'),
(2 , 'book2'),
(3 , 'book3');

CREATE TABLE relationships (
    id int AUTO_INCREMENT,
    id_book int,
    id_author int,
    PRIMARY KEY (id)
) ENGINE=InnoDB;

INSERT INTO relationships VALUES
(null, 1 , 2),
(null, 1 , 3),
(null, 1 , 4),
(null, 2 , 2),
(null, 1 , 1),
(null, 3 , 4);

SELECT
    b.name_book, GROUP_CONCAT(a.name_author) authors
FROM relationships r
JOIN books b ON r.id_book = b.id
JOIN author a ON r.id_author = a.id
GROUP BY r.id_book
HAVING COUNT(r.id_book) = 4
;
+-----------+---------------------+
| name_book | authors             |
+-----------+---------------------+
| book1     | kate,tom,mark,david |
+-----------+---------------------+

DROP TABLE author;
DROP TABLE books;
DROP TABLE relationships;
