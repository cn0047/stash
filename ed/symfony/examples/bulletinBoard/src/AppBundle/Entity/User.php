<?php

namespace AppBundle\Entity;

use Symfony\Component\Security\Core\User\UserInterface;
use Symfony\Component\Validator\Constraints as Assert;

/**
 * @Assert\GroupSequence({"User", "Strict"})
 */
class User // implements UserInterface
{
    /**
    * @Assert\Email(groups={"registration"})
    */
    private $email;

    /**
     * @Assert\NotBlank
     */
    private $username;

    /**
    * @Assert\NotBlank(groups={"registration"})
    * @Assert\Length(min=7, groups={"registration"})
    */
    private $password;

    /**
    * @Assert\Length(min=2)
    */
    private $city;

    /**
     * @Assert\CardScheme(
     * schemes={"VISA"},
     * groups={"Premium"},
     * )
     */
    private $creditCard;

    /**
     * @Assert\True(message="The password cannot match your username", groups={"Strict"})
     */
    public function isPasswordLegal()
    {
      return ($this->username !== $this->password);
    }

    public function getGroupSequence()
    {
        $groups = array('User');
        if ($this->isPremium()) {
           $groups[] = 'Premium';
        }
        return $groups;
    }

    public function addEmailAction($email)
    {
        $emailConstraint = new Assert\Email();
        // all constraint "options" can be set this way
        $emailConstraint->message = 'Invalid email address';
        // use the validator to validate the value
        $errorList = $this->get('validator')->validate(
            $email,
            $emailConstraint
        );
        if (0 === count($errorList)) {
        } else {
            $errorMessage = $errorList[0]->getMessage();
        }
    }
}