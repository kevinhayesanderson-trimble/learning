#include <unistd.h>

int main(){
    //ls | wc
    //0, 1, and 2 are already open as stdin, stdout, and stderr
    int fds[2];

    /*
    pipe(fds) asks the kernel to allocate an in-memory pipe buffer and wire two new file descriptors into the current process's descriptor table:
        fds[0] = read end of the pipe
        fds[1] = write end of the pipe
    */
    pipe(fds);

    if(0 == fork()){
        //child
        
        /*
        dup2(oldfd, newfd)
        File descriptor 1 (stdout) now points to the write end of the pipe (fds[1]).
        When ls runs next and prints filenames to stdout (fd 1), those bytes go straight into the pipe buffer instead of the terminal.
        */
        dup2(fds[1], 1); 
        
        /*
        child - close the read end
        The child only writes (ls), never reads. Closing its read end also avoids keeping an extra reference open unnecessarily.
        */
        close(fds[0]);

        /*
        child - close the original
        fd 1 already points to the write end, so the extra fds[1] descriptor is no longer needed.
        */
        close(fds[1]);
        execlp("ls", "ls", NULL);
    }
    //parent

    /*
    dup2(oldfd, newfd)
    File descriptor 0 (stdin) now points to the read end of the pipe (fds[0]).
    When wc runs next and reads from stdin (fd 0), it pulls bytes straight out of the pipe buffer instead of your keyboard.
    */
    dup2(fds[0], 0);

    /*
    parent - close the original
    fd 0 already points to the read end, so the extra fds[0] descriptor is no longer needed.
    */
    close(fds[0]);

    /*
    parent - close the write end
    Crucial. 
    The parent never writes to the pipe (wc reads). 
    If the parent left its fds[1] open, 
    the kernel's write-end reference count would stay greater than zero even after ls exited. 
    Because wc waits for EOF (read() returning 0), 
    and EOF on a pipe only happens when all write descriptors across all processes are closed,
    wc would hang forever waiting for more input.
    */
    close(fds[1]);
    execlp("wc", "wc", NULL);
}