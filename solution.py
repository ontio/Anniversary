from ontology.interop.System.Runtime import Notify

def Main(operation, args):
    if operation == "match":
        pattern = args[0]
        text = args[1]
        return match(pattern, text)

    return False

def match(pattern, text):
  result = False
  if pattern[0] == "^":
    result = matchExact(pattern, text, 1, 0, 1)
  else:
    result = matchExact(concat(".*", pattern), text, 0, 0, 1)

  Notify(["match", pattern, text, result])
  return result

def matchExact(pattern, text, pi, ti, inc):
  if pi == len(pattern):
    return True
  elif ti == len(text) and pattern[pi] == "$":
    return True
  elif ti > len(text):
    return False
  elif pi + inc < len(pattern):
    if pattern[pi] == "\\":
      return matchSegment(pattern, text, pi + 1, ti, inc, True) and matchExact(pattern, text, pi + inc + 1, ti + 1, 1)
    elif pattern[pi + inc] == "?":
      # return matchQuestion(pattern, text, pi, ti, inc)
      return (matchSegment(pattern, text, pi, ti, inc, False) and matchExact(pattern, text, pi + inc + 1, ti + 1, 1)) or matchExact(pattern, text, pi + inc + 1, ti, 1)
    elif pattern[pi + inc] == "*":
      # return matchStar(pattern, text, pi, ti, inc)
      return (matchSegment(pattern, text, pi, ti, inc, False) and matchExact(pattern, text, pi, ti + 1, inc)) or matchExact(pattern, text, pi + inc + 1, ti, 1)
    elif pattern[pi + inc] == "+":
      # return matchPlus(pattern, text, pi, ti, inc)
      modifiedPattern = concat(concat(substr(pattern, 0, pi + inc), "*"), substr(pattern, pi + inc + 1, len(pattern) - (pi + inc + 1)))
      return matchSegment(pattern, text, pi, ti, inc, False) and ((matchSegment(modifiedPattern, text, pi, ti + 1, inc, False) and matchExact(modifiedPattern, text, pi, ti + 2, inc)) or matchExact(modifiedPattern, text, pi + inc + 1, ti + 1, 1))
  
  if inc == 1 and pattern[pi] == "[":
    # return matchBracket(pattern, text, pi, ti)
    # endBracket = findEndBracket(pattern, pi)
    escaping = False

    for i in range(pi, len(pattern)):
      if not escaping and pattern[i] == "\\":
        escaping = True
      elif not escaping and pattern[i] == "]":
        return matchExact(pattern, text, pi, ti, i - pi + 1)
      else:
        escaping = False
    return False
  
  return matchSegment(pattern, text, pi, ti, inc, False) and matchExact(pattern, text, pi + inc, ti + 1, 1)


def matchSegment(pattern, text, pi, ti, inc, escaping):
  if inc > 1:
    invert = pattern[pi + 1] == "^"
    escapingMulti = False

    if ti == len(text):
      return False

    for i in range(pi + 1 + invert, pi + inc - 1):
      if not escapingMulti and pattern[i] == "\\":
        escapingMulti = True
      else:
        # if matchSingle(pattern, text, i, ti, True):
        if pattern[i] == text[ti]:
          return not invert
        
        escapingMulti = False

    return invert
  else:
    # return matchSingle(pattern, text, pi, ti, escaping)
    if pi == len(pattern):
      return True
    if ti == len(text):
      return False

    return (not escaping and pattern[pi] == ".") or pattern[pi] == text[ti]
