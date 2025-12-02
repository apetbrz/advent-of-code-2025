import fs from 'fs';

const isSilly = (number) => {
    let n = number.toString()
    return n.substring(0, n.length / 2) == n.substring(n.length / 2)
}

fs.readFile(import.meta.dirname + '/sequence.txt', (err, data) => {
    if (err) {
        console.error(err);
        return;
    }

    let ranges = data.toString().split(",");
    let sum = 0;

    for (let r of ranges) {
        let range = r.trim()
        let bounds = range.split("-")
        let hi = parseInt(bounds[1])
        let lo = parseInt(bounds[0])

        for (let i = lo; i <= hi; i++) {
            if (isSilly(i)) {
                sum += parseInt(i);
            }
        }
    }

    console.log(sum);
});
