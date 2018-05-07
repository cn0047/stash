<?php
/**
 * Strategy
 *
 * @category Behaviour
 */

interface Comparator
{
    public function compare($a, $b);
}

class NumericComparator implements Comparator
{
    public function compare($a, $b)
    {
        if ($a == $b) {
            return 0;
        }
        return $a < $b ? -1 : 1;
    }
}

class CountableObjectComparator implements Comparator
{
    public function compare($a, $b)
    {
        if (count($a) == count($b)) {
            return 0;
        }
        return count($a) < count($b) ? -1 : 1;
    }
}

class Collection implements Countable
{
    private $elements;
    private $comparator;

    public function __construct(array $elements = array())
    {
        $this->elements = $elements;
    }

    public function initComparator(Comparator $comparator)
    {
        $this->comparator = $comparator;
    }

    public function sort()
    {
        $callback = array($this->comparator, 'compare');
        uasort($this->elements, $callback);
    }

    public function __toString()
    {
        $elements = array();
        foreach ($this->elements as $value) {
            if (is_array($value)) {
                $value = 'Array with ' . count($value) . ' elements';
            }
            $elements[] = $value;
        }
        return '(' . implode(', ', $elements) . ')';
    }

    public function count()
    {
        return count($this->elements);
    }
}

// ordering numbers
$numbers = new Collection(array(4, 6, 1, 7, 3));
$numbers->initComparator(new NumericComparator);
$numbers->sort();
echo $numbers, "\n";
// ordering Countable objects
$first = array(1, 2, 3);
$second = array(1, 2, 3, 4);
$third = new Collection(array(4, 2, 3, 5, 1));
$objects = new Collection(array($third, $second, $first));
$objects->initComparator(new CountableObjectComparator);
$objects->sort();
echo $objects, "\n";

/*
(1, 3, 4, 6, 7)
(Array with 3 elements, Array with 4 elements, (4, 2, 3, 5, 1))
*/
