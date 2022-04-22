import { exec } from 'child_process';

export async function simpleExample(
  dir: string,
): Promise<void> {
  const cmd = `ls -la`;
  exec(cmd, (error, stdout, stderr) => {
    console.log(`ERROR: ${error}`);
    console.log(`STDERR: ${stderr}`);
    console.log(`STDOUT: ${stdout}`);
  });
}

simpleExample();
