<?php

namespace VO;

/**
 * Positive integer wrapper.
 *
 * This class ensure that it contains valid and positive integer.
 */
class PositiveInteger
{
    /** @var integer $value Value. */
    private $value;

    /**
     * Constructor.
     *
     * @param integer $value Value.
     *
     * @throws \InvalidArgumentException In case when value is invalid.
     */
    public function __construct($value)
    {
        if (!is_int($value) or $value < 0) {
            throw new \InvalidArgumentException('Invalid value.');
        }
        $this->value = $value;
    }

    /**
     * Gets value.
     *
     * @return integer Value.
     */
    public function get()
    {
        return $this->value;
    }
}
