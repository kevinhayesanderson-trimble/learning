 #include <stdio.h>
 #include <unistd.h>
 
 int main(){
    printf("Hello, World!\n");// new line character flushes the output buffer in the user space before the segfault occurs.
    // we can also use fflush(stdout) to flush the output buffer before the segfault occurs.
    // we can also use write(1, "Hello, World!\n", 14); to flush the output buffer before the segfault occurs.
    write(1, "Hello, World!\n", 14);
    int *a;
    *a =5;
 }