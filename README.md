# ZFS snapshotter plugin

WORK IN PROGRESS

ZFS snapshotter plugin for containerd, based on the [official containerd zfs snapshotter](https://github.com/containerd/containerd), yet git different from the original since it's using a real API interface to ZFS rather than executing command. 

This plugin is tested on Arch Linux for ZFS On Linux 2.x and Ubuntu for 1.8.x, targetting 2.x first.

## Usage

The plugin is built-in by default since containerd 1.1.
No need to recompile containerd or execute a proxy snapshotter process.

1. Set up a ZFS filesystem.
The ZFS filesystem name is arbitrary but the mount point needs to be `/var/lib/containerd/io.containerd.snapshotter.v1.zfs`, when the containerd root is set to `/var/lib/containerd`.
```console
$ zfs create -o mountpoint=/var/lib/containerd/io.containerd.snapshotter.v1.zfs your-zpool/containerd
```
