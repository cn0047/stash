<?php

namespace AppBundle\Controller\Api;

use FOS\RestBundle\Controller\FOSRestController;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;

/**
 * @Route("/api")
 */
class DefaultController extends FOSRestController
{
    /**
     * @Route("/", name="api_homepage")
     */
    public function indexAction(Request $request)
    {
        $response = new JsonResponse(array(
            'desc' => 'api',
        ));
        return $response;
    }
}
