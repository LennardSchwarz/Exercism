export function find(haystack: number[], needle: number): number {
  // check if haystack is an array
  if (!Array.isArray(haystack)) {
    throw new Error('Value not an array');
  }
  // find the middle index
  const middleIndex = findMiddleIndex(haystack);
  
  // check if the middle index is the needle
  if (haystack[middleIndex] === needle) {
    return middleIndex;
  }
  // check if the needle is less than the middle index
  if (needle < haystack[middleIndex]) {
    return find(haystack.slice(0, middleIndex), needle);
  }
  // check if the needle is greater than the middle index
  if (needle > haystack[middleIndex]) {
    return find(haystack.slice(middleIndex + 1), needle) + middleIndex + 1;
  }
  // if the needle is not found
  throw new Error('Value not in array');
}


function findMiddleIndex(haystack: number[]): number {
  if (haystack.length % 2 === 0) {
    return haystack.length / 2;
  }
  return (haystack.length - 1) / 2;
}