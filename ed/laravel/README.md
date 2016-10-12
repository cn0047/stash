Laravel
-
5.3

````
composer global require "laravel/installer" 
laravel new blogProject
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

https://laravel.com/docs/5.3/passport

https://laravel.com/docs/5.3/broadcasting
