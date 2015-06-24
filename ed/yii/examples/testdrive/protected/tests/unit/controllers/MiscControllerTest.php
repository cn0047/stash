<?php

class MiscControllerTest extends CTestCase
{
    public function testPutData()
    {
        $this->assertEquals(404, Yii::app()->QRCode->get());
    }
}
