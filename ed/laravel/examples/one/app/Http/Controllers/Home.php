<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Requests;

class Home extends Controller
{
    public function index($code)
    {
        echo $code;
    }

    public function home()
    {
        return view('front.home', ['time' => time()]);
    }
}
