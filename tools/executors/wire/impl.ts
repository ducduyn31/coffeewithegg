import { ExecutorContext } from '@nrwl/devkit'
import { promisify } from 'util'
import { exec } from 'child_process'

export interface WireGenOptions {}

export default async function wireExecutor(
  options: WireGenOptions,
  context: ExecutorContext,
) {
  const appCwd = `${context.workspace.projects[context.projectName].root}`
  console.info(`Executing "wire" at ${appCwd}`)

  const { stdout, stderr } = await promisify(exec)(
    `go run github.com/google/wire/cmd/wire`,
    { cwd: appCwd },
  )
  console.log(stdout)
  console.error(stderr)

  const success = !stderr
  return { success }
}
