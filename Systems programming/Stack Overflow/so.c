#include <stdio.h>
#include <unistd.h>
#include <sys/resource.h>

void f(int depth, long int bottom)
{
    if (depth % 10000 == 0)
        printf("[pid %d] frame=%d %ld address=%p\n", getpid(), depth, bottom - (long)&depth, (void *)&depth);
    
    // for(;;)
    //     ;
    f(depth + 1, bottom);
}

void start()
{
    int depth = 0;
    f(depth, (long)&depth);
}

int main()
{
    struct rlimit rlim;
    getrlimit(RLIMIT_STACK, &rlim);
    printf("pid=%d stack current max=%ld stack hard limit=%ld\n", getpid(), rlim.rlim_cur, rlim.rlim_max);
    printf("OK\n");
    start();
}

/*
./so &
pid=$!
sleep 0.1
ps -p "$pid" -o pid,comm,args
pmap -x "$pid"
kill "$pid"
wait "$pid" 2>/dev/null
*/