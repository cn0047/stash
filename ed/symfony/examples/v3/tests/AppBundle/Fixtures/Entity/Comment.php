<?php

namespace Tests\AppBundle\Fixtures\Entity;

use Doctrine\Common\DataFixtures\FixtureInterface;
use Doctrine\Common\Persistence\ObjectManager;
use AppBundle\Entity;

class Comment implements FixtureInterface
{
    // php bin/console doctrine:fixtures:load --fixtures tests/AppBundle/Fixtures/ -e test
    public function load(ObjectManager $manager)
    {
        $c = new Entity\Comment();
        $c->setRoute('/');
        $c->setMessage('Hello from db fixture generated on fly!');
        $c->setUser('phpunit');
        $manager->persist($c);
        $manager->flush();
    }
}
