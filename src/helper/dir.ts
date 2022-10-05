import * as deps from '../deps.deno.ts';

/**
 * Make project directory
 *
 * @param root project path
 * @param options options
 * @returns
 */
export function MakeDir(
  root: string,
  options = { recursive: true }
): Promise<string | undefined> {
  return deps.fs.promises.mkdir(root, options);
}

/**
 * Verify that the project directory is empty and conflict-free
 *
 * @param root project path
 * @param name project name
 * @returns
 */
export function IsFolderEmpty(root: string, name: string): boolean {
  const validFiles = [
    '.DS_Store',
    '.git',
    '.gitattributes',
    '.gitignore',
    '.gitlab-ci.yml',
    '.idea',
    '.npmignore',
    'LICENSE',
    'npm-debug.log',
    'yarn-debug.log',
    'yarn-error.log',
  ];

  const conflicts = deps.fs
    .readdirSync(root)
    .filter((file) => !validFiles.includes(file))
    .filter((file) => !/\.iml$/.test(file));

  if (conflicts.length > 0) {
    console.log(
      `The directory ${deps.chalk.green(name)} contains files that could conflict:`
    );
    console.log();

    for (const file of conflicts) {
      try {
        const stats = deps.fs.lstatSync(deps.path.join(root, file));
        if (stats.isDirectory()) {
          console.log(`  ${deps.chalk.blue(file)}/`);
        } else {
          console.log(`  ${file}`);
        }
      } catch {
        console.log(`  ${file}`);
      }
    }

    console.log();
    console.log(
      'Either try using a new directory name, or remove the files listed above.'
    );
    console.log();
    return false;
  }

  return true;
}