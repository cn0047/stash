
# whatever
java  ed/l/java/examples/whatever/hw.java
java  ed/l/java/examples/whatever/static1.java
java  ed/l/java/examples/whatever/cliArgs.java
java  ed/l/java/examples/whatever/str.java
java  ed/l/java/examples/whatever/collections.java
java  ed/l/java/examples/whatever/class1.java
java  ed/l/java/examples/whatever/class2.java
java  ed/l/java/examples/whatever/map1.java
java  ed/l/java/examples/whatever/cryptoAES.java
java  ed/l/java/examples/whatever/date.java

# javac whatever
cd ed/l/java/examples/whatever
javac hw.java
java  hw
# json
jar=$HOME/.m2/repository/com/googlecode/json-simple/json-simple/1.1.1/json-simple-1.1.1.jar
# javac -Xlint:unchecked -cp $jar json.java
# java  json

# pkg
cd ed/l/java/examples/pkg
javac src/com.x/*.java
# java -classpath src/com.x com.x.Main

# jar
cd ed/l/java/examples/jar
javac Main.java
jar cmfv MANIFEST.MF /tmp/x.jar *.class
java -jar /tmp/x.jar

# maven
cd ed/l/java/examples/maven
mvn clean
mvn compile
mvn package
# java -jar target/xmaven.jar

#### Java Spring

# demo
# cd ed/l/java/java.spring/examples/demo
# ./mvnw spring-boot:run
# curl http://localhost:8080/hello
