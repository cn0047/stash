<?php

namespace AppBundle\Entity;

use APY\DataGridBundle\Grid\Mapping as GRID;
use Doctrine\ORM\Mapping as ORM;
use Symfony\Component\Validator\Constraints as Assert;

/**
 * Db analog for line record from log file.
 *
 * @ORM\Entity(repositoryClass="AppBundle\Entity\LogRepository")
 * @ORM\Table(name="log")
 * @GRID\Source(columns="id, owner, host, user, dateTime, firstRequestLine, status, size, userAgent")
 */
class Log
{
    /**
     * @ORM\Id
     * @ORM\Column(type="integer")
     * @ORM\GeneratedValue(strategy="AUTO")
     */
    private $id;

    /**
     * @ORM\Column(type="string", length=50, options={"default" = ""})
     * @Assert\NotBlank()
     * @Assert\Length(max=50)
     */
    private $owner;

    /**
     * @ORM\Column(type="string", length=50, options={"default" = ""})
     * @Assert\Length(max=50)
     */
    private $host;

    /**
     * @ORM\Column(type="string", length=50, options={"default" = ""})
     * @Assert\Length(max=50)
     */
    private $user;

    /**
     * @ORM\Column(type="datetime", options={"default" = "0000-00-00 00:00:00"})
     * @Assert\Date()
     */
    private $dateTime;

    /**
     * @ORM\Column(type="string", length=100, options={"default" = ""})
     * @Assert\Length(max=100)
     */
    private $firstRequestLine;

    /**
     * @ORM\Column(type="integer")
     * @Assert\Type(type="integer")
     */
    private $status;

    /**
     * @ORM\Column(type="integer")
     * @Assert\Type(type="integer")
     */
    private $size;

    /**
     * @ORM\Column(type="string", length=250, options={"default" = ""})
     * @Assert\Length(max=250)
     */
    private $userAgent;

    /**
     * Convenient way to set all properties in easiest way.
     * @param array $args Parameters that will setted as entity properties.
     */
    public function init(array $args)
    {
        foreach ($args as $field => $value) {
            $this->$field = $value;
        }
    }
}
