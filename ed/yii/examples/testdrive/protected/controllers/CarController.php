<?php

class CarController extends Controller
{
    public function actionIndex()
    {
        $car = new Car('search');
        $query = Yii::app()->request->getQuery(get_class($car));
        if (!is_null($query)) {
            $car->attributes = $query;
            if (!$car->validate()) {
                $car->unsetAttributes();
            }
        }
        $helper = new Helper;
        $templete = 'index';
        if (Yii::app()->request->getQuery('add')) {
            $templete = 'indexAdd';
        }
        $this->render($templete, ['model' => $car, 'helper' => $helper]);
    }
}
