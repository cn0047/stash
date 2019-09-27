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

    public function actionArray()
    {
        $get['name'] = \Yii::app()->request->getQuery('name');
        if (!preg_match('/^\w+$/', $get['name'])) {
            var_dump('Error.');
            $get['name'] = '';
        }
        $dataProvider = new CArrayDataProvider(
            [
                [
                    'id' => '1',
                    'name' => 'Lexus',
                    'country' => 'JPN',
                ],
                [
                    'id' => '2',
                    'name' => 'Acura',
                    'country' => 'JPNE',
                ],
                [
                    'id' => '3',
                    'name' => 'Lexus 2',
                    'country' => 'JPN',
                ],
                [
                    'id' => '5',
                    'name' => 'Acura 2',
                    'country' => 'JPNE',
                ],
                [
                    'id' => '5',
                    'name' => 'Lexus 3',
                    'country' => 'JPN',
                ],
                [
                    'id' => '6',
                    'name' => 'Acura 3',
                    'country' => 'JPNE',
                ],
                [
                    'id' => '7',
                    'name' => 'Lexus 4',
                    'country' => 'JPN',
                ],
                [
                    'id' => '8',
                    'name' => 'Acura 4',
                    'country' => 'JPNE',
                ],
                [
                    'id' => '9',
                    'name' => 'Lexus 5',
                    'country' => 'JPN',
                ],
                [
                    'id' => '10',
                    'name' => 'Acura 5',
                    'country' => 'JPNE',
                ],
            ]
        );
        $this->render(
            'indexArray',
            ['dataProvider' => $dataProvider, 'get' => $get]
        );
    }
}
