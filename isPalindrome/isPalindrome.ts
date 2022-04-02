const isPalindrome = (s: string) => {
  let mirror = "";
  for (let i = s.length - 1; i >= 0; i--) {
    mirror += s[i];
  }

  return s === mirror;
};
