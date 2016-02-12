<?php

namespace VO;

/**
 * Not empty string wrapper.
 *
 * This class ensure that it contains valid and not empty string.
 */
class NotEmptyString
{
    /** @var string $value Value. */
    private $value;

    /**
     * Constructor.
     *
     * @param string $value Value.
     *
     * @throws \InvalidArgumentException In case when value is invalid.
     */
    public function __construct($value)
    {
        if (!is_string($value) or $value === '') {
            throw new \InvalidArgumentException('Invalid value.');
        }
        $this->value = $value;
    }

    /**
     * Gets value.
     *
     * @return string Value.
     */
    public function get()
    {
        return $this->value;
    }

    /**
     * Magic gets value.
     *
     * @return string Value.
     */
    public function __toString()
    {
        return $this->value;
    }
}
