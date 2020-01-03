<?php

Yii::import('application.controllers.MiscController');

class MiscControllerTest extends CTestCase
{
    public function testGetQRCode1()
    {
        $this->assertEquals(404, Yii::app()->QRCode->get());
    }

    public function testGetQRCode2()
    {
        $controller = new MiscController('MiscControllerTest');
    }
}
