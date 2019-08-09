VirtualBox
-

````sh
VBoxManage list vms
````

[macOS High Sierra Installation](https://www.youtube.com/watch?v=UUydJTFurMU)
````sh
VBoxManage modifyvm "km2" --cpuidset 00000001 000306a9 04100800 7fbae3ff bfebfbff
VBoxManage setextradata "km2" "VBoxInternal/Devices/efi/0/Config/DmiSystemProduct" "MacBookPro11,3"
VBoxManage setextradata "km2" "VBoxInternal/Devices/efi/0/Config/DmiSystemVersion" "1.0"
VBoxManage setextradata "km2" "VBoxInternal/Devices/efi/0/Config/DmiBoardProduct" "Mac-2BD1B31983FE1663"
VBoxManage setextradata "km2" "VBoxInternal/Devices/smc/0/Config/DeviceKey" "ourhardworkbythesewordsguardedpleasedontsteal(c)AppleComputerInc"
VBoxManage setextradata "km2" "VBoxInternal/Devices/smc/0/Config/GetKeyFromRealSMC" 1

# During install OSX press F8 > Boot Maintenance Manager > Boot from file > last > macOS Install > ... > boot.efi
````
