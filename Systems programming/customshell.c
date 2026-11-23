// execute a program by name
// support a few buitins (e.g quit, help)
// do this in a loop
// no support for strings (tokenize by whitespace and tabs)
// handle EOF and SIGINT coorectly(EOF should terminate shell, SIGINT should terminate child)

//1. implement REPL, but echoing instead of eval'ing (fgets/printf, ie not readline)
//2. tokenize cmd, print argc/argv (strtok)
//3. actually execute (fork/execvp)
//4. refactor/details, check success criteria incl signal handling

#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>
#include <signal.h>

volatile pid_t childpid = 0;

void sigint_handler(int sig){
    int saved_errno = errno;
    if (childpid > 0 && kill(childpid, SIGINT) == -1 && errno != ESRCH) {
        static const char message[] = "Error sending SIGINT to child\n";
        (void)write(STDERR_FILENO, message, sizeof(message) - 1);
    }
    errno = saved_errno;
}

#define MAX_CMD 4096
#define MAX_ARGV 256
#define PROMPT "> "
#define SEP " \t\n"

int main(){
    if (signal(SIGINT, sigint_handler) == SIG_ERR) {
        perror("signal");
        return 1;
    }

    char cmd[MAX_CMD];
    int argc;
    char* argv[MAX_ARGV];

    while(1){
        printf(PROMPT);
        if (NULL == fgets(cmd, MAX_CMD, stdin) && ferror(stdin)) {
            perror("fgets error");
            exit(1);
        }
        if(feof(stdin))exit(0);

        // tokenize
        argc = 0;
        argv[argc] = strtok(cmd, SEP);

        while(argv[argc] != NULL){
            argv[++argc] = strtok(NULL, SEP);
        }

        // printf("%d args: \n", argc);
        // for(int i=0; i < argc; i++){
        //     printf("%s, ",argv[i]);
        // }

        // eval
        if(argc == 0) continue;

        if(0 == strcmp(argv[0], "quit")) exit(0);
        if(0 == strcmp(argv[0], "help")) {
            printf("Welcome to my shell, builtins are quit and help\n");
            continue;
        }

        if((childpid = fork()) < 0) {
            perror("fork error");
            exit(1);
        }

        if (childpid == 0){// child
            //exec
            if(execvp(argv[0], argv) < 0){
                perror("execvp error");
                exit(1);
            }
            // we should never get here, since the exec passes
            exit(1);
        }

        //parent
        int status;
        while (waitpid(childpid, &status, 0) == -1) {
            if (errno != EINTR) {
                perror("waitpid");
                exit(1);
            }
        }
        childpid=0;
        
    }
}