<?php
/**
 * Chain of Responsibility
 *
 * Is a design pattern consisting of objects, each object contains logic and link to the next
 * processing object in the chain, that will invoked.
 *
 * @category Behaviour
 */

abstract class AbstractHandler
{
    protected $next;

    abstract public function sendRequest($message);

    public function setNext($next)
    {
        $this->next = $next;
    }

    public function getNext()
    {
        return $this->next;
    }
}

class ConcreteHandlerA extends AbstractHandler
{
    public function sendRequest($message)
    {
        echo __CLASS__ . " try process this message\n";
        if ($message == 1) {
            echo __CLASS__ . " process this message !!!\n";
        } else {
            if ($this->getNext()) {
                $this->getNext()->sendRequest($message);
            }
        }
    }
}

class ConcreteHandlerB extends AbstractHandler
{
    public function sendRequest($message)
    {
        echo __CLASS__ . " try process this message\n";
        if ($message == 2) {
            echo __CLASS__ . " process this message !!!\n";
        } else {
            if ($this->getNext()) {
                $this->getNext()->sendRequest($message);
            }
        }
    }
}

class ConcreteHandlerC extends AbstractHandler
{
    public function sendRequest($message)
    {
        echo __CLASS__ . " try process this message\n";
        if ($message == 3) {
            echo __CLASS__ . " process this message !!!\n";
        } else {
            if ($this->getNext()) {
                $this->getNext()->sendRequest($message);
            }
        }
    }
}
$handler = new ConcreteHandlerA();
$handler->setNext(new ConcreteHandlerB());
$handler->getNext()->setNext(new ConcreteHandlerC());
$handler->sendRequest(3);

/*
ConcreteHandlerA try process this message
ConcreteHandlerB try process this message
ConcreteHandlerC try process this message
ConcreteHandlerC process this message !!!
*/
