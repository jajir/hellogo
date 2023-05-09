# Printer

This program allows to print simple pattern.

EV3 have layoyour of connector:

```
 A   B   C   D
+-------------+
|+-----------+|
||           ||
|+-----------+|
| <>          |
|     /=\     |
|  <- <=> ->  |
|     \=/     |
|EV3          |
+-------------+
 1   2   3   4 
```
Printer should be conected in a following way

* `A` large LEGO tacho motor for axis X.
* `1` touch sensor detecting edges of axis X.
* `B` large LEGO tacho motor for axis Y.
* `2` touch sensor detecting edges of axis Y.
* `C` large LEGO tacho motor for pen.
* `3` touch sensor for immediate ending of printing.

Motors `A` and `B` controll movement of pen on paper in all directions. Motor `C` control when pen is actualy drawing on paper or is in upper position.

Axis and paper coordinates are oriented in a following way:

```
[0,0]
+-------------+ ^
|             | a
|             | x
|             | i
|             | s
|             | 
|             | Y
+-------------+
<-- axis X ---> 
```


