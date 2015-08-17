<?php

class Vendor extends CActiveRecord
{
    public function tableName()
    {
        return '';
    }

    public function rules()
    {
        return [
            ['name', 'type', 'type' => 'string'],
        ];
    }
}
