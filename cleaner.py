import validators
import re

def awsMatch(txt):
    x = re.search(".*aws.amazon.com", txt)
    if x == None:
        return False
    else:
        return True

uSets = {"https://www.geeksforgeeks.org/sets-in-python/"}

with open("b52") as b:
    lines = b.readlines()
    for l in lines:
        a = l.removesuffix("\n")
        if validators.url(a) and (not awsMatch(a)):
            uSets.add(a)
        
for u in uSets:
    print(u)