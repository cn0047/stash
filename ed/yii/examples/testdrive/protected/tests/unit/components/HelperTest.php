<?php

class HelperTest extends CDbTestCase
{
    protected $fixtures = [
        'brand' => 'Brand',
    ];

    public function testGetAvailableBrands()
    {
        $actualBrands = (new Helper)->getAvailableBrands();
        $this->assertContains('Lexus', $actualBrands);
        $this->assertContains('Acura', $actualBrands);
    }

    /**
     * @expectedException CDbException CDbCommand failed to execute the SQL statement: SQLSTATE[42S01]: Base table or view already exists: 1050 Table 'brand' already exists
     */
    public function testGetAvailableBrands2()
    {
        $sql = <<<"SQL"
            CREATE TABLE brand (
                id INT AUTO_INCREMENT,
                name VARCHAR(100) NOT NULL DEFAULT '',
                country VARCHAR(50) NOT NULL DEFAULT '',
                PRIMARY KEY (id)
            );
            INSERT INTO brand VALUES
                (null, 'Lamborghini', 'IT')
            ;
SQL;
        \Yii::app()->db->connectionString = 'sqlite::memory:';
        \Yii::app()->db->setActive(true);
        \Yii::app()->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        \Yii::app()->db->createCommand($sql)->execute();
    }

    public function testGetAllAvailableBrands()
    {
        // Get brands from fixture.
        $brand = new Brand;
        $allBrandsFromFixture = $brand->findAll();
        // Generate brands.
        $lexus = new Brand;
        $lexus->scenario = 'update';
        $lexus->name = 'Lexus';
        $lexus->country = 'JPN';
        $lexus->id = 1;
        $lexus->setPrimaryKey(1);
        $lexus->setIsNewRecord(false);
        $allBrandsGeneratedAtRuntime[] = $lexus;
        $acura = new Brand;
        $acura->scenario = 'update';
        $acura->name = 'Acura';
        $acura->country = 'JPNE';
        $acura->id = 2;
        $acura->setPrimaryKey(2);
        $acura->setIsNewRecord(false);
        $allBrandsGeneratedAtRuntime[] = $acura;
        // Brands from fixture should be equals to generated brands.
        $this->assertEquals($allBrandsFromFixture, $allBrandsGeneratedAtRuntime);
    }
}
