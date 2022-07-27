import { ExecutorContext } from '@nrwl/devkit'
import { promisify } from 'util'
import { exec } from 'child_process'
import * as path from 'path'

export interface GinkgoOptions {
  ci: boolean
  ignoreExtra: boolean
  _: string[]
}

export default async function ginkgoExecutor(
  options: GinkgoOptions,
  context: ExecutorContext,
) {
  const atPath =
    options._?.find((op) => op.startsWith('at='))?.substring(3) ?? ''
  const appCwd = path.join(
    `${context.workspace.projects[context.projectName].root}`,
    atPath,
  )
  const extraOptions =
    !options?.ignoreExtra && options._ ? options._?.join(' ') : ''

  console.info(`Executing "ginkgo ${extraOptions}" at ${appCwd}`)

  console.info(`Options: ${JSON.stringify(options, null, 2)}`)

  const isRecursive = options.ci ? '-r -v -p' : ''

  const { stdout, stderr } = await promisify(exec)(
    `go run github.com/onsi/ginkgo/v2/ginkgo ${isRecursive} ${extraOptions}`,
    { cwd: appCwd },
  )
  console.log(stdout)
  console.error(stderr)

  const success = !stderr
  return { success }
}
