# Daily Coding Problem: Problem #937 [Hard]

This problem was asked by Google.

You're given a string consisting solely of `(`, `)`, and `*`.
`*` can represent either a `(`, `)`, or an empty string.
Determine whether the parentheses are balanced.

For example, `(()*` and `(*)` are balanced.
`)*(` is not balanced.

## Analysis

An asterisk can represent any single token among the inputs.
The programmer has a choice of what to use for any `*` that comes up.
The examples seem to say
"choose so as to favor balanced parens strings if you can".

This screams "backtracking" to me.
Recursively remove a '(', ')' or '\*' from the input string.
Call the recursive function with remaining input string,
and a string composed of the prefix-characters of the input string.
If the prefix character is '\*', the code has to handle
prefixing the prefix-characters-string with a '(', ')' or
no new prefix character at all,
recursing, then popping any prefix paren off the prefix-characters-string.
Validate the prefix-characters string so composed
when the recursion runs out of input characters.

[My code](a1.go) doesn't strictly meet the problem statement,
in that it doesn't give a plain "Yes, balanced" or "No, not balanced" answer.
It does print out all possible balanced parens strings that
it can generate from the alternatives.

## Interview Analysis

The candidate should talk through this one, if possible.
It seemed easy to mess up all the prefix-character removals,
and composing prefix-character-strings.

Validation of parens balance
is itself seemingly a common "easy"-rated interview question.
I used a parity-counting algorithm that starts with
zero, adds 1 for every '(' and subtracts 1 for every ')'.
If the parity ends at something other than zero,
or ever goes negative, the string does not contain balanced parens.

This may be why this problem is rated "[Hard]".
You have to have 2 algorithms,
one to generate candidates based on substituting for '\*',
and one to check the candidates.

I'm guessing that catching ')(' as unbalanced is one of the
things an interviewer would look for.
It's probably good to include test cases for that.

Enumerating potential test cases would be advisable:

* `(()*`, `(*)`, `)*(` are given in the clarifying example.
* `()`, `(*)` `(**)` would be good, see if empty string works.
* `*)`, `(*` would check if code examines both parens.
* Multiple-asterisk cases: `((*)*`, `()*()*`, `(()*()*` seem like a good thing.
* Try some unbalanced cases: `)(`, `)*(` to see if your validating works.

I'll buy this as a "[Hard]" daily coding problem.
You have to understand 2 different algorithms to solve the problem.
Thankfully, you don't need any physical insight to solve this one.

I think the interviewer might expect candidates that don't
see or remember a "counting" do-parens-balance validity check
to do some flailing, especially if they decide to use a destructive
algorithm for checking balance.

Looks like it's possible to have multiple balanced strings
for inputs with multiple asterisks.
Is it possible to have a single-asterisk input that gives
multiple result balanced strings?

# Related Daily Coding Problem

Haven't done this one yet.

## Daily Coding Problem: Problem #984 [Hard]

This problem was asked by Facebook.

Given a string of parentheses,
find the balanced string that can be produced from it using the minimum number
of insertions and deletions.
If there are multiple solutions,
return any of them.

For example,
given "(()", you could return "(())".
Given "))()(", you could return "()()()()".

## Analysis

This is roughly the same as Problem #937,
but does not give an '\*' where you should try to insert something.
