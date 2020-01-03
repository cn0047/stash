<?php

namespace Screecher\Entity;

use Symfony\Component\Validator\Constraints as Assert;

/**
 * @Entity
 * @Table(name="maintainer")
 */
class Maintainer
{
    /**
     * @Id
     * @Column(type="integer")
     * @GeneratedValue(strategy="AUTO")
     */
    private $id;

    /**
     * @Column(name="api_id", type="integer")
     */
    private $apiId;

    /**
     * @Column(type="string", length=100)
     */
    private $email;

    static public function loadValidatorMetadata($metadata)
    {
        $metadata->addPropertyConstraint('apiId', new Assert\Regex('/^\d+$/'));
        $metadata->addPropertyConstraint('email', new Assert\Email());
    }

    public function getId()
    {
        return $this->id;
    }

    public function getApiId()
    {
        return $this->apiId;
    }

    public function setApiId($apiId)
    {
        $this->apiId = $apiId;
    }

    public function getEmail()
    {
        return $this->email;
    }

    public function setEmail($email)
    {
        $this->email = $email;
    }

    public function setProperties(array $properties)
    {
        foreach ($properties as $name => $value) {
            $this->$name = $value;
        }
    }
}
