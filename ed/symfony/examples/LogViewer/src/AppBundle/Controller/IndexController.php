<?php

namespace AppBundle\Controller;

use APY\DataGridBundle\Grid\Source\Entity;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;

/**
 * Main controller.
 *
 * This controller provides access to data grid.
 */
class IndexController extends Controller
{
    /**
     * @Route("/", name="index")
     */
    public function indexAction(Request $request)
    {
        $isAdmin = $this
            ->get('security.authorization_checker')
            ->isGranted('ROLE_ADMIN')
        ;
        $source = new Entity('AppBundle:Log');
        $grid = $this->get('grid');
        $grid->setSource($source);
        $grid->setLimits([10, 25, 50, 100]);
        if (!$isAdmin) {
            $grid->hideColumns('owner');
            $grid->setDefaultFilters([
                'owner' => ['operator' => 'eq', 'from' => 'user'],
            ]);
        }
        return $grid->getGridResponse('AppBundle:Index:Index.html.twig');
    }
}
