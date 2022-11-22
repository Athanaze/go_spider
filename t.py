import re

txt = "https://aws.amazon.com/training/digital/partners/aws-training-partners/?sb_tr_atp"
x = re.search(".*aws.amazon.com", txt)
if x == None:
    print("no aws match")
else:
    print("aws match")
    print(x)
