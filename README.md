AWS CloudWatch Monitoring Program
=========================================================

# Put Linux System metrics to AWS CloudWatch

## Memory
* Memory Utilization -  Memory usage in percent
* Memory Used - Used memory in bytes
* Memory Available - Available memory in bytes
* Swap Utilization - Swap usage in percent
* Swap Used - Swap used in bytes

## Disk
* Disk Space Utilization - Disk space usage in percent
* Disk Space Used - Disk space used in bytes
* Disk Space Available - Disk space available in bytes
* Linux partition inode usage - Disk parttion inodes usage in percent

# Usage

'''
go-aws-mon --namespace=<NAMESPACE> --mem-util --mem-used --mem-avail --swap-util --swap-used  --disk-space-util --disk-inode-util --disk-space-used --disk-space-avail --disp-path=PATH
'''


Allen Chen(a3linux X gmail.com)
