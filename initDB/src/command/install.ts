import fs from 'node:fs';
import path from 'node:path';

import pg from 'pg';

import {AssetsPathDir} from '../asset/path_dir';

export interface InstallConfig {
    dbConfig : {
        host : string,
        port : number,
        username : string,
        password : string,
        dbname : string
    }
}

const collectionDir : AssetsPathDir = new AssetsPathDir(process.execPath, "1.0","collection");
const spmsDir : AssetsPathDir = new AssetsPathDir(process.execPath, "1.0","spms");
const syncDir : AssetsPathDir = new AssetsPathDir(process.execPath, "1.0","sync");
const webDir : AssetsPathDir = new AssetsPathDir(process.execPath, "1.0","web");

const pathLessFunc = (l : string, r : string) : boolean => {
    const l_prefix = l[0];
    const r_prefix = r[0];

    return parseInt(l_prefix, 10) < parseInt(r_prefix, 10);
}

const runQueryFiles = async (client : pg.Client, dir : string, fileList : Array<string>) => {
    await fileList.map(async (value) => {
        const query = fs.readFileSync(path.join(dir, value));
        await client.query(query.toString());
    })
}

export const InstallCommand = async (cfg : InstallConfig) => {
    const client = new pg.Client();

    await Promise.all([
        runQueryFiles(client, collectionDir.dir(), collectionDir.getFileSortPaths(pathLessFunc)),
        runQueryFiles(client, spmsDir.dir(), spmsDir.getFileSortPaths(pathLessFunc)),
        runQueryFiles(client, syncDir.dir(), syncDir.getFileSortPaths(pathLessFunc)),
        runQueryFiles(client, webDir.dir(), webDir.getFileSortPaths(pathLessFunc))
    ]);
    
}

