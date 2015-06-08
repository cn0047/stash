<?php

class CarController extends Controller
{
    public function actionIndex()
    {
        $model = new Car('search');
        $model->unsetAttributes();
        if (isset($_GET['Car'])) {
            $model->attributes = $_GET['Car'];
        }
        $this->render('index', array('model' => $model));
    }
}
