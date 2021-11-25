# range 1
for i in $(seq 1 5); do
    echo $i
done

# range 2
for i in `seq 1 2 11`; do
    echo $i;
done

# read line by line from file
for i in $( cat file.txt ); do
    echo item: $i;
done

for i in $( ls ); do
    echo item: $i;
done
