<?php

class Car extends CActiveRecord
{
    public function tableName()
    {
        return 'cars';
    }

    public function rules()
    {
        return [
            ['brand, model, maxSpeed', 'required', 'on' => 'insert'],
            ['brand', 'type', 'type' => 'string'],
            ['model', 'type', 'type' => 'string'],
            ['maxSpeed', 'numerical'],
        ];
    }

    public function search()
    {
        $criteria = new CDbCriteria;
        $criteria->compare('brand', $this->brand, true);
        $criteria->compare('model', $this->model, true);
        $criteria->compare('maxSpeed', $this->maxSpeed, true);
        return new CActiveDataProvider($this, array(
            'criteria' => $criteria,
            'pagination' => array('pageSize' => 20),
        ));
    }
}
