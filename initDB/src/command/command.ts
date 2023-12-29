import fs from 'fs/promises';

import installCommand,{InstallConfig} from './install';

const runCommand = async (cmd : string, configPath : string) => {
    const cfg_str : string = (await fs.readFile(configPath)).toString();
    switch(cmd){
        case "install":
            installCommand(JSON.parse(cfg_str));
            break;
        default:
            throw `command : ${cmd} not support`;
    }
};

export default runCommand;