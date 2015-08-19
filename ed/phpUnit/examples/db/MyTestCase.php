<?php

class MyTestCase extends PHPUnit_Extensions_Database_TestCase
{
    public function getDataSet()
    {
        // XML
        return $this->createFlatXmlDataSet('dataSets/myFlatXmlFixture.xml');
        return $this->createXMLDataSet('dataSets/myXmlFixture.xml');
        // Run: mysqldump --xml -t -u root --password= testdrive > /tmp/testdrive.xml
        // then:
        return $this->createMySQLXMLDataSet('/tmp/testdrive.xml');
        // YAML
        return new PHPUnit_Extensions_Database_DataSet_YamlDataSet('dataSets/guestbook.yml');
        // CSV
        $dataSet = new PHPUnit_Extensions_Database_DataSet_CsvDataSet();
        $dataSet->addTable('guestbook', 'dataSets/guestbook.csv');
        return $dataSet;
        // ARRAY
        return new MyApp_DbUnit_ArrayDataSet([
            'guestbook' => [
                ['id' => 1, 'content' => 'Hello buddy!', 'user' => 'joe', 'created' => '2010-04-24 17:15:23'],
                ['id' => 2, 'content' => 'I like it!',   'user' => null,  'created' => '2010-04-26 12:14:20'],
            ],
        ]);
    }
}