/* eslint-disable @typescript-eslint/no-unused-vars */

import evaluate from './luhn'
const cc: string = '4539479425649801' // 4539479425649801 to test valid
const inputArray: string[] = ['45394d9425649801', '45394794256asdf49801', '45394794d25649801', '4539479425649801', '4539479425649801', '4539479425649801', '45394d79425649801', '4539479425649801', '4539479425649801']


async function bootstrap() {

    const startTime = performance.now()
    
    for (let x in inputArray) {
        let result = await evaluate(inputArray[x])
        //console.log(result)
    }

    const endTime = performance.now()

    console.log(`It takes ${(endTime - startTime)} ms to evalute ${inputArray.length} credit cards`)
}

bootstrap()