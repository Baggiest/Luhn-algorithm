/* eslint-disable @typescript-eslint/no-unused-vars */

export default async function evaluate(cc: string) {

    // just making 1234-5134-1424 to 1234515 blah blah

    const ccArray = cc.split("");

    let ccArrayNumber = ccArray.map((ccArray) => Number(ccArray));
    let finalSum: number = 0;
    let i: number;

    for (i = 0; i < ccArrayNumber.length; i++) {

        let num = ccArrayNumber[i];

        if (i % 2 === 0) {
            evenIndex(num)
        }

        else {
            oddIndex(num)
        }
    }

    function evenIndex(num: number) {

        let digits = num * 2;
        let firstDigit = digits % 10
        let secondDigit = Math.floor(digits / 10)

        finalSum += firstDigit + secondDigit
    }

    function oddIndex(num: number) {
        finalSum += num
    }

    // console.log(finalSum)

    async function determineResult() {

        if (finalSum % 10 === 0) {
            console.log("It is a valid cc number !")
            return true;
        }
        else {
            console.log("its not a valid cc number")
            return false;
        }
    }

    return (await determineResult() ? true : false)
}

