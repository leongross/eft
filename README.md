# Elf To Shellcode (eft)

Tool to extract `.text` section of `ELF` executable and display it as an escaped byte sequence.

```
$ go build
$ go install
$ ./eft <elf>
\x54\x68\x9d\x80\x04\x08\x31\xc0\x31\xdb\x31\xc9\x31\xd2\x68\x43\x54
\x46\x3a\x68\x74\x68\x65\x20\x68\x61\x72\x74\x20\x68\x73\x20\x73\x74
\x68\x4c\x65\x74\x27\x89\xe1\xb2\x14\xb3\x01\xb0\x04\xcd\x80\x31\xdb
\xb2\x3c\xb0\x03\xcd\x80\x83\xc4\x14\xc3\x5c\x31\xc0\x40\xcd\x80
```

Verify using `objdump`:
```
$ objdump -s -j .text <elf>
Contents of section .text:
 8048060 54689d80 040831c0 31db31c9 31d26843  Th....1.1.1.1.hC
 8048070 54463a68 74686520 68617274 20687320  TF:hthe hart hs
 8048080 7374684c 65742789 e1b214b3 01b004cd  sthLet'.........
 8048090 8031dbb2 3cb003cd 8083c414 c35c31c0  .1..<........\1.
 80480a0 40cd80
```
