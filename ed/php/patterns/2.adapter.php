<?php
/**
 * Adapter (wrapper)
 *
 * To translate one interface for a class into a compatible interface.
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
