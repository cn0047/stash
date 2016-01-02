<?php

namespace AppBundle\Controller\Api;

use FOS\RestBundle\Controller\Annotations\Route;
use FOS\RestBundle\Controller\Annotations\RouteResource;
use FOS\RestBundle\Controller\FOSRestController;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;

/**
 * @Route("/comments")
 */
class CommentsController extends FOSRestController
{
    /**
     * @Route("/", name="post_comments")
     * @Method("POST")
     */
    public function postCommentsAction(Request $request)
    {
        // curl -X POST http://127.0.0.1:8000/comments/ -d message=msg -d user=007 -d route=/home
        var_export($request->request->get('message'));
        return new JsonResponse([]);
    }
}
