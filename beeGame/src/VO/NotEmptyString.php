<?php

namespace VO;

class NotEmptyString
{
    private $value;

    public function __construct($value)
    {
        if (!is_string($value) or $value === '') {
            throw new \InvalidArgumentException('Invalid value.');
        }
        $this->value = $value;
    }

    /**
     * @return string
     */
    public function get()
    {
        return $this->value;
    }

    public function __toString()
    {
        return $this->value;
    }
}
