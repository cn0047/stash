<?php
/**
 * Iterator
 *
 * Is a pattern in which an iterator is used to traverse a container and access the container's elements.
 *
 * @category Behaviour
 */

interface IIterator
{
    public function hasNext();

    public function next();
}

interface IContainer
{
    public function createIterator();
}

class BooksCollection implements IContainer
{
    private $a_titles = array();

    public function createIterator()
    {
        return new BookIterator($this);
    }

    public function setTitle($string)
    {
        $this->a_titles[] = $string;
    }

    public function getTitles()
    {
        return $this->a_titles;
    }
}

class BookIterator implements IIterator
{
    private $i_position = 0;
    private $booksCollection;

    function __construct(BooksCollection $booksCollection)
    {
        $this->booksCollection = $booksCollection;
    }

    public function hasNext()
    {
        if ($this->i_position < count($this->booksCollection->getTitles())) {
            return true;
        }
        return false;
    }

    public function next()
    {
        $m_titles = $this->booksCollection->getTitles();

        if ($this->hasNext()) {
            return $m_titles[$this->i_position++];
        } else {
            return null;
        }
    }
}

$booksCollection = new BooksCollection();
$booksCollection->setTitle('Design Patterns');
$booksCollection->setTitle('1');
$booksCollection->setTitle('2');
$booksCollection->setTitle('3');
$iterator = $booksCollection->createIterator();
while ($iterator->hasNext()) {
    echo $iterator->next().PHP_EOL;
}

/*
Design Patterns
1
2
3
*/
