Composer
-

````sh
curl -sS https://getcomposer.org/installer | php
sudo mv composer.phar /usr/local/bin/composer
````

````sh
composer clear-cache
composer config github-oauth.github.com 2c58b0e2f9a6c0f7e8f1953310f676e249ff62d2

composer require --dev phpspec/phpspec
````

#### Versions

````js
"require": {
    "vendor/package": "1.3.2", // exactly 1.3.2
    // >, <, >=, <= | specify upper / lower bounds
    "vendor/package": ">=1.3.2", // anything above or equal to 1.3.2
    "vendor/package": "<1.3.2", // anything below 1.3.2
    // * | wildcard
    "vendor/package": "1.3.*", // >=1.3.0 <1.4.0
    // ~ | allows last digit specified to go up
    "vendor/package": "~1.3.2", // >=1.3.2 <1.4.0
    "vendor/package": "~1.3", // >=1.3.0 <2.0.0
    // ^ | doesn't allow breaking changes (major version fixed - following semver)
    "vendor/package": "^1.3.2", // >=1.3.2 <2.0.0
    "vendor/package": "^0.3.2", // >=0.3.2 <0.4.0 // except if major version is 0
    "vendor/package": "^0.3", // >=0.3.0 <0.4.0
}
````
