blocks - 
blockSize - 4096 bytes
blockCount - 1024
blockSize * blockCount = 4096 * 1024 = 4194304 bytes = 4MB
blockSize vs diskSize - 4096 bytes vs 4194304 bytes = 1/1024 = 0.09765625% of disk size
why is blockSize different from diskSize? - blockSize is the size of a single block, while diskSize is the total size of the disk. The disk is divided into multiple blocks, and each block can store a fixed amount of data. The blockSize is typically smaller than the diskSize to allow for efficient storage and retrieval of data.
