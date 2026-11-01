#define _GNU_SOURCE

#include <sched.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/time.h>
#include <sys/resource.h>
#include <stdio.h>
#include <time.h>

#define SLEEP_SEC 3
#define NUM_MULS 100000000
#define NUM_MALLOCS 100000
#define MALLOC_SIZE 1000

struct profile_times
{
  struct timeval wall_start;
  struct rusage cpu_start;
};

void profile_start(struct profile_times *t)
{
  gettimeofday(&t->wall_start, NULL);
  getrusage(RUSAGE_SELF, &t->cpu_start);
}

void log_prefix(const struct timespec *program_start)
{
  struct timespec epoch_now;
  struct timespec monotonic_now;
  if (clock_gettime(CLOCK_REALTIME, &epoch_now) == -1)
  {
    perror("clock_gettime");
    exit(EXIT_FAILURE);
  }
  if (clock_gettime(CLOCK_MONOTONIC, &monotonic_now) == -1)
  {
    perror("clock_gettime");
    exit(EXIT_FAILURE);
  }
  double epoch_seconds = epoch_now.tv_sec + epoch_now.tv_nsec / 1000000000.0;
  double elapsed_seconds = (monotonic_now.tv_sec - program_start->tv_sec) + (monotonic_now.tv_nsec - program_start->tv_nsec) / 1000000000.0;
  printf("epoch=%.3f elapsed=%.3f [pid %d, cpu %d] ", epoch_seconds, elapsed_seconds, getpid(), sched_getcpu());
}

void profile_log(struct profile_times *t, const struct timespec *program_start)
{
  struct timeval wall_end;
  struct rusage cpu_end;
  gettimeofday(&wall_end, NULL);
  getrusage(RUSAGE_SELF, &cpu_end);
  double wall_seconds = (wall_end.tv_sec - t->wall_start.tv_sec) + (wall_end.tv_usec - t->wall_start.tv_usec) / 1000000.0;
  double user_seconds = (cpu_end.ru_utime.tv_sec - t->cpu_start.ru_utime.tv_sec) + (cpu_end.ru_utime.tv_usec - t->cpu_start.ru_utime.tv_usec) / 1000000.0;
  double system_seconds = (cpu_end.ru_stime.tv_sec - t->cpu_start.ru_stime.tv_sec) + (cpu_end.ru_stime.tv_usec - t->cpu_start.ru_stime.tv_usec) / 1000000.0;
  double cpu_seconds = user_seconds + system_seconds;
  log_prefix(program_start);
  printf("real: %.3fs user: %.3fs sys: %.3fs\n", wall_seconds, user_seconds, system_seconds);
  //printf("CPU time:%.6f , Wall time:%.6f \n", cpu_seconds, wall_seconds);
}

int main(int argc, char *argv[])
{
  struct timespec program_start;
  if (clock_gettime(CLOCK_MONOTONIC, &program_start) == -1)
  {
    perror("clock_gettime");
    return EXIT_FAILURE;
  }
  struct profile_times t;

  // profile doing a bunch of floating point muls
  float x = 1.0;
  log_prefix(&program_start);
  printf("%d fmuls\n", NUM_MULS);
  profile_start(&t);
  for (int i = 0; i < NUM_MULS; i++)
    x *= 1.1;
  profile_log(&t, &program_start);

  // profile doing a bunch of mallocs
  log_prefix(&program_start);
  printf("%d mallocs of size %d\n", NUM_MALLOCS, MALLOC_SIZE);
  profile_start(&t);
  void *p;
  for (int i = 0; i < NUM_MALLOCS; i++)
    p = malloc(MALLOC_SIZE);
  profile_log(&t, &program_start);

  // profile sleeping
  log_prefix(&program_start);
  printf("sleeping for %d\n", SLEEP_SEC);
  profile_start(&t);
  sleep(SLEEP_SEC);
  profile_log(&t, &program_start);
}

/*
kevin@KevinLaptop:/mnt/c/Users/kevin/learning/Systems programming/cputime$ gcc -Wall -Wextra -o cputime cputime.c
cputime.c: In function ‘profile_log’:
cputime.c:56:10: warning: unused variable ‘cpu_seconds’ [-Wunused-variable]
   56 |   double cpu_seconds = user_seconds + system_seconds;
      |          ^~~~~~~~~~~
cputime.c: In function ‘main’:
cputime.c:85:9: warning: variable ‘p’ set but not used [-Wunused-but-set-variable]
   85 |   void *p;
      |         ^
cputime.c:62:14: warning: unused parameter ‘argc’ [-Wunused-parameter]
   62 | int main(int argc, char *argv[])
      |          ~~~~^~~~
cputime.c:62:26: warning: unused parameter ‘argv’ [-Wunused-parameter]
   62 | int main(int argc, char *argv[])
      |                    ~~~~~~^~~~~~
kevin@KevinLaptop:/mnt/c/Users/kevin/learning/Systems programming/cputime$ ./cputime
epoch=1793539129.747 elapsed=0.000 [pid 15276, cpu 3] 100000000 fmuls
epoch=1793539130.139 elapsed=0.392 [pid 15276, cpu 3] real: 0.392s user: 0.392s sys: 0.000s
epoch=1793539130.139 elapsed=0.392 [pid 15276, cpu 3] 100000 mallocs of size 1000
epoch=1793539130.174 elapsed=0.427 [pid 15276, cpu 3] real: 0.034s user: 0.000s sys: 0.034s
epoch=1793539130.174 elapsed=0.427 [pid 15276, cpu 3] sleeping for 3
epoch=1793539133.174 elapsed=3.427 [pid 15276, cpu 3] real: 3.000s user: 0.000s sys: 0.000s
kevin@KevinLaptop:/mnt/c/Users/kevin/learning/Systems programming/cputime$ ./cputime & ./cputime
[1] 15277
epoch=1793539137.722 elapsed=0.000 [pid 15278, cpu 3] 100000000 fmuls
epoch=1793539137.722 elapsed=0.000 [pid 15277, cpu 4] 100000000 fmuls
epoch=1793539138.131 elapsed=0.408 [pid 15277, cpu 4] real: 0.408s user: 0.406s sys: 0.003s
epoch=1793539138.131 elapsed=0.408 [pid 15277, cpu 4] 100000 mallocs of size 1000
epoch=1793539138.137 elapsed=0.414 [pid 15278, cpu 3] real: 0.414s user: 0.414s sys: 0.000s
epoch=1793539138.137 elapsed=0.414 [pid 15278, cpu 3] 100000 mallocs of size 1000
epoch=1793539138.166 elapsed=0.444 [pid 15277, cpu 4] real: 0.035s user: 0.003s sys: 0.032s
epoch=1793539138.166 elapsed=0.444 [pid 15277, cpu 4] sleeping for 3
epoch=1793539138.175 elapsed=0.453 [pid 15278, cpu 3] real: 0.038s user: 0.000s sys: 0.038s
epoch=1793539138.175 elapsed=0.453 [pid 15278, cpu 3] sleeping for 3
epoch=1793539141.166 elapsed=3.444 [pid 15277, cpu 4] real: 3.000s user: 0.000s sys: 0.000s
epoch=1793539141.175 elapsed=3.453 [pid 15278, cpu 3] real: 3.000s user: 0.000s sys: 0.000s
[1]+  Done                       ./cputime
kevin@KevinLaptop:/mnt/c/Users/kevin/learning/Systems programming/cputime$
*/