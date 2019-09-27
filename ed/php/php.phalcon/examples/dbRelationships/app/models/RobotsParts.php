<?php

use Phalcon\Mvc\Model;

class RobotsParts extends Model
{
    public $id;

    public $robots_id;

    public $parts_id;

    public function initialize()
    {
        $this->belongsTo('robots_id', 'Robots', 'id');
        $this->belongsTo('parts_id', 'Parts', 'id');
    }
}
