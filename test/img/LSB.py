from PIL import Image
import numpy as np


def split(text: str, slen: int):
    splitlist = []
    counter = 0
    newsplit = ""
    for char in text:
        newsplit += char
        counter += 1
        if counter == slen:
            splitlist.append(newsplit)
            counter = 0
            newsplit = ""
    splitlist.append(newsplit)
    return splitlist


def imgtoparr(path: str):
    img = Image.open(path).convert("RGB")  # get rid of the transparency
    parr = np.asarray(img)
    return parr


def parrtoimg(parr: np.ndarray, outpath: str):
    img = Image.fromarray(parr)
    img.save(outpath)


def steno(parr: np.ndarray, secret: str):
    binsecret = ""
    blen = 0
    newimg = []
    for char in secret:
        binsecret += f"{ord(char):08b}"
    binsecretsplit = split(binsecret, 3)
    i = 0
    bslen = len(binsecretsplit)
    for pixel in parr.reshape(
        -1, 3
    ):  # in this case it was better not to use enumerate IMO
        if i == bslen:
            bdg3 = "000"
        else:
            try:
                bdg3 = binsecretsplit[i]
            except:
                bdg3 = "000"
        i += 1
        red, green, blue = pixel
        RB = list(f"{red:08b}")
        GB = list(f"{green:08b}")

        BB = list(f"{blue:08b}")
        while len(bdg3) < 3:
            bdg3 += ("0")
        RB[-1], GB[-1], BB[-1] = bdg3
        newred = int("".join(RB), 2)
        newgreen = int("".join(GB), 2)
        newblue = int("".join(BB), 2)
        newpixel = [newred, newgreen, newblue]
        newimg.append(newpixel)
    return np.array(newimg, dtype=np.uint8).reshape(parr.shape)


def rev_steno(parr: np.ndarray):
    binsecret = ""
    secret = ""
    for pixel in parr.reshape(-1, 3):
        r, g, b = pixel
        RB, GB, BB = f"{r:08b}", f"{g:08b}", f"{b:08b}"
        binsecret += RB[-1]
        binsecret += GB[-1]
        binsecret += BB[-1]
    bssplit = split(binsecret, 8)
    for char in bssplit:
        if char:
            secret += chr(int(char, 2))
    return secret


if __name__ == "__main__":
    imginput = imgtoparr('image.png')
    secret = steno(imginput, "mypassword")
    secret2 = rev_steno(secret)
    print(secret2)
    parrtoimg(secret, "test.png")
    with open('test.txt','w',encoding='utf-8') as fh:
        pass