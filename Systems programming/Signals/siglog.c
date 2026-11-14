//#define _POSIX_C_SOURCE 200809L

#include <signal.h>
#include <stdio.h>
#include <stdint.h>
#include <string.h>
#include <unistd.h>

volatile uint64_t handled = 0;

void handle(int sig)
{
    handled |= (1 << sig);
    printf("Caught %d: %s (%d total)\n", 
        sig,
        strsignal(sig),
        __builtin_popcount(handled));

    if (sig == SIGINT || sig == SIGTERM || sig == SIGQUIT)
    {
        printf("Exiting...\n");
        _exit(0);
    }
}

int main(int argc, char *argv[])
{
    // register all valid signals
    for (int i = 0; i < NSIG; i++)
    {
        signal(i, handle);
    }

    // spin
    for (;;)
        sleep(1);
}