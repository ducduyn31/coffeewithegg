import React, { SVGProps, useEffect, useRef, useState } from 'react'

interface IconProps {
  name: string
  fill?: string
  stroke?: string
}

const Icon: React.FC<IconProps & { [k: string]: unknown }> = ({
  name,
  fill,
  stroke,
  ...rest
}) => {
  const ImportedIconRef = useRef<React.FC<SVGProps<SVGSVGElement>> | never>()

  const [loading, setLoading] = useState(false)

  useEffect(() => {
    let isSubscribed = true
    setLoading(true)
    const importIcon = async (): Promise<void> => {
      try {
        ImportedIconRef.current = (
          await import(
            `!!@svgr/webpack!../../../../base/assets/icons/${name}.svg`
          )
        ).default
      } finally {
        if (isSubscribed) setLoading(false)
      }
    }
    importIcon()
    return () => {
      isSubscribed = false
    }
  }, [name])

  const ImportedIcon = ImportedIconRef.current

  if (loading || !ImportedIcon) return null

  return <ImportedIcon fill={fill} stroke={stroke} {...rest} />
}

export default Icon
