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
    if(process.argv.length <= 0) {
        displays.mainDisplay();
        displays.versionListDisplay(support_version);
    }
    else {
        asyncRunCommand(process.argv[1], process.argv[2]);
    }
};

main();