#include <fcntl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>

#define ONE_MEG (1 << 20)

int main(void)
{
    int f = open("bw-output.bin", O_WRONLY | O_CREAT | O_TRUNC, 0644);
    if (f == -1)
    {
        perror("open");
        return 1;
    }

    struct stat st;
    blkcnt_t priorblocks = -1;
    for (int i = 0; i < ONE_MEG; i++)
    {
        ssize_t bytes_written = write(f, ".", 1);
        if (bytes_written == -1)
        {
            perror("write");
            close(f);
            return 1;
        }
        if (bytes_written != 1)
        {
            fprintf(stderr, "write: expected 1 byte, wrote %zd\n", bytes_written);
            close(f);
            return 1;
        }

        if (fstat(f, &st) == -1)
        {
            perror("fstat");
            close(f);
            return 1;
        }

        if (st.st_blocks != priorblocks)
        {
            printf("Size=%lld blocks=%lld on disk=%lld\n",
                   (long long)st.st_size,
                   (long long)st.st_blocks,
                   (long long)st.st_blocks * 512);
            priorblocks = st.st_blocks;
        }
    }

    if (close(f) == -1)
    {
        perror("close");
        return 1;
    }

    return 0;
}

/*
gcc -Wall -Wextra -o bw bw.c && ./bw
*/