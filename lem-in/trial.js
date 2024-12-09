var isValid = (s) => {
  for (let index = 1; index < s.length; index++) {
    if (
      (s[index - 1] == "[" && s[index] == "]") ||
      (s[index - 1] == "{" && s[index] == "}") ||
      (s[index - 1] == "(" && s[index] == ")")
    ) {
      if (index == s.length - 1) {
        s = s.slice(0, index - 1);
      } else {
        s = s.slice(0, index - 1) + s.slice(index + 1);
      }
    }
  }
  if (s.length == 0) {
    return true;
  }
  if (containPairs(s)) {
    return isValid(s);
  } else {
    console.log(s);
    return false;
  }
};

var containPairs = (s) => {
  for (let index = 1; index < s.length; index++) {
    if (
      (s[index - 1] == "[" && s[index] == "]") ||
      (s[index - 1] == "{" && s[index] == "}") ||
      (s[index - 1] == "(" && s[index] == ")")
    ) {
      return true;
    }
  }
  return false;
};

console.log(isValid("[]{})"));
