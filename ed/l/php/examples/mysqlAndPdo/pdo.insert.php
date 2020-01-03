<?php

$dbh = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');

/**
 * VARIANT 1
 *
 * Everything OK.
 */
$s = $dbh->prepare('INSERT INTO stockists VALUES (null, :s1, :s2)');
$s->bindValue(':s1', uniqid(), PDO::PARAM_STR);
$s->bindValue(':s2', uniqid(), PDO::PARAM_STR);
if (!$s->execute()) {
    var_dump($s->errorInfo());
} else {
    echo 'VARIANT 1 - OK, at: '.time();
}

/**
 * VARIANT 2
 *
 * Everything OK.
 */
$s = $dbh->prepare('INSERT INTO stockists VALUES (null, :s1, :s2)');
if (!$s->execute([':s1' => uniqid(), ':s2' => uniqid()])) {
    var_dump($s->errorInfo());
} else {
    echo 'VARIANT 2 - OK, at: '.time();
}

/**
 * VARIANT 3
 *
 * Everything OK.
 */
$s = $dbh->prepare('INSERT INTO stockists (state1, state2) VALUES (:s1, :s2)');
if (!$s->execute([':s1' => uniqid(), ':s2' => uniqid()])) {
    var_dump($s->errorInfo());
} else {
    echo 'VARIANT 3 - OK, at: '.time();
}

/**
 * VARIANT 4
 *
 * Everything OK.
 */
$s = $dbh->prepare('INSERT INTO stockists SET state1 = :s1, state2 = :s2');
if (!$s->execute([':s1' => uniqid(), ':s2' => uniqid()])) {
    var_dump($s->errorInfo());
} else {
    echo 'VARIANT 4 - OK, at: '.time();
}

/**
 * VARIANT 5
 *
 * Everything OK.
 */
$s = $dbh->prepare('INSERT INTO stockists SET state1 = ?, state2 = ?');
if (!$s->execute([uniqid(), uniqid()])) {
    var_dump($s->errorInfo());
} else {
    echo 'VARIANT 5 - OK, at: '.time();
}

/**
 * VARIANT 6
 *
 * PHP Warning:  PDO::query(): SQLSTATE[HY000]: General error: mode must be an integer in ...
 */
// $query = 'INSERT INTO stockists SET state1 = ?, state2 = ?';
// $result = $dbh->query($query, [uniqid(), uniqid()]);
// var_export($result);
