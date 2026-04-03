import c
from typing import Literal
def shift(string:str, type:Literal[1,-1]=1):
    newstring = []
    stringsplit = string.split(" ")
    for word in stringsplit:
        newword = ""
        if stringsplit.index(word) % 2 == 0:
            strengh = len(word)
            for char in word:
                newword += c.CeasarShift(char,type*strengh)
                strengh += -1
        
        else:
            strengh = len(word)
            for char in word:
                newword += c.CeasarShift(char,type*-strengh)
                strengh += -1
        newstring.append(newword)
        newword = ""
    return " ".join(newstring)
print(shift("hi there"))
print(shift("jj odbpd",-1))