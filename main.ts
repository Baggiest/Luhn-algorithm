/* eslint-disable @typescript-eslint/no-unused-vars */

import evaluate from './algorithm'
const cc: string = '4539-4794-2564-9801' // 4539479425649801 to test valid

async function bootstrap() {
    let result = await evaluate(cc)
    console.log(result)
}

bootstrap()