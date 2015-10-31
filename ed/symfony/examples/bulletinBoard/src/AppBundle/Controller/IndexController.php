<?php

namespace AppBundle\Controller;

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use APY\DataGridBundle\Grid\Source\Entity;
use APY\DataGridBundle\Grid\Source\Vector;

class IndexController extends Controller
{
    public function indexAction()
    {
        return $this->render('default/index.html.twig');
    }

    /**
    * @Route("/myArrayGrid", name="myArrayGrid")
    */
    public function myArrayGridAction()
    {
        $data = [
            ['id' => 1, 'title' => 'book1', 'publication' => '2012-04-06'],
            ['id' => 2, 'title' => 'book2', 'publication' => 'Apr. 6, 2012'],
            ['id' => 3, 'title' => 'book3', 'publication' => '2013-04-06'],
            ['id' => 4, 'title' => 'book5', 'publication' => '2014-04-06'],
            ['id' => 5, 'title' => 'book6', 'publication' => '2015-04-06'],
            ['id' => 6, 'title' => 'book6', 'publication' => '2015-04-06'],
            ['id' => 7, 'title' => 'book7', 'publication' => '2015-04-07'],
            ['id' => 8, 'title' => 'book8', 'publication' => '2015-04-08'],
            ['id' => 9, 'title' => 'book9', 'publication' => '2015-04-09'],
            ['id' => 10, 'title' => 'book10', 'publication' => '2015-04-10'],
            ['id' => 11, 'title' => 'book11', 'publication' => '2015-04-11'],
            ['id' => 12, 'title' => 'book12', 'publication' => '2015-04-12'],
            ['id' => 13, 'title' => 'book13', 'publication' => '2015-04-13'],
        ];
        $source = new Vector($data);
        $grid = $this->get('grid');
        $grid->setSource($source);
        $grid->setLimits([5, 10, 15]);
        return $grid->getGridResponse('AppBundle:default:index.html.twig');
    }

    /**
    * @Route("/myDbGrid", name="myDbGrid")
    */
    public function myDbGridAction()
    {
        $source = new Entity('AppBundle:GridEntity');
        $grid = $this->get('grid');
        $grid->setSource($source);
        $grid->setLimits([5, 10, 15]);
        return $grid->getGridResponse('AppBundle:default:index.html.twig');
    }
}
