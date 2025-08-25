# oro

Oro is a scripting language that takes inspiration from Lua, Python,Go, and C.

Disclaimer: **For educational and recreational purposes only.**

Disclaimer: Everything is subject to change.

## Notes

- interpreted as AST
- TODO: refactor to interpret bytecode after completion of MVP
- automatic memory management with garbage collection
- dynamically typed
- number - integers (later (maybe): floats,'e' 'E', hexadecimal)
- string - strings and characters (ASCII), delimited by matching single/double quotes
- identifiers - `a-zA-Z0-9_`, cannot start with number
- ignore white space and comments between tokens
- white space - ASCII white space characters
- TODO: case sensitive or insensitive
- TODO: variables (global, local?)
- TODO: How to handle overflows?
- TODO: How to handle type conversions and coercions
- TODO: How to handle errors
- TODO: standard library
- TODO: look into regex

## Keywords

- variable
- function
- if
- else
- for
- while
- and
- or
- not
- null
- eof
- break
- continue

```
 +     -     *     /     %     ^     #
 &     ~     |     <<    >>    //
 ==    ~=    <=    >=    <     >     =
 (     )     {     }     [     ]
 ;     :     ,     .    ..
 '      "
```

## Escape Sequences

```
 '\a' (bell), '\b' (backspace), '\f' (form feed), '\n' (newline),
 '\r' (carriage return), '\t' (horizontal tab), '\v' (vertical tab),
 '\\' (backslash), '\"' (quotation mark [double quote]), and '\'' (apostrophe [single quote]).
```

## Arithmetic Operators

- `+` - add
- `-` - subtract, negate
- `*` - multiply
- `/` - divide
- `%` - modulo

## Bitwise Operators

- convert to integers

- `&` - bitwise AND
- `|` - bitwise OR
- `~` - bitwise NOT
- `^` - bitwise XOR
- `>>` - bitwise right shift
- `<<` - bitwise left shift

## Relational Operators

- returns boolean

- `==` - equality
- `!=` - inequality
- `<` - less than
- `>` - greater than
- `<=` - less or equal
- `>=` - greater or equal

## Logical Operators

- `||` - logical OR (disjunction)
- `&&` - logical AND (conjunction)
- `!` - logical NOT (negation)

## Precedence

- low to high
- FIX: some operators need to be checked

```
or
 and
 <     >     <=    >=    ~=    ==
 |
 ~
 &
 <<    >>
 ..
 +     -
 *     /     //    %
 unary operators (not   #     -     ~)
 ^
```

## Examples

```oro
var x = 5;
var y = 3;

fn add(a, b) {
    return a + b;
}

var result = add(x, y);
print(result);
```

## References

- [Learning Go](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/)
- [Lua 5.4 Reference Manual](https://www.lua.org/manual/5.4/)
- [Crafting Interpreters](https://craftinginterpreters.com/index.html)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)
- [Writing An Interpreter in Go](https://interpreterbook.com/)
