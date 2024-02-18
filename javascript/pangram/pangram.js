//
// This is only a SKELETON file for the 'Pangram' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const isPangram = (pangram) => {
  const alphabet = new Set('abcdefghijklmnopqrstuvwxyz');
  const cleanedPangram = pangram.toLowerCase().replace(/[^a-z]/g, '');

  for (let char of cleanedPangram) {
    alphabet.delete(char);
    if (alphabet.size === 0) {
      return true;
    }
  }

  return false;
};
