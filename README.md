# Example disk utility snap

Build using snapcraft and connect the interfaces:
```
sudo snap connect disk-util:hardware-observe
sudo snap connect disk-util:udisks2
sudo snap connect disk-util:block-devices
sudo snap connect disk-util:udisks2-cons disk-util:udisks2-svc
```


## Usage

*Listing block devices*
```
disk-util.list
```


*Formatting a partition*
```
sudo disk-util.format sdd3 ext4
```


