---
# Golang Concurrency
---

---
---
### Concurrency 
             +Concurrency is about multiple things happening at same time in random order
             + Concurrent programs may or may not run in parallel. Concurrency is more about the structure of a program, which enables it to potentially run in parallel.
             + Go provides a built-in support for concurrency.
#### ex
       + Every morning I drink my coffee while I fix my breakfast. I could say I am working on both tasks concurrently, but I’m not actually executing them simultaneously (in parallel). At any given point in time I am either drinking from my cup or cutting pancakes, but never both at the same moment. I’m bouncing back-and-forth between the two activities concurrently.
---
---
### Parallelim
           +  Parallelism is when multiple processes are carried out simultaneously by multiple threads or processors.
#### ex
       + Every morning I drink my coffee while watching my Tv. I could say I am working on both  in Parallel(im dring also and watching tv also).
---

    + Concurrency involves structuring a program so that `two or more tasks may be in progress simultaneously`, whereas parallelism allows for `two or more tasks to be executed simultaneously`. `SNote that while parallelism requires more than one processor or thread, concurrency does not`.
### Go’s Concurrency Tool Set
    1. goroutines
    2. channels
    3. select
    4. sync package
             
---