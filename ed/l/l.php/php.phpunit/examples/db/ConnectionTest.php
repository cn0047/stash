<?php

class ConnectionTest extends PHPUnit_Extensions_Database_TestCase
{
    /**
     * @return PHPUnit_Extensions_Database_DB_IDatabaseConnection
     */
    public function getConnection()
    {
        $database = 'myguestbook';
        $user = 'root';
        $password = '';
        $pdo = new PDO('mysql:host=localhost;dbname=myguestbook', $user, $password);
        $pdo->exec('CREATE TABLE IF NOT EXISTS guestbook (id int, content text, user text, created text)');
        return $this->createDefaultDBConnection($pdo, $database);
        // $pdo = new PDO('sqlite::memory:');
        // $pdo->exec('CREATE TABLE guestbook (id int, content text, user text, created text)');
        // $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        // return $this->createDefaultDBConnection($pdo, ':memory:');
    }

    /**
     * @return PHPUnit_Extensions_Database_DataSet_IDataSet
     */
    public function getDataSet()
    {
        return $this->createFlatXMLDataSet(__DIR__.'/dataSets/myFlatXmlFixture.xml');
        // return new PHPUnit_Extensions_Database_DataSet_YamlDataSet(__DIR__.'/dataSets/guestbook.yml');
        // return $this->createFlatXMLDataSet(__DIR__.'/dataSets/myXmlFixture.xml');
    }

    // public function testCreateDataSet()
    // {
    //     $tableNames = array('guestbook');
    //     $dataSet = $this->getConnection()->createDataSet();
    // }

    // public function testCreateQueryTable()
    // {
    //     $tableNames = array('guestbook');
    //     $queryTable = $this->getConnection()->createQueryTable('guestbook', 'SELECT * FROM guestbook');
    // }

    public function testGetRowCount()
    {
        $this->assertEquals(2, $this->getConnection()->getRowCount('guestbook'));
    }
}
