<?php

/**
 * @example vendor/bin/phpunit ed/phpUnit/examples/DBSQLLiteTest.php
 */
class DBSQLLiteTest extends PHPUnit_Framework_TestCase
{
    private static $pdo;

    public static function setUpBeforeClass()
    {
        self::$pdo = new PDO('sqlite::memory:');
        self::$pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        self::$pdo->exec('CREATE TABLE tmp (id int)');
        $stmt = self::$pdo->prepare('INSERT INTO tmp (id) values (:id)');
        $tmp = [
            [':id' => 1],
            [':id' => 2],
            [':id' => 3],
        ];
        foreach ($tmp as $id) {
            $stmt->execute($id);
        }
    }

    public function testOne()
    {
        $query = self::$pdo->prepare('select * from tmp where id > :search');
        $query->execute([':search' => 1]);
        $actualResult = [];
        foreach($query->fetchAll(PDO::FETCH_ASSOC) as $row) {
            $actualResult[] = $row;
        }
        $this->assertSame([['id' => '2'], ['id' => '3']], $actualResult);
    }
}
