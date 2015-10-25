doctrine
-

2.0

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

#### Basic Mapping

````php
<?php
/**
 * @Entity
 * @Table(name="my_persistent_class")
 */
class MyPersistentClass {
    /**
     * @Id @Column(type="integer")
     * @GeneratedValue(strategy="AUTO")
     */
    private $id;
    /** @Column(type="integer") */
    private $intId;
    /** @Column(type="string") */
    private $name;
    /** @Column(name="db_name", type="string") */
    private $dbName;
    /** @Column(name="`number`", type="integer") */
    private $number;
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

Custom Mapping Types:
````php
<?php
namespace My\Project\Types;
use Doctrine\DBAL\Types\Type;
use Doctrine\DBAL\Platforms\AbstractPlatform;

/**
 * My custom datatype.
 */
class MyType extends Type
{
    public function getSqlDeclaration(array $fieldDeclaration, AbstractPlatform $platform)
    {
    // return the SQL used to create your column type.
    // To create a portable column type, use the $platform.
    }

    public function convertToPHPValue($value, AbstractPlatform $platform)
    {
    // This is executed when the value is read from the database.
    // Make your conversions here, optionally using the $platform.
    }

    public function convertToDatabaseValue($value, AbstractPlatform $platform)
    {
    // This is executed when the value is written to the database.
    // Make your conversions here, optionally using the $platform.
    }
}
?>

<?php
use Doctrine\DBAL\Types\Type;
// Register my type
$config->setCustomTypes(array('mytype' => 'My\Project\Types\MyType'));
?>

<?php
class MyPersistentClass
{
    /** @Column(type="mytype") */
    private $field;
}
````

#### Inheritance Mapping

````php
<?php
/** @MappedSuperclass */
class MappedSuperclassBase
{
    /** @Column(type="integer") */
    private $mapped1;
    /** @Column(type="string") */
    private $mapped2;
    /**
     * @OneToOne(targetEntity="MappedSuperclassRelated1")
     * @JoinColumn(name="related1_id", referencedColumnName="id")
     */
    private $mappedRelated1;
    // ... more fields and methods
}

/** @Entity */
class EntitySubClass extends MappedSuperclassBase
{
    /** @Id @Column(type="integer") */
    private $id;
    /** @Column(type="string") */
    private $name;
    // ... more fields and methods
}
````

````sql
CREATE TABLE EntitySubClass (
    mapped1 INTEGER NOT NULL,
    mapped2 TEXT NOT NULL,
    id INTEGER NOT NULL,
    name TEXT NOT NULL,
    related1_id INTEGER DEFAULT NULL,
    PRIMARY KEY(id)
);
````

Single Table Inheritance:

````php
<?php
namespace MyProject\Model;
/**
 * @Entity
 * @SubClasses({"MyProject\Model\Employee"})
 * @InheritanceType("SINGLE_TABLE")
 * @DiscriminatorColumn(name="discr", type="string")
 * @DiscriminatorValue("person")
 */
class Person
{
}

/**
 * @Entity
 * @DiscriminatorValue("employee")
 */
class Employee extends Person
{
}
````

Class Table Inheritance:

````php
<?php
namespace MyProject\Model;
/**
 * @Entity
 * @SubClasses({"MyProject\Model\Employee"})
 * @InheritanceType("JOINED")
 * @DiscriminatorColumn(name="discr", type="string")
 * @DiscriminatorValue("person")
 */

class Person
{
}

/**
 * @Entity
 * @DiscriminatorValue("employee")
 */
class Employee extends Person
{
}
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

#### Batch processing

Bulk Inserts:

````php
<?php
$batchSize = 20;
for ($i = 1; $i <= 10000; ++$i) {
    $user = new CmsUser;
    $user->setStatus('user');
    $user->setUsername('user' . $i);
    $user->setName('Mr.Smith-' . $i);
    $em->persist($user);
    if (($i % $batchSize) == 0) {
        $em->flush();
        $em->clear(); // Detaches all objects from Doctrine!
    }
}
````

Bulk Updates:

````php
<?php
$q = $em->createQuery('update MyProject\Model\Manager m set m.salary = m.salary * 0.9');
$numUpdated = $q->execute();
````

Iterating results:

````php
<?php
$batchSize = 20;
$i = 0;
$q = $em->createQuery('select u from MyProject\Model\User u');
$iterableResult = $q->iterate();
while (($row = $iterableResult->next()) !== false) {
    $user = $row[0];
    $user->increaseCredit();
    $user->calculateNewBonuses();
    if (($i % $batchSize) == 0) {
        $em->flush(); // Executes all updates.
        $em->clear(); // Detaches all objects from Doctrine!
    }
    ++$i;
}
````

Bulk Deletes:

````php
<?php
$q = $em->createQuery('delete from MyProject\Model\Manager m where m.salary > 100000');
$numDeleted = $q->execute();
````

Iterating results:

````php
<?php
$batchSize = 20;
$i = 0;
$q = $em->createQuery('select u from MyProject\Model\User u');
$iterableResult = $q->iterate();
while (($row = $iterableResult->next()) !== false) {
    $em->remove($row[0]);
    if (($i % $batchSize) == 0) {
        $em->flush(); // Executes all deletions.
        $em->clear(); // Detaches all objects from Doctrine!
    }
    ++$i;
}
````

#### DQL (Doctrine Query Language)

````sql
SELECT u FROM MyProject\Model\User u WHERE u.age > 20
````

````php
<?php
// $em instanceof EntityManager
// example1: passing a DQL string
$q = $em->createQuery('select u from MyProject\Model\User u');
// example2: usin setDql
$q = $em->createQuery();
$q->setDql('select u from MyProject\Model\User u');
````
