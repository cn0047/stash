
read -p "Are you sure [y/N]" -n 1 -r
echo 

if [[ $REPLY =~ ^[Yy]$ ]]
then
    echo 
    echo "OK"
fi
