doctrine
-

2.0

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
