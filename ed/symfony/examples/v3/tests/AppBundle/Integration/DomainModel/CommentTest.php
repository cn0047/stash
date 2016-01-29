<?php

namespace Tests\AppBundle\Integration\DomainModel;

use AppBundle\DomainModel\Comment;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;
use Doctrine\Common\DataFixtures\Loader;
use Tests\AppBundle\Fixtures\Entity\Comment as CommentFixture;
use Doctrine\Common\DataFixtures\Executor\ORMExecutor;
use Doctrine\Common\DataFixtures\Purger\ORMPurger;

class CommentTest extends KernelTestCase
{
    /**
     * @var \Doctrine\ORM\EntityManager
     */
    private $em;

    /**
     * {@inheritDoc}
     */
    protected function setUp()
    {
        self::bootKernel();

        $this->em = static::$kernel->getContainer()
            ->get('doctrine')
            ->getManager()
        ;

        $loader = new Loader();
        $loader->addFixture(new CommentFixture());
        $purger = new ORMPurger();
        $executor = new ORMExecutor($this->em, $purger);
        $executor->execute($loader->getFixtures());
    }

    /**
     * {@inheritDoc}
     */
    protected function tearDown()
    {
        parent::tearDown();

        $this->em->close();
    }

    public function testGet()
    {
        $c = static::$kernel->getContainer()->get('comment');
        $comments = $c->get();
        $this->assertSame('Hello from db fixture generated on fly!', $comments[0]['message']);
    }
}
