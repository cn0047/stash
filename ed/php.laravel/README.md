Laravel
-
5.3

````
composer global require "laravel/installer" 
laravel new blogProject

composer require barryvdh/laravel-ide-helper
````
````
php artisan serve
````

* Service Container - for managing class dependencies.
* Service Providers - central place of application bootstrapping
  (registering service container bindings, event listeners, middleware, and even routes).
* Facades - "static" interface to classes that are available in the application\'s service container.
* Contracts - a set of interfaces that define the core services provided by the framework.
* Response Macros - custom response that you can re-use.
* Collections - provides a fluent, convenient wrapper for working with arrays of data.
* Collections - multi-result sets returned by Eloquen.

## Service Container

````
# get service from container
$h = $this->app->make('App\Services\Helper');
````

## Controller

````
# Create CRUD controller.
php artisan make:controller PhotoController --resource

# Route Caching
php artisan route:cache
php artisan route:clear
````

## DB

````php
$shopper = DB::table('virtual_cards')
    ->where('card_number','=', $post_data['card_number'])
    ->first()
;
$r = DB::table('virtual_cards')
    ->select('shoppers.id', 'shoppers.active_status', 'shoppers.is_verified')
    ->leftJoin('shoppers','virtual_cards.shopper_id','=','shoppers.id')
    ->where('virtual_cards.card_number','=', $post_data['card_number'])
    ->first()
;
````

## Blade

````blade
@section - section of content
@yield - display the contents of a given section
@include('sub-view')

{{ trans('messages.welcome') }}
@lang('messages.welcome')
````

## Misc

````
Illuminate\Support\Facades\Hash::make('str');
Illuminate\Support\Facades\Hash::check('plain-text', $hash);
````

## Migrations

````
php artisan migrate # running migrations
````

DEBT:
https://laravel.com/docs/5.4/lifecycle
https://laravel.com/docs/5.4/providers
https://laravel.com/docs/5.4/container
https://laravel.com/docs/5.4/facades
https://laravel.com/docs/5.4/contracts
https://laravel.com/docs/5.4/passport
https://laravel.com/docs/5.4/broadcasting
