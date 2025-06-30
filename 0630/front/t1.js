console.log('1');

setTimeout(() => {
    console.log('2');
    new Promise(resolve => {
        console.log('3');
        resolve();
    }).then(() => {
        console.log('4');
    });
}, 0);

new Promise((resolve) => {
    console.log('5');
    resolve();
}).then(() => {
    console.log('6');
    return new Promise(resolve => {
        console.log('7');
        resolve();
    }).then(() => {
        console.log('8');
    });
}).then(() => {
    console.log('9');
});

Promise.resolve().then(() => {
    console.log('10');
    setTimeout(() => {
        console.log('11');
    }, 0);
    Promise.resolve().then(() => {
        console.log('12');
    });
});

console.log('13');

setTimeout(() => {
    console.log('14');
}, 0);

console.log('15');