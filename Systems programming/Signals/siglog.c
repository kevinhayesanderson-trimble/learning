#define _POSIX_C_SOURCE 200809L

#include <errno.h>
#include <signal.h>
#include <stdio.h>
#include <stddef.h>
#include <sys/wait.h>
#include <unistd.h>

static void handle(int sig)
{
    const char *message;
    size_t length;

    switch (sig)
    {
    case SIGBUS:
        message = "Caught SIGBUS\n";
        length = sizeof("Caught SIGBUS\n") - 1;
        break;
    case SIGSEGV:
        message = "Caught SIGSEGV\n";
        length = sizeof("Caught SIGSEGV\n") - 1;
        break;
    case SIGFPE:
        message = "Caught SIGFPE\n";
        length = sizeof("Caught SIGFPE\n") - 1;
        break;
    case SIGILL:
        message = "Caught SIGILL\n";
        length = sizeof("Caught SIGILL\n") - 1;
        break;
    case SIGCHLD:
        message = "Caught SIGCHLD\n";
        length = sizeof("Caught SIGCHLD\n") - 1;
        break;
    case SIGINT:
    case SIGTERM:
    case SIGQUIT:
        message = "Caught termination signal; exiting\n";
        length = sizeof("Caught termination signal; exiting\n") - 1;
        (void)write(STDERR_FILENO, message, length);
        _exit(0);
    default:
        message = "Caught another signal\n";
        length = sizeof("Caught another signal\n") - 1;
        break;
    }

    (void)write(STDERR_FILENO, message, length);
}

int main(void)
{
    const int signals_to_raise[] = {SIGBUS, SIGSEGV, SIGFPE, SIGILL};
    struct sigaction sa = {0};
    sa.sa_handler = handle;
    sa.sa_flags = 0;
    if (sigemptyset(&sa.sa_mask) == -1)
    {
        perror("sigemptyset");
        return 1;
    }

    const int signals_to_catch[] = {
        SIGBUS, SIGSEGV, SIGFPE, SIGILL, SIGCHLD, SIGINT, SIGTERM, SIGQUIT
    };
    for (size_t i = 0; i < sizeof(signals_to_catch) / sizeof(signals_to_catch[0]); ++i)
    {
        if (sigaction(signals_to_catch[i], &sa, NULL) == -1)
        {
            perror("sigaction");
            return 1;
        }
    }

    for (size_t i = 0; i < sizeof(signals_to_raise) / sizeof(signals_to_raise[0]); ++i)
    {
        if (raise(signals_to_raise[i]) != 0)
        {
            fprintf(stderr, "raise failed for signal %d\n", signals_to_raise[i]);
            return 1;
        }
    }

    pid_t child = fork();
    if (child == -1)
    {
        perror("fork");
        return 1;
    }
    if (child == 0)
    {
        _exit(0);
    }

    while (waitpid(child, NULL, 0) == -1)
    {
        if (errno != EINTR)
        {
            perror("waitpid");
            return 1;
        }
    }

    return 0;
}
