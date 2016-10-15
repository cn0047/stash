<?php

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It's a breeze. Simply tell Laravel the URIs it should respond to
| and give it the controller to call when that URI is requested.
|
*/

Route::get('/', function () {
    return view('welcome');
});
Route::get('blade', function () {
    return view('layouts.child');
});

Route::get('/get200', function () {
    echo 200;
});

Route::get('/get/{code}', 'Home@index');
Route::get('/home', 'Home@home');
Route::get('/conf', 'Home@conf');
Route::resource('photos', 'PhotoController');

Route::auth();
Route::get('/home', 'HomeController@index');
