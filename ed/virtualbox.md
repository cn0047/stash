VirtualBox
-

````
VBoxManage list vms

# VBoxManage modifyvm kmac  --cpuidset 00000001 000306a9 00020800 80000201 178bfbff
# VBoxManage setextradata "kmac" "VBoxInternal/Devices/efi/0/Config/DmiSystemProduct" "iMac11,3"
# VBoxManage setextradata "kmac" "VBoxInternal/Devices/efi/0/Config/DmiSystemVersion" "1.0"
# VBoxManage setextradata "kmac" "VBoxInternal/Devices/efi/0/Config/DmiBoardProduct" "Iloveapple"
# VBoxManage setextradata "kmac" "VBoxInternal/Devices/smc/0/Config/DeviceKey" "ourhardworkbythesewordsguardedpleasedontsteal(c)AppleComputerInc"
# VBoxManage setextradata "kmac" "VBoxInternal/Devices/smc/0/Config/GetKeyFromRealSMC" 1
````
