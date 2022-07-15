/* eslint-disable @typescript-eslint/no-unused-vars */

import evaluate from './luhn'
const cc: string = '4539479425649801' // 4539479425649801 to test valid

async function bootstrap() {
    let result = await evaluate(cc)
    console.log(result)
}

bootstrap()