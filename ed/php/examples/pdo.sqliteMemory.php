<?php

header('Content-type: text/plain');
$pdo = new PDO('sqlite::memory:');
$pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
$result = $pdo->exec('CREATE TABLE tmp (id int)');
$stmt = $pdo->prepare('INSERT INTO tmp (id) values (:id)');
try {
    $tmp = [
        [':id' => 1],
        [':id' => 2],
        [':id' => 3],
    ];
    foreach ($tmp as $id) {
        $stmt->execute($id);
    }
} catch (PDOException $e) {
    echo $e;
}
$query = $pdo->prepare('select * from tmp where id > :search');
$query->execute([':search' => 1]);
foreach($query->fetchAll(PDO::FETCH_ASSOC) as $row) {
    var_export($row);
}
