# Copyright (c) 2019 Yunindyo Prabowo. All Rights Reserved.
#
# This code is licensed under MIT license (see LICENSE for details)

# TASK :
# create number pattern 0 to 0 like below here :
#   xxxxx
#   x   x
#   x   x
#   x   x
#   xxxxx
#   ***************************

#   #
#   #
#   #
#   #
#   #
#   ***************************

#   #####
#       #
#   #####
#   #
#   #####
#   ***************************

#    xxxxx
#        x
#    xxxxx
#        x
#    xxxxx
#   ***************************
# to 9

import sys


def numberPattern0(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if (i == 1 or i == size or y == size or y == 1):
                sys.stdout.write("x")
            else:
                sys.stdout.write(" ")
        sys.stdout.write("\n")


def numberPattern0(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if (i == 1 or i == size or y == size or y == 1):
                sys.stdout.write("x")
            else:
                sys.stdout.write(" ")
        sys.stdout.write("\n")


def numberPattern1(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if (y == 1):
                sys.stdout.write("#")
            else:
                sys.stdout.write(" ")
        sys.stdout.write("\n")


def numberPattern2(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if size % 2 == 0:
                if (i == 1 or i == size or i == size / 2):
                    sys.stdout.write("#")
                elif (y == 1 and i > size / 2):
                    sys.stdout.write("#")
                elif y == size and i < size / 2:
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")
            else:
                if (i == 1 or i == size or i == (size + 1) / 2):
                    sys.stdout.write("#")
                elif (y == 1 and i > (size + 1) / 2):
                    sys.stdout.write("#")
                elif y == size and i < (size + 1) / 2:
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")
        sys.stdout.write("\n")


def numberPattern3(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if size % 2 == 0:
                if (i == 1 or i == size or y == size or i == (size) / 2):
                    sys.stdout.write("h")
                else:
                    sys.stdout.write(" ")
            else:
                if (i == 1 or i == size or y == size or i == (size + 1) / 2):
                    sys.stdout.write("x")
                else:
                    sys.stdout.write(" ")
        sys.stdout.write("\n")


def numberPattern4(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if size % 2 == 0:
                if (i == size / 2 or y == size):
                    sys.stdout.write("#")
                elif y == 1 and i < size / 2:
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")
            else:
                if (i == (size + 1) / 2 or y == size):
                    sys.stdout.write("x")
                elif y == 1 and i < (size + 1) / 2:
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")
        sys.stdout.write("\n")


def numberPattern5(size):

    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if size % 2 == 0:
                if (i == 1 or i == size or i == size / 2):
                    sys.stdout.write("#")
                elif (y == size and i > size / 2):
                    sys.stdout.write("#")
                elif y == 1 and i < size / 2:
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")
            else:
                if (i == 1 or i == size or i == (size + 1) / 2):
                    sys.stdout.write("#")
                elif (y == 1 and i < (size + 1) / 2):
                    sys.stdout.write("#")
                elif y == size and i > (size + 1) / 2:
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")
        sys.stdout.write("\n")


def numberPattern6(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if size % 2 == 0:
                if (i == 1 or i == size or i == (size) / 2 or y == 1):
                    sys.stdout.write("#")
                elif i > (size / 2) and y == size:
                    sys.stdout.write("x")
                else:
                    sys.stdout.write(" ")
            else:
                if (i == 1 or i == size or i == (size + 1) / 2 or y == 1):
                    sys.stdout.write("x")
                elif i > ((size + 1) / 2) and y == size:
                    sys.stdout.write("x")
                else:
                    sys.stdout.write(" ")

        sys.stdout.write("\n")


def numberPattern7(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if (i == 1 or y == size):
                sys.stdout.write("x")
            else:
                sys.stdout.write(" ")
        sys.stdout.write("\n")


def numberPattern8(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if size % 2 == 0:
                if (i == 1 or i == size or i == (size) / 2 or y == 1 or
                        y == size):
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")
            else:
                if (i == 1 or i == size or i == (size + 1) / 2 or y == 1 or
                        y == size):
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")
        sys.stdout.write("\n")


def numberPattern9(size):
    for i in range(1, size + 1):
        for y in range(1, size + 1):
            if size % 2 == 0:
                if (i == 1 or i == size or i == (size) / 2 or y == size):
                    sys.stdout.write("#")
                elif i < (size / 2) and y == 1:
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")
            else:
                if (i == 1 or i == size or i == (size + 1) / 2 or y == size):
                    sys.stdout.write("#")
                elif i < ((size + 1) / 2) and y == 1:
                    sys.stdout.write("#")
                else:
                    sys.stdout.write(" ")

        sys.stdout.write("\n")


size = int(input("Input size: "))
for x in range(0, 10):
    exec("numberPattern" + str(x) + "(size)")
    sys.stdout.write("***************************\n\n")