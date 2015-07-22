package main

import (
	"syscall"
)

// Disk Space total and free bytes available for path like "df command"
func DiskSpace(path string) (diskspaceUtil float64, diskspaceUsed, diskspaceAvail int, diskinodesUtil float64, err error) {
	s := syscall.Statfs_t{}
	err = syscall.Statfs(path, &s)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	_total := int(s.Bsize) * int(s.Blocks)
	_avail := int(s.Bsize) * int(s.Bavail)
	_used := _total - _avail
	_diskspaceUtil := (float64(_used) / float64(_total)) * 100

	_inodesTotal := int(s.Files)
	_inodesFree := int(s.Ffree)
	_inodesUtil := 100 * (1 - float64(_inodesFree)/float64(_inodesTotal))
	return _diskspaceUtil, _used, _avail, _inodesUtil, err
}
