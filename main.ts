/* eslint-disable @typescript-eslint/no-unused-vars */

import evaluate from './luhn'
const cc: string = '4539479425649801' // 4539479425649801 to test valid

async function bootstrap() {

    const startTime = performance.now()

    let result = await evaluate(cc)
    console.log(result)

    const endTime = performance.now()

    console.log(`Call to doSomething took ${endTime - startTime} milliseconds`)
}

bootstrap()