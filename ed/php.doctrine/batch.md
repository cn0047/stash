Batch
-

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
