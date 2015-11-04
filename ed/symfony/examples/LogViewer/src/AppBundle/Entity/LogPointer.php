<?php

namespace AppBundle\Entity;

use Doctrine\ORM\Mapping as ORM;
use Symfony\Bridge\Doctrine\Validator\Constraints\UniqueEntity;
use Symfony\Component\Validator\Constraints as Assert;

/**
 * @ORM\Entity
 * @ORM\Table(name="log_pointer")
 * @UniqueEntity(fields="file", message="File already taken.")
 */
class LogPointer
{
    /**
     * @ORM\Id
     * @ORM\Column(type="integer")
     * @ORM\GeneratedValue(strategy="AUTO")
     */
    private $id;

    /**
     * @ORM\Column(type="string", length=255, options={"default" = ""}, unique=true)
     * @Assert\NotBlank()
     * @Assert\Length(max=255)
     */
    private $file = '';

    /**
     * @ORM\Column(type="integer", options={"default" = 0})
     * @Assert\NotBlank()
     * @Assert\Type(type="integer")
     */
    private $pointer = 0;

    public function setFile($file)
    {
        $this->file = $file;
    }

    public function getPointer()
    {
        return $this->pointer;
    }

    public function setPointer($pointer)
    {
        $this->pointer = $pointer;
    }
}
