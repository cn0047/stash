<?php
/**
 * Composite
 *
 * The intent of a composite is to "compose" objects into tree structures.
 *
 * @category Structural
 */

class Node
{
    private $elements = [];

    public function add($el)
    {
        $this->elements[] = $el;
    }

    public function show()
    {
        foreach ($this->elements as $el) {
            $el->show();
        }
    }
}

class Leaf
{
    private $name = '';

    public function __construct($n)
    {
        $this->name = $n;
    }

    public function show()
    {
        var_dump($this->name);
    }
}

$n = new Node();
$n->add(new Leaf('first l'));
$n->add(new Leaf('2nd l'));
$n1 = new Node();
$n->add(new Leaf('NLLL --- first l'));
$n->add(new Leaf('NLLL --- 2nd l'));
$n->add($n1);
$n->add(new Leaf('3nd l'));
$n->show();

/*
string(7) "first l"
string(5) "2nd l"
string(16) "NLLL --- first l"
string(14) "NLLL --- 2nd l"
string(5) "3nd l"
*/

/**
 * Example 2.
 */

abstract class Component
{
    protected $name;

    public function __construct($name)
    {
        $this->name = $name;
    }

    abstract public function add(Component $c);
    abstract public function remove(Component $c);
    abstract public function display();
}

class Composite extends Component
{
    private $children = array();

    public function add(Component $component)
    {
        $this->children[$component->name] = $component;
    }

    public function remove(Component $component)
    {
        unset($this->children[$component->name]);
    }

    public function display()
    {
        foreach ($this->children as $child) {
            $child->display();
        }
    }
}

class Leaf extends Component
{
    public function add(Component $c)
    {
        print ('Cannot add to a leaf');
    }

    public function remove(Component $c)
    {
        print('Cannot remove from a leaf');
    }

    public function display()
    {
        var_dump($this->name);
    }
}

$root = new Composite('Composite');
$root->add(new Leaf('Leaf A'));
$root->add(new Leaf('Leaf B'));

$comp = new Composite('Composite X');
$comp->add(new Leaf('Leaf XA'));
$comp->add(new Leaf('Leaf XB'));

$root->add($comp);
$root->add(new Leaf('Leaf C'));

$leaf = new Leaf('Leaf D');
$root->add($leaf);
$root->remove($leaf);

$root->display();

/*
string(6) "Leaf A"
string(6) "Leaf B"
string(7) "Leaf XA"
string(7) "Leaf XB"
string(6) "Leaf C"
*/
