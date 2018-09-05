File Systems
-

* overlayfs
* rootfs
* ext4 (best for backend)

````
mount -t overlay overlay -o lowerdir=/lower1:/lower2:/lower3,upperdir=/upper,workdir=/work /merged
````
