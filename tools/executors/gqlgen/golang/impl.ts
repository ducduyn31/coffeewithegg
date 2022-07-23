import { ExecutorContext } from '@nrwl/devkit'
import { promisify } from 'util'
import { exec } from 'child_process'

export interface GqlGenOptions {
  config: string
}

export default async function gqlgenExecutor(
  options: GqlGenOptions,
  context: ExecutorContext,
) {
  const appCwd = `${context.workspace.projects[context.projectName].root}`
  console.info(`Executing "gqlgen generate" at ${appCwd}`)
  const verbose = context.isVerbose ? '--verbose' : ''
  if (verbose && verbose.length > 0) console.info(`VERBOSE: true`)
  console.info(`Options: ${JSON.stringify(options, null, 2)}`)

  const customConfig = options.config ? `--config=${options.config}` : ''

  const { stdout, stderr } = await promisify(exec)(
    `go run github.com/99designs/gqlgen generate ${customConfig} ${verbose}`,
    { cwd: appCwd },
  )
  console.log(stdout)
  console.error(stderr)

  const success = !stderr
  return { success }
}
