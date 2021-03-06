package disk

type linuxBindMounter struct {
	delegateMounter Mounter
}

func NewLinuxBindMounter(delegateMounter Mounter) Mounter {
	return linuxBindMounter{delegateMounter}
}

func (m linuxBindMounter) Mount(partitionPath, mountPoint string, mountOptions ...string) error {
	// Filesystems should not be bind mounted
	if partitionPath != "tmpfs" {
		mountOptions = append(mountOptions, "--bind")
	}

	return m.delegateMounter.Mount(partitionPath, mountPoint, mountOptions...)
}

func (m linuxBindMounter) RemountAsReadonly(mountPoint string) error {
	// Remounting mount points mounted originally by warden with '-o ro --bind' flags does not work.
	// See https://lwn.net/Articles/281157/.
	return nil
}

func (m linuxBindMounter) Remount(fromMountPoint, toMountPoint string, mountOptions ...string) error {
	mountOptions = append(mountOptions, "--bind")
	return m.delegateMounter.Remount(fromMountPoint, toMountPoint, mountOptions...)
}

func (m linuxBindMounter) SwapOn(partitionPath string) (err error) {
	return m.delegateMounter.SwapOn(partitionPath)
}

func (m linuxBindMounter) Unmount(partitionOrMountPoint string) (bool, error) {
	return m.delegateMounter.Unmount(partitionOrMountPoint)
}

func (m linuxBindMounter) IsMountPoint(path string) (string, bool, error) {
	return m.delegateMounter.IsMountPoint(path)
}

func (m linuxBindMounter) IsMounted(partitionOrMountPoint string) (bool, error) {
	return m.delegateMounter.IsMounted(partitionOrMountPoint)
}

func (m linuxBindMounter) RemountInPlace(mountPoint string, mountOptions ...string) (err error) {
	return m.delegateMounter.RemountInPlace(mountPoint, mountOptions...)
}
