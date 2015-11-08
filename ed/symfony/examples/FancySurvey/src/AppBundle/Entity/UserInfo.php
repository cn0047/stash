<?php

namespace AppBundle\Entity;

use AppBundle\Entity\User;
use Doctrine\ORM\Mapping as ORM;
use Symfony\Bridge\Doctrine\Validator\Constraints\UniqueEntity;
use Symfony\Component\Validator\Constraints as Assert;

/**
 * @ORM\Entity
 */
class UserInfo
{
    /**
     * @ORM\Id
     * @ORM\Column(type="integer")
     * @ORM\GeneratedValue(strategy="AUTO")
     */
    private $id;

    /**
     * @ORM\OneToOne(targetEntity="User")
     * @ORM\JoinColumn(name="user_id", referencedColumnName="id")
     **/
    private $user;

    /**
     * @ORM\Column(type="string", length=50)
     * @Assert\NotBlank()
     */
    private $iceCream;

    /**
     * @ORM\Column(type="string", length=50)
     * @Assert\NotBlank()
     */
    private $superHero;

    /**
     * @ORM\Column(type="string", length=50)
     * @Assert\NotBlank()
     */
    private $movieStar;

    /**
     * @ORM\Column(type="string", length=50)
     * @Assert\NotBlank()
     */
    private $worldEnd;

    /**
     * @ORM\Column(type="string", length=50)
     * @Assert\NotBlank()
     */
    private $whoWinSuperBowl;

    public function setUser(User $user)
    {
        $this->user = $user;
        return $this;
    }

    public function getUser()
    {
        return $this->user;
    }

    public function getIceCream()
    {
        return $this->iceCream;
    }

    public function setIceCream($iceCream)
    {
        $this->iceCream = $iceCream;
    }

    public function getSuperHero()
    {
        return $this->superHero;
    }

    public function setSuperHero($superHero)
    {
        $this->superHero = $superHero;
    }

    public function getMovieStar()
    {
        return $this->movieStar;
    }

    public function setMovieStar($movieStar)
    {
        $this->movieStar = $movieStar;
    }

    public function setWorldEnd($worldEnd)
    {
        $this->worldEnd = $worldEnd;
    }

    public function getWorldEnd()
    {
        return $this->worldEnd;
    }

    public function getWhoWinSuperBowl()
    {
        return $this->whoWinSuperBowl;
    }

    public function setWhoWinSuperBowl($whoWinSuperBowl)
    {
        $this->whoWinSuperBowl = $whoWinSuperBowl;
    }
}