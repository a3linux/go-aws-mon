package main

import (
	"syscall"
)

// Disk Space total and free bytes available for path like "df command"
func DiskSpace(path string) (diskspaceUtil float64, total, free int, err error) {
	s := syscall.Statfs_t{}
	err = syscall.Statfs(path, &s)
	if err != nil {
		return 0, 0, 0, err
	}
	_total := int(s.Bsize) * int(s.Blocks)
	_free := int(s.Bsize) * int(s.Bfree)
	_used := _total - _free
	_diskSpaceUtil := (float64(_used) / float64(_total)) * 100
	return _diskSpaceUtil, _used, _free, err
}
