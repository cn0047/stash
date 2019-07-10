<?php

// default
// $dbh = new PDO('mysql:host=127.0.0.1;dbname=test', 'root');
// docker
// $dbh = new PDO('mysql:host=mysql-master;port=3306;dbname=test', 'dbu', 'dbp');
$dbh = new PDO('mysql:host=xmysql;port=3306;dbname=test', 'dbu', 'dbp'); // OK
// $dbh = new PDO('mysql:host=xmysql;port=3306;dbname=testdrive', 'user', 'pass'); // OK
// $dbh = new PDO('mysql:host=xmysql;port=3306;dbname=testdrive_unit_test', 'user', 'pass'); // OK
$s = $dbh->prepare('SELECT NOW()');
$s->execute();
$result = $s->fetchAll(PDO::FETCH_ASSOC);
var_export($result);
