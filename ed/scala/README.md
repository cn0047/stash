Scala
-
2.12.3

Scala - is static typed programming language,
fuses both imperative-programming and functional-programming styles.

````sh
scala hw.scala

sbt scalaVersion

# in proj root
sbt
run
````

Basic types:

* Byte
* Short
* Int
* Char
* Float
* Double
* Boolean
* String


Traversable:

Iterable
  Seq
  Set
  Map
  IndexedSeq
  Vector
  ResizableArray
  GenericArray
  LinearSeq
  MutableList
  List
  Stream
  Buffer
  ListBuffer
  ArrayBuffer
  SortedSet
  TreeSet
  HashSet (mutable)
  LinkedHashSet
  HashSet (immutable)
  BitSet
  EmptySet,
  Set1, Set2, Set3, Set4
  SortedMap
  TreeMap
  HashMap (mutable)
  LinkedHashMap (mutable)
  HashMap (immutable)
  EmptyMap,
  Map1, Map2, Map3, Map4

````
// keywords
4 to 6 
abs
capitalize
drop
isInfinity
max
min
round

boolean nameHasUpperCase = false;
for (int i = 0; i < name.length(); ++i) {
    if (Character.isUpperCase(name.charAt(i))) {
        nameHasUpperCase = true;
        break;
    }
}

// switch
receive {
    case Msg1 => ... // handle Msg1
    case Msg2 => ... // handle Msg2
    // ...
}
````

Loops:

````
args.foreach(arg => println(arg))

args.foreach((arg: String) => println(arg))

for (i <-0.to(2)) print(greetStrings.apply(i))
for (z <- arr.zipWithIndex if z._2 % 2 == 1) yield z._1
{for{a <- arr; if a%2 !=0 }yield a}.sum

while (a != 0) {
    val temp = a a = b % a b = temp
}

do {
    line = readLine() println("Read: "+ line)
} while (line != "")

list.map(_.grouped(2).map(_.reverse).mkString("")).map(println)
````

Functions:

````
def f () = println("Hello World")

def f(delim:Int,arr:List[Int]):List[Int] = for{a <- arr if a < delim } yield a

arr.flatMap(e => List.fill(num)(e))

def factorial(x: BigInt): BigInt = if (x == 0) 1 else x * factorial(x -1)

def f (args: Number) {
    for (a <- 1 to n) {
        println("Hello World");
    }
}
````

Data structures:

````
val buf = collection.mutable.ArrayBuffer.empty[Int]

val input = io.Source.stdin.getLines.drop(1).toList;

val x: HashMap[Int, String] = new HashMap[Int, String]()

import scala.collection.mutable.ArrayBuffer
val ab = ArrayBuffer[String]()
ab += "hello"
ab += "world"
ab += "!"
ab.toArray
println(ab.mkString(" "))

import scala.collection.mutable.MutableList
val x = MutableList(1, 2, 3, 4, 5)
x += 6

val twoThree = List(2, 3)
val oneTwoThree = 1 :: twoThree
println(oneTwoThree)

var capital = Map("US" -> "Washington", "France" -> "Paris")
capital += ("Japan" -> "Tokyo")
println(capital("France"))

var jetSet = Set("Boeing", "Airbus")
jetSet += "Lear"
println(jetSet.contains("Cessna"))

var jetSet = scala.collection.mutable.SortedSet[String]()
jetSet += "Lear"
println(jetSet)

import scala.collection.mutable.Map
val treasureMap = Map[Int, String]() 
reasureMap += (1 -> "Go to island.")
treasureMap += (2 -> "Find big X on ground.")
treasureMap += (3 -> "Dig.")
println(treasureMap(2))
````

Class:

````
class ChecksumAccumulator {
    private var sum = 0
    def add(b: Byte): Unit = { sum += b }
    def checksum(): Int = {
        return ~(sum & 0xFF) + 1
    }
}

class MyClass {
    private int index;
    private String name;
    public MyClass(int index, String name) {
        this.index = index;
        this.name = name;
    }
}
````
