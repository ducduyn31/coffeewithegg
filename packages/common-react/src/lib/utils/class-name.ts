export default function (
  ...classNames: (string | boolean | undefined | null | never)[]
): string {
  return classNames
    .filter((name) => typeof name === 'string')
    .map((name) => name as string)
    .reduce(
      (previousValue, currentValue) => `${previousValue} ${currentValue}`,
      '',
    )
}
