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
    // Dumps user.
    // var_export(\Illuminate\Support\Facades\Auth::user());die;
    return view('welcome');
});
Route::get('blade', function () {
    return view('layouts.child');
});

Route::group(['middleware' => 'age'], function () {
    Route::get('/get200', function () {
        echo 200;
    });
});

Route::get('/get/{code}', 'Home@index');
Route::get('/home', 'Home@home');
Route::get('/di', 'Home@di');
Route::get('/db', 'Home@db');
Route::get('/auth', 'Home@auth');
Route::get('/conf', 'Home@conf');
Route::get('/flash', 'Home@flash');
Route::get('/validation', 'Home@validation');
Route::get('/localization', 'Home@localization');
Route::resource('photos', 'PhotoController');

Route::auth();
Route::get('/home', 'HomeController@index');
