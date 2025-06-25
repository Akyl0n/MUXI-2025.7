// 请你编写一个函数，检查给定的值是否是给定类或超类的实例。

// 可以传递给函数的数据类型没有限制。例如，值或类可能是  undefined 。

function checkIfInstanceOf(obj , classfunction){
    if (obj === null || obj === undefined || classfunction === null || classfunction === undefined){
        return false;
    }
    if(typeof classfunction !== 'function'){
        return false;
    }
    let proto = Object.getPrototypeOf(obj);
    while(proto !== null){
        if(proto === classfunction.prototype){
            return true;
        }
        proto = Object.getPrototypeOf(proto);
    }
    return false;
}
