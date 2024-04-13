# Basic WC

This is Challenge 1 from [Coding Challenges.fyi](https://codingchallenges.fyi/challenges/intro). The idea was to build a simple word count tool.

More details about the challenge can be found [here](https://codingchallenges.fyi/challenges/challenge-wc)

## Learnings

1. When you use the `:=` syntax in Golang, you are creating a new variable. This manifested itself in a very interesting way [here](https://github.com/saphal1998/basic-wc/blob/b27a7c80f5d069fec226b2deef31992e9d5a6168/wc/wc.go#L58). 

Using the `:=` syntax, makes it so that the lifetime of the variable is only for the active scope. So despite the file being declared outside, we lose the handle to the file, since the `defer` was called as soon as we exit the inner `else` scope.

I found this to be very sneaky. But I guess it makes sense.

2. I cannot use a `tee` reader to create multiple streams from one stream, since they share their seek position. Reading from one, lead to reading from another as well. Now sure, I can pipe what I read into something else, but that will also take up its own memory, so I figured, it is better to just open the file again. All of this is [here](https://github.com/saphal1998/basic-wc/blob/512d78abf9ab1fec5bca1880a638f18f80d73826/wc/wc.go#L87). 
I initally thought that using a pipe and a tee together might work, but since all of them share seek positions, it is (in my opinion), impossible to read the same file stream multiple times while holding the contents of the file in memory. (Or rather even if it is, I didn't find it worth the time).
