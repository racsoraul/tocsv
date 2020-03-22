# Easing my life
Tool to automate the transformation of a really ugly file with lots of inconsistent data to CSV format.

> **NOTE:** This a very specific tool to help me with a really trivial task. It's not in any way generic. However, I'm leaving this open source in case it's useful for anyone as a reference, maybe.

# Format of the input file

Ideally, each line was expected to have the following data:

```
CORRELATIVE_NUMBER      NIT     DUI
```
> **NIT** and **DUI** are documents of identification in my country (El Salvador). NIT is expected to have 14 numbers, and DUI is expected to have 9 numbers.

## Reality 

However, the file has lines with the following data:

- Empty line. No data at all
- Both DUI and NIT are NULL:
```
123     NULL        NULL
```
- The data is incomplete. Either NIT or DUI is missing, and the other one is NULL. Difficult to know for which that NULL stands for:
```
123     NULL
```
- The data is incomplete. It has the correlative and NIT or DUI:
```
123     12345678901234
```
- The data is complete. All 3 values are there:
```
123     12345678901234      123456789
```
- The data is complete or incomplete, and has random extra spaces at the beggining, end or in the middle of the values:
```
      123        12345678901234      123456789   
      123               123456789   
```
