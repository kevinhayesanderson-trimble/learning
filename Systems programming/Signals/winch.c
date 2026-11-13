#include <signal.h>
#include <stdio.h>
#include <sys/ioctl.h>

volatile sig_atomic_t resized = 0;
_Atomic int resize_count = 0;

void handler(int sig) {
    (void)sig;
    resized = 1;
    resize_count++;
}

int main(void) {
    struct sigaction sa = {0};
    sa.sa_handler = handler;
    sa.sa_flags = 0;
    if (sigemptyset(&sa.sa_mask) == -1) {
        perror("sigemptyset");
        return 1;
    }
    if (sigaction(SIGWINCH, &sa, NULL) == -1) {
        perror("sigaction");
        return 1;
    }
    struct winsize ws;
    printf("Waiting for signal\n");
    for (;;) {
        if (resized) {
            ioctl(0, TIOCGWINSZ, &ws);
            printf("%d: Columns: %d, Rows: %d\n", resize_count, ws.ws_col, ws.ws_row);
            resized = 0;
        }
    }
}