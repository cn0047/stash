<?php

/*
create table mostCurrentRow (
    id int auto_increment,
    week date,
    firstName char(5),
    lastName char(5),
    score int,
    primary key (id)
);
insert into mostCurrentRow values
(1, '2015-03-06', 'Bill', 'Jones', 25),
(2, '2015-03-13', 'Bill', 'Jones', 45),
(3, '2015-03-06', 'Dave', 'Smith', 32),
(4, '2015-03-13', 'Dave', 'Smith', 52),
(5, '2015-03-06', 'Mary', 'Chu'  , 28),
(6, '2015-03-13', 'Mary', 'Chu'  , 45);
*/

$link = mysql_connect('localhost', 'root');
$selectedDB = mysql_select_db('test', $link);

$date = mysql_escape_string('2015-03-13');
$myQuery = "SELECT * FROM mostCurrentRow WHERE week = '$date' ORDER BY score DESC";
$query = mysql_query($myQuery);
if (!$myQuery) {
    throw new Exception(mysql_error());
}
$data = [];
while ($row = mysql_fetch_assoc($query)) {
    $data[] = $row;
}
echo '<pre>';
var_export($data);

mysql_close();
