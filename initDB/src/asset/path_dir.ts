import path from 'node:path';
import fs from 'node:fs';

const quickSort = function (arr : Array<string>, lessFunc : (l: string, r: string) => boolean) : Array<string> {
    if (arr.length <= 1) return arr;
  
    const pivot = arr[0];
    const left = [];
    const right = [];
  
    for (let i = 1; i < arr.length; i++) {
      if (lessFunc(arr[i], pivot)) left.push(arr[i]);
      else right.push(arr[i]);
    }
  
    const lSorted = quickSort(left, lessFunc);
    const rSorted = quickSort(right, lessFunc);
    return [...lSorted, pivot, ...rSorted];
  };

export class AssetsPathDir {
    private readonly root: string;
    private readonly version : string;
    private readonly prefix : string;

    public constructor(root: string, version: string, prefix: string) {
        this.root = root;
        this.version = version;
        this.prefix  = prefix;   
    }
    public dir = () : string => path.join(this.root, this.version, this.prefix);
    public getFilePaths = (): Array<string> => fs.readdirSync(path.join(this.root, this.version, this.prefix));

    public getFileSortPaths = (lessFunc: (l: string, r: string) => boolean) => 
        quickSort(fs.readdirSync(path.join(this.root, this.version, this.prefix)), lessFunc);

}
