<?php

class MiscController extends Controller
{
    public function actionIndex()
    {
        $form = new TestForm;
        $form->ids = [
            5,
            'a',
        ];
        $form->adIds = [
            'f895a078e81c11e4a5ee11aa24b31963',
            7,
        ];
        $d = [
            $form->validate(),
            $form->getErrors(),
        ];
        $this->render('index', ['d' => $d]);
    }

    public function actionGetQRCode()
    {
        Yii::app()->QRCode->onGet = [$this, 'onGetQRCode'];
        $d = Yii::app()->QRCode->get();
        $this->render('index', ['d' => $d]);
    }

    public function onGetQRCode($event)
    {
        var_dump(['captured event' => $event]);
    }
}
