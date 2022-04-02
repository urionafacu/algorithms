const selectionSort = (nums: number[]): number[] => {
  for (let i = 0; i < nums.length; i++) {
    var min = i;
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[j] < nums[min]) {
        min = j;
      }
    }
    if (i !== min) {
      var aux = nums[i];
      nums[i] = nums[min];
      nums[min] = aux;
    }
  }
  return nums;
};

const medianOfTwoSortedArrays = (nums1: number[], nums2: number[]): number => {
  const nums = [...nums1, ...nums2];

  selectionSort(nums);

  const isOddSize = nums.length % 2 === 0;
  const size = nums.length;
  if (isOddSize) {
    const middle = size / 2;
    const left = nums[middle - 1];
    const right = nums[middle];
    const sum = left + right;
    return sum / 2;
  }
  const middle = Math.floor(size / 2);
  return nums[middle];
};

console.log(medianOfTwoSortedArrays([1, 3], [2, 3]));
