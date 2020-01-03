<?php

/*
select * from user;
+----+----------------+------+--------------------------------------------------------------+
| id | email          | name | password                                                     |
+----+----------------+------+--------------------------------------------------------------+
|  1 | test@gmail.com | Test | $2y$10$/UlSlT.bCntE1L2099KnmeM4I1v1Vjr/278UDxIghCeBndVnfjD.y |
|  2 | bond@mi6.com   | Bond | $2y$10$DK4NjvRPICtM57e..V7hlOr6sVKAvwbm1M3u.uFrJsLIWkO6R179m |
+----+----------------+------+--------------------------------------------------------------+
*/

class User
{
    public function __toString()
    {
        return sprintf('Name: %s, Email: %s'.PHP_EOL, $this->name, $this->email);
    }
}

$dbh = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');
$s = $dbh->prepare('SELECT * FROM user');
$s->execute();
$result = $s->fetchAll(PDO::FETCH_CLASS, 'User');
foreach ($result as $obj) {
    echo $obj;
}

/*
Name: Test, Email: test@gmail.com
Name: Bond, Email: bond@mi6.com
*/
