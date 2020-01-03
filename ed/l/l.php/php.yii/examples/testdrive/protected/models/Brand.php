<?php

class Brand extends CActiveRecord
{
    public function tableName()
    {
        return 'brand';
    }

    public function relations()
    {
        return [
            'car' => [self::HAS_MANY, 'Car', 'brand_id'],
        ];
    }

    public function rules()
    {
        return [
            ['name, country', 'on' => 'insert'],
            ['name', 'type', 'type' => 'string'],
            ['name', 'length', 'max' => 100],
            ['country', 'type', 'type' => 'string'],
            ['country', 'length', 'max' => 50],
        ];
    }
}
