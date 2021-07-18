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

What about `(()*()*`? Choose empty string for the left-most `*`
and `)` for the right-most: the expression has balanced parens.
