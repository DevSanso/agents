function versionDisplay(list : Array<String>,displayLen : number) {
    console.log("*" + ' '.repeat(38) + "*");
    list.forEach((value) => {
        console.log(`*${' '.repeat(4)}- ${value.padEnd(displayLen - 8)}*`)
    });
}

export function _versionListDisplay(list : Array<String>) {
    console.log("*".repeat(40));
    console.log("*" + " - command-list:".padEnd(38) + "*");
    console.log("*" + " ".repeat(38) + "*");
    versionDisplay(list, 40);
    console.log("*".repeat(40));
}