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
