<?php
/**
 * Dependency injection
 *
 * @category Behaviour
 */

/**
 * Interface injection.
 */
interface Enquire
{
    public function initAuthor(Author $author);
}

class Author
{
    private $firstName;
    private $lastName;

    public function __construct($firstName, $lastName)
    {
        $this->firstName = $firstName;
        $this->lastName = $lastName;
    }

    public function getFirstName()
    {
        return $this->firstName;
    }

    public function getLastName()
    {
        return $this->lastName;
    }
}

class Question implements Enquire
{
    private $author;
    private $question;

    /**
     * Constructor injection.
     */
    public function __construct($question, Author $author)
    {
        $this->author = $author;
        $this->question = $question;
    }

    /**
     * Setter injection.
     */
    public function setAuthor(Author $author)
    {
        $this->author = $author;
    }

    /**
     * Interface injection.
     */
    public function initAuthor(Author $author)
    {
        $this->author = $author;
    }

    public function getAuthor()
    {
        return $this->author;
    }

    public function getQuestion()
    {
        return $this->question;
    }
}
