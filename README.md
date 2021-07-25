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

## Interview Analysis

The candidate should talk through this one, if possible.
It seemed easy to mess up all the prefix-character removals,
and composing prefix-character-strings.

[Validation of parens balance]((https://github.com/bediger4000/binary-tree-odd-string-rep)
is itself seemingly a common "[Easy]"
rated interview question.
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

Looks like it's possible to have multiple balanced strings
for inputs with multiple asterisks.
Is it possible to have a single-asterisk input that gives
multiple result balanced strings?
