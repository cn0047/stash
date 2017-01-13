<?php
/**
 * Adapter (wrapper)
 *
 * To translate one interface for a class into a compatible interface.
 * An adapter allows classes to work together
 * that normally could not because of incompatible interfaces
 * by providing its interface to clients while using the original interface.
 *
 * @category Structural
 */

class BookAdapter
{
    protected $book;

    public function __construct($book)
    {
        $this->book = $book;
    }

    public function open()
    {
        $this->book->pressStart();
    }

    public function turnPage()
    {
        $this->book->pressNext();
    }
}
