<?php

namespace AppBundle\Twig;

use AppBundle\AppBundle;
use AppBundle\DomainModel\Comment;
use Symfony\Component\DependencyInjection\ContainerInterface;

class CommentsExtension extends \Twig_Extension
{
    private $container;

    /**
     * Constructor.
     *
     * @param ContainerInterface $container A container.
     */
    public function __construct(ContainerInterface $container)
    {
        $this->container = $container;
    }

    public function getFunctions()
    {
        return array(
            new \Twig_SimpleFunction('comments', array($this, 'commentsFunction')),
        );
    }

    public function commentsFunction()
    {
        $c = new Comment($this->container);
        $comments = $c->get();
        $comments = [
            ['route' => '/', 'message' => 'hi', 'user' => 'unknown'],
            ['route' => '/index', 'message' => 'are you there?', 'user' => 'nobody'],
        ];
        $r = $this->container->get('templating')->render(
            'AppBundle:default:comments.html.twig',
            ['comments' => $comments]
        );
        return $r;
    }

    public function getName()
    {
        return 'comments_extension';
    }
}
