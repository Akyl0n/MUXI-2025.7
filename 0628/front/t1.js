// 改写为箭头函数
// function multiply(a, b) {
// return a * b;
// }

const multiply = (a , b) => {return a+b;}

// 使用解构从数组 [1, 2, 3] 中取出第一个和第三个值。

const [a , b , c] = [1 , 2 , 3];
console.log(a , c);