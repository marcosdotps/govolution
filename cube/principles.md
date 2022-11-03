## Introduction

This project shows a rotating cube using ASCII character. I created this golang implementation inspired by [this amazing video](https://www.youtube.com/watch?v=p09i_hoFdd0) made by [Servet Gulnaroglu](https://www.youtube.com/c/ServetGulnaroglu). His job is written in C and wanted to check how it looks using golang. Follow him for more amazing coding videos!

The mathematical principles to do this can be found [here](https://en.wikipedia.org/wiki/Rotation_matrix). To sum up, we will use a rotation matix to perform a rotation in Euclidean Space. Therefore to perform a rotationon a plane point with coordinates `v=(x,y)` it should be written as a column vector and multiplied by the matix R.

A rotation matrix describes rotation about the origin and can be characterized as real square matrices whose columns and rows are orthogonal vectors with determinant 1. 

To make it easier, [symolab](https://www.symbolab.com/solver/matrix-multiply-calculator) has been used to perform matrix calculations.

