The Bee Game
-

To play this game just run in shell: `php index.php`.

During game you'll see game progress like:
````
Bee\Drone: 50 50 50 50 50 50 50 50
Bee\Worker: 75 75 75 75 75
Bee\Queen: 100
Do you wish hit bee (y) or exit game (n)?
````
or something like this:
````
Bee\Drone: 2
Bee\Worker: 15
Game over. Start new game (y/n)?
````

### Specification:

#### Bees:

There are three types of bees in this game:

* Queen Bee:

    * The Queen Bee has a lifespan of 100 Hit Points.
    * When the Queen Bee is hit, 8 Hit Points are deducted from her lifespan.
    * If/When the Queen Bee has run out of Hit Points, All remaining alive Bees automatically run out of hit points.
    * There is only 1 Queen Bee.

* Worker Bee:
    * Worker Bees have a lifespan of 75 Hit Points.
    * When a Worker Bee is hit, 10 Hit Points are deducted from his lifespan.
    * There are 5 Worker Bees.

* Drone Bee:
    * Drone Bees have a lifespan of 50 Hit Points.
    * When a Drone Bee is hit, 12 Hit Points are deducted from his lifespan.
    * There are 8 Drone Bees.

#### Gameplay:

To play, there must be a button that enables a user to "hit" a random bee.
The selection of a bee must be random.
When the bees are all dead, the game must be able to reset itself with full life bees for another round.
