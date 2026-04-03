#section CONST
letters = "abcdefghijklmnopqrstuvwxyz"
LCLetters = list(letters)
UCLetters = list(letters.upper())
#section MAIN
def CeasarShift(text:str, shift:int) -> str:
    modified_text = ""
    for char in text:
        if char in LCLetters:
            index = LCLetters.index(char)
            new_index = (index + shift) % 26
            modified_text += LCLetters[new_index]
        elif char in UCLetters:
            index = UCLetters.index(char)
            new_index = (index + shift) % 26
            modified_text += UCLetters[new_index]
        else:
            modified_text += char
    return modified_text