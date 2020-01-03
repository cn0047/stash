<?php

class CarArrayForm extends CFormModel
{
    public $name;
    public $country;

    public function rules()
    {
        return [
            ['name', 'type', 'type' => 'string'],
            ['country', 'type', 'type' => 'string'],
        ];
    }
}
