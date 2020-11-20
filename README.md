# Example disk utility snap

Build using snapcraft and connect the interfaces:
```
sudo snap connect disk-util:hardware-observe
sudo snap connect disk-util:udisks2
sudo snap connect disk-util:block-devices
```


## Usage

*Listing block devices*
```
disk-util.list
```

*Creating a partition*
Partitions need to be created with a valid type e.g. 0x83, size and offset (in bytes).
Note that the block device needs to be provided (not the partition) i.e. `sdd` not `sdd3`.
```
sudo disk-util.create sdd 0x83 55574528 1048576
```

*Create and format a partition*
The `create-format` command creates a partition as a single step. The parameters need
to be supplied as per the `create` command, with the additional format of the partition
e.g. `ext4`.
```
sudo disk-util.create-format sdd 0x83 55574528 1048576 ext4
```

*Delete a partition*
Remove an existing partition.
```
sudo disk-util.delete sdd3
```
