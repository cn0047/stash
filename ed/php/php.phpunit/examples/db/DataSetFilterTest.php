<?php

class DataSetFilterTest extends PHPUnit_Extensions_Database_TestCase
{
    /**
     * @return PHPUnit_Extensions_Database_DB_IDatabaseConnection
     */
    public function getConnection()
    {
        $pdo = new PDO('sqlite::memory:');
        $pdo->exec('CREATE TABLE guestbook (id int, content text, user text, created text)');
        $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        return $this->createDefaultDBConnection($pdo, ':memory:');
    }

    /**
     * @return PHPUnit_Extensions_Database_DataSet_IDataSet
     */
    public function getDataSet()
    {
        return $this->createFlatXMLDataSet(__DIR__.'/dataSets/myFlatXmlFixture.xml');
    }

    public function testIncludeFilteredGuestbook()
    {
        $tableNames = array('guestbook');
        $dataSet = $this->getConnection()->createDataSet();

        $filterDataSet = new PHPUnit_Extensions_Database_DataSet_DataSetFilter($dataSet);
        $filterDataSet->addIncludeTables(array('guestbook'));
        $filterDataSet->setIncludeColumnsForTable('guestbook', array('id', 'content'));
        // ..
    }

    public function testExcludeFilteredGuestbook()
    {
        $tableNames = array('guestbook');
        $dataSet = $this->getConnection()->createDataSet();

        $filterDataSet = new PHPUnit_Extensions_Database_DataSet_DataSetFilter($dataSet);
        $filterDataSet->addExcludeTables(array('foo', 'bar', 'baz')); // only keep the guestbook table!
        $filterDataSet->setExcludeColumnsForTable('guestbook', array('user', 'created'));
        // ..
    }
}
