<?php

class Helper
{
    public function getAvailableBrands()
    {
        $brand = new Brand;
        $sql = 'SELECT id, name FROM '.$brand->tableName();
        $command = Yii::app()->db->getPdoInstance()->prepare($sql);
        $command->execute();
        $data = $command->fetchAll(\PDO::FETCH_GROUP|\PDO::FETCH_ASSOC);
        foreach ($data as &$el) {
            $el = $el[0]['name'];
        }
        return $data;
    }

    public function getAvailableModels()
    {
        $car = new Car;
        $data = Yii::app()->db->createCommand()
            ->selectDistinct(['model'])
            ->from($car->tableName())
            ->queryColumn()
            ;
        $data = array_combine($data, $data);
        return $data;
    }
}
