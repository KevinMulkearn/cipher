# cipher #

Simple cipher generator Go (golang) application. Takes an input file (or string) to encrypt and produces an encrpyted output file (or string). Note: currently setup for handle lowercase characters only.

## How to use ##
Select the file to be encrpyted followed by the tpye of cipher you would like to use and optionally the name of the output file.

Input/OUtput file Example:
```
    cipher -i myinputfile.txt -c caeser -o myoutputfile.txt
```

You may also pass a string literal as an input where the result will be printed to the screen.

Test string Example:
```
    cipher -c caeser -m "Hello World"
```

## Ciphers ##
1. [Caesar Cipher](https://en.wikipedia.org/wiki/Caesar_cipher)
2. [Monoalphabetic Cipher](https://crypto.interactive-maths.com/monoalphabetic-substitution-ciphers.html) - TODO
3. [Homophonic Substitution Cipher](https://crypto.interactive-maths.com/homophonic-substitution.html) - TODO
4. [Polygram Substitution Cipher](https://en.wikipedia.org/wiki/Polygraphic_substitution) - TODO
5. [Polyalphabetic Substitution Cipher](https://crypto.interactive-maths.com/polyalphabetic-substitution-ciphers.html) - TODO
6. [Playfair Cipher](https://en.wikipedia.org/wiki/Playfair_cipher) - TODO
7. [Hill Cipher](https://en.wikipedia.org/wiki/Hill_cipher) - TODO
