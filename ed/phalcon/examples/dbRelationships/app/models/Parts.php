<?php

use Phalcon\Mvc\Model;

class Parts extends Model
{
    public $id;

    public $name;

    public function initialize()
    {
        $this->hasMany('id', 'RobotsParts', 'parts_id');
    }
}
