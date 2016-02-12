<?php

namespace AppBundle\VO;

use Symfony\Component\Validator\Constraints as Assert;

class UserForValidation
{
    /**
     * @Assert\NotBlank(groups={"g1"})
     */
    private $username;

    /**
     * @Assert\NotBlank(groups={"g2"})
     */
    private $password;

    public function getRoles()
    {
        return array('ROLE_USER');
    }

    public function __toString()
    {
        return __CLASS__;
    }
}
