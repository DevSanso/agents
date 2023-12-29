import displays from './console/console';
import runCommand from './command/command';

const support_version = [
    "install",
    "uninstall"
]

const asyncRunCommand = async (cmd : string, configPath : string) => {
    await runCommand(cmd, configPath);
};
const main = () => {
    if(process.argv.length <= 2) {
        displays.mainDisplay();
        displays.versionListDisplay(support_version);
    }
    else {
        asyncRunCommand(process.argv[2], process.argv[3])
        .then(()=> {
            console.log("command is done!");
        })
        .catch((value) => {
            console.error(`initDB run error : ${value}`);
        });
    }
};

main();