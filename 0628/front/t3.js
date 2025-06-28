// 定义一个类 Animal,包含构造函数和一个 speak 方法,然后扩展一个子类 Dog,
// 添加一个bark 方法并输出 "Woof! "

class Animal{
    constructor(name){
        this.name = name;
    }
    speak(){
        console.log('this is speaker');
    }
}

class Dog extends Animal{
    constructor(name){
        super(name);
    }

    bark(){
        console.log('Wolf!');
    }
}

const dog = new Dog('dog');
dog.bark();