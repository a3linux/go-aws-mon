package main

import "github.com/guillermo/go.procmeminfo"

// Memory Usage(percent of used, bytes of Used, Avaliable and Total)
func memoryUsage() (memUtil, memUsed, memAvail, swapUtil, swapUsed float64, err error) {
	meminfo := &procmeminfo.MemInfo{}
	meminfo.Update()

	_memUsed := float64(meminfo.Used())
	_memAvail := float64(meminfo.Available())
	_memTotal := float64(meminfo.Total())
	_memUtil := (_memUsed / _memTotal) * 100
	_swapUtil := float64(meminfo.Swap())
	_swapUsed := float64((*meminfo)["SwapCached"])

	return _memUtil, _memUsed, _memAvail, _swapUtil, _swapUsed, err
}
