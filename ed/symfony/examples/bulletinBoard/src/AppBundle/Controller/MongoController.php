<?php

namespace AppBundle\Controller;

use AppBundle\Document\Product;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;

class MongoController extends Controller
{
    /**
    * @Route("/mongo_create", name="mongo_create")
    */
    public function createAction()
    {
        $product = new Product();
        $product->setName('A Foo Bar');
        $product->setPrice('19.99');

        $dm = $this->get('doctrine_mongodb')->getManager();
        $dm->persist($product);
        $dm->flush();

        return new Response('Created product id '.$product->getId());
    }

    /**
     * @Route("/mongo_show/{id}", name="mongo_show")
     */
    public function showAction($id)
    {
        $product = $this->get('doctrine_mongodb')
            ->getRepository('AppBundle:Product')
            ->find($id);
        if (!$product) {
            throw $this->createNotFoundException('No product found for id '.$id);
        }
        var_dump($product);
    }
}
