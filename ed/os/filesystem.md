File System
-

* overlayfs
* rootfs
* ext4 (best for backend)
* procfs

````sh
mount -t overlay overlay -o lowerdir=/lower1:/lower2:/lower3,upperdir=/upper,workdir=/work /merged
````
