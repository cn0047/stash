<?php

class Url
{
    public static function getProdId()
    {
        return 'Category1';
    }
}

class Product
{
    public static $name = 'Car';

    public static function instance($prodId)
    {
        if ($prodId === 'Category1') {
            return new Category1();
        }
    }
}

class Category1 extends Product
{
    public $model = 'DB9';

    public function info()
    {
        return 'Aston Martin DB9 v12';
    }
}

class Page
{
    public $html;

    /**
     * DI.
     */
    public function createPage(Product $product)
    {
        $this->html = $product->info() . PHP_EOL;
    }

    public function showPage()
    {
        echo $this->html;
    }
}

$page = new Page();
$page->createPage(Product::instance(Url::getProdId()));
$page->showPage();
var_export($page);

/*
Aston Martin DB9 v12
Page::__set_state(array(
   'html' => 'Aston Martin DB9 v12
',
))
*/
