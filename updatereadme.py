# Adapted from https://stackoverflow.com/questions/33365621:
import re
with open('README.md', 'r') as md, open('examples.txt', 'r') as ex:
        readme = md.read()
        examples = ("\n    " +
                    "\n    ".join([l.rstrip() for l in ex.readlines()]) +
                    "\n")
        ntext = re.sub(r'(.*BEGIN EXAMPLES.*?\n+)(.*)(\<\!\-\-\- END EXAMPLES)',r'\1' +
                       r'\1' + examples + r'\3',
                       readme,
                       flags=re.M|re.DOTALL)

with open('README.md', 'w') as fout:
    fout.write(ntext)
