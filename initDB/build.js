const esbuild = require('esbuild');
const fs = require('fs/promises');
const path = require('path');

const build = async () => {
    await esbuild.build({
        entryPoints : ['./src/main.ts'],
        bundle : true,
        outfile : './dist/initdb.js',
        platform : 'node'
    });
};

const copyAssets = async () => {
    const root = process.cwd();
    const src = path.join(root, "assets");
    const dst = path.join(root, "dist", "assets");
    
    await fs.rm(dst, {recursive : true});
    await fs.cp(src, dst, {recursive : true});
};

const start = () => {
    console.log("build typescript codes(1/2)\n\n");
    build();
    console.log("copy assets directory(2/2)\n\n");
    copyAssets();
}

start();