<?php
/**
 * Adapter (wrapper)
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
