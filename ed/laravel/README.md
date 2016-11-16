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

## Controller

````
# Create CRUD controller.
php artisan make:controller PhotoController --resource

# Route Caching
php artisan route:cache
php artisan route:clear
````

## Blade

````blade
@section - section of content
@yield - display the contents of a given section

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

https://laravel.com/docs/5.3/passport
