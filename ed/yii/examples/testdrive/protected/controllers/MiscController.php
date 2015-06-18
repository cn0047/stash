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
        echo '<pre>';
        var_export([
            $form->validate(),
            $form->getErrors(),
        ]);
        echo '</pre>';
    }
}
