<?php

use Phalcon\Mvc\Model;

class Robots extends Model
{
    public $id;

    public $name;

    public function initialize()
    {
        $this->hasMany('id', 'RobotsParts', 'robots_id');
    }
}
