Doctrine
-

2.0

Object relational mapper (ORM) for PHP that sits on top of a powerful `Database Abstraction Layer (DBAL)`.
One of its key features is the option to write database queries
in a proprietary object oriented SQL dialect called `Doctrine Query Language (DQL)`.
`DBAL` - database abstraction layer
with many features for database schema introspection, schema management and PDO abstraction.

Doctrine use `ODM` (Object Document Mapper) for MongoDB.

`EntityManager` - central access point (facade) to ORM functionality (UnitOfWork, Query Language, Repository API).

`UnitOfWork` - responsible for tracking changes to objects
during "object-level" transaction and for writing out changes to the database.

A new UnitOfWork is implicitly started when an EntityManager is initially created 
or after `EntityManager#flush()`.

<br>`flush` - write operations into DB,
<br>`persist`, `remove` - only notify the UnitOfWork to perform these operations during flush.
<br>not calling `flush` - all changes will lost.

For relationships between entities, you don't have to have physical foreign key in db.

@HasLifecycleCallbacks - required for:

* @PostLoad
* @PrePersist
* @PostPersist
* @PreRemove
* @PostRemove
* @PreUpdate
* @PostUpdate

@GeneratedValue - strategy for identifier generation annotated by @Id. Available values:

* AUTO (default)
* SEQUENCE - Oracle, PostgreSql and SQL Anywhere.
* TABLE - Use a separate table for ID generation (Not implemented).
* IDENTITY - MySQL (AUTO_INCREMENT), MSSQL (IDENTITY), PostgreSQL (SERIAL.
* UUID - Built-in Universally Unique Identifier generator.
* CUSTOM
* NONE

#### Installing

````
svn co http://svn.doctrine-project.org/trunk doctrine
cd doctrine

"require": {
    "doctrine/orm": "2.4.*"
}
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

#### Working with objects

````php
<?php
$uowSize = $em->getUnitOfWork()->size();
?>

<?php
$user = new User;
$user->setName('Mr.Right');
$em->persist($user);
$em->flush();
// If $user had a generated identifier, it would now be available.
?>

<?php
$em->remove($user);
$em->flush();
?>

<?php
// Detaching entities
$detachedEntity = $em->detach($entity);
?>

<?php
// Merging entities
$detachedEntity = unserialize($serializedEntity); // some detached entity
$entity = $em->merge($detachedEntity);
// $entity now refers to the fully managed copy returned by the merge operation.
// The EntityManager $em now manages the persistence of $entity as usual.
?>
````
