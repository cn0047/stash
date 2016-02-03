<?php

namespace VO;

class PositiveInteger
{
    private $value;

    public function __construct($value)
    {
        if (!is_int($value) or $value < 0) {
            throw new \InvalidArgumentException('Invalid value.');
        }
        $this->value = $value;
    }

    public function get()
    {
        return $this->value;
    }
}
