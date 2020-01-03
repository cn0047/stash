<?php

require_once 'PHPUnit/Extensions/SeleniumTestCase.php';

class STest extends PHPUnit_Extensions_SeleniumTestCase
{
    protected function setUp()
    {
        $this->setBrowser('*firefox');
        $this->setBrowserUrl('http://www.example.com/');
    }

    public function testTitle()
    {
        $this->open('http://www.example.com/');
        $this->assertTitle('Example Domain');
        $this->click('link=More information...');
        $this->waitForPageToLoad('30000');
    }
}
?>