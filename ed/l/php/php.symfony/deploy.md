Deploy
-

````sh
php bin/symfony_requirements

# check all parameters in parameters.yml.dist

composer install --no-dev --optimize-autoloader

php bin/console cache:clear --env=prod --no-debug --no-warmup
php bin/console cache:warmup --env=prod

php bin/console assetic:dump --env=prod --no-debug
````
