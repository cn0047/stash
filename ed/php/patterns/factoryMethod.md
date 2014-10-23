Factory method
-

Group 1.

````php
<?php
class Twitter {
    public function share()
    {
        return $this->tweet();
    }
}

class Facebook {
    public function share()
    {
        return $this->post();
    }
}

class Factory {
    public static function create($name)
    {
        if (class_exists($name)) {
            return new $name();
        }
    }
}

$twitter = Factory::create('Twitter');
$facebook = Factory::create('Facebook');
````
