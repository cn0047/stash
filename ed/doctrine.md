doctrine
-

2.0

#### Installing
````
svn co http://svn.doctrine-project.org/trunk doctrine
cd doctrine
````

#### Configuration
````php
<?php
// test.php
require '/path/to/lib/Doctrine/Common/ClassLoader.php';
$classLoader = new \Doctrine\Common\ClassLoader();
````

````php
<?php
$classLoader->setBasePath($prefix, $basePath);
$classLoader->setBasePath('Doctrine', '/usr/local/phplibs/doctrine/lib');
````

````php
<?php
use Doctrine\ORM\EntityManager,
Doctrine\ORM\Configuration;
$config = new Configuration;
$cache = new \Doctrine\Common\Cache\ApcCache;
$config->setMetadataCacheImpl($cache);
$config->setQueryCacheImpl($cache);
$connectionOptions = array(
    'driver' => 'pdo_sqlite',
    'path' => 'database.sqlite'
);
$em = EntityManager::create($connectionOptions, $config);
````

````
/** @ChangeTrackingPolicy("DEFERRED_IMPLICIT") */
/** @ChangeTrackingPolicy("DEFERRED_EXPLICIT") */
````

````php
<?php
use Doctrine\Common\NotifyPropertyChanged,
    Doctrine\Common\PropertyChangedListener;
/**
 * @Entity
 * @ChangeTrackingPolicy("NOTIFY")
 */
class MyEntity implements NotifyPropertyChanged
{
    private $_listeners = array();
    public function addPropertyChangedListener(PropertyChangedListener $listener)
    {
        $this->_listeners[] = $listener;
    }
    protected function _onPropertyChanged($propName, $oldValue, $newValue)
    {
        if ($this->_listeners) {
            foreach ($this->_listeners as $listener) {
                $listener->propertyChanged($this, $propName, $oldValue, $newValue);
            }
        }
    }
}
?>
<?php
public function setData($data)
{
    if ($data != $this->data) {
        $this->_onPropertyChanged('data', $this->data, $data);
        $this->data = $data;
    }
}
````

````php
<?php
$q = $em->createQuery("select u.id, u.name from MyApp\Domain\User u");
$q->setHint(Query::HINT_FORCE_PARTIAL_LOAD, true);
````

#### Basic Mapping
````php
<?php
/**
 * @Entity
 * @Table(name="my_persistent_class")
 */
class MyPersistentClass {
    /** @Column(type="integer") */
    private $id;
    /** @Column(type="string") */
    private $name;
    /** @Column(name="db_name", type="string") */
    private $dbName;
}
````

Built-in mapping types:

* string: Type that maps an SQL VARCHAR to a PHP string.
* integer: Type that maps an SQL INT to a PHP integer.
* smallint: Type that maps a database SMALLINT to a PHP integer.
* bigint: Type that maps a database BIGINT to a PHP string.
* boolean: Type that maps an SQL boolean to a PHP boolean.
* decimal: Type that maps an SQL DECIMAL to a PHP double.
* date: Type that maps an SQL DATETIME to a PHP DateTime object.
* time: Type that maps an SQL TIME to a PHP DateTime object.
* datetime: Type that maps an SQL DATETIME/TIMESTAMP to a PHP DateTime object.
* text: Type that maps an SQL CLOB to a PHP string.

Column annotation attributes:

* type: The mapping type to use for the column.
* name: (optional, defaults to field name) The name of the column in the database.
* length: (optional, default 255) The length of the column in the database. (Applies only if a string-valued column is used).
* unique: (optional, default FALSE) Whether the column is a unique key.
* nullable: (optional, default FALSE) Whether the database column is nullable.
* precision: (optional, default 0) The precision for a decimal (exact numeric) column. (Applies only if a decimal column is used.)
* scale: (optional, default 0) The scale for a decimal (exact numeric) column. (Applies only if a decimal column is used.)

Page:20
