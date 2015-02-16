<?php
/**
 * Dependency injection
 *
 * Pattern separates the creation of a client's dependencies from its own behavior,
 * which allows program designs to be loosely coupled and to follow the
 * dependency inversion and single responsibility principles.
 *
 * @category Behaviour
 */

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

class Question
{
    private $author;
    private $question;

    public function __construct($question, Author $author)
    {
        $this->author = $author;
        $this->question = $question;
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
