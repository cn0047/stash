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
}
