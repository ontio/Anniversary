"""
Register Bind Contract
"""
from ontology.interop.System.Runtime import Notify

def Main(operation, args):

    if operation == "match":
        pattern = args[0]
        text = args[1]
        return match(pattern, text)

    return False

def match(pattern, text):
    result = True      # Your Implementation Here
    # Notify(["match", pattern, text, result])
    return result

