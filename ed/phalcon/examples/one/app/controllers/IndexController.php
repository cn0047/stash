<?php

use Phalcon\Mvc\Controller;

class IndexController extends Controller
{

    public function indexAction()
    {
        echo "<h1>Hello!</h1>";
        echo $this->tag->linkTo("signup", "Sign Up Here!");
    }
}
