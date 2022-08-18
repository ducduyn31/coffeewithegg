import React from 'react'
import { resolveClassNames } from '../utils'
import Spinner from '../loader/components/spinner'

interface ButtonProps
  extends React.DetailedHTMLProps<
    React.ButtonHTMLAttributes<HTMLButtonElement>,
    HTMLButtonElement
  > {
  children?: React.ReactElement | string
  variant?: ButtonVariant
  size?: ButtonSize
  selected?: boolean
  loading?: boolean
}

type ButtonVariant = 'default' | 'outline' | 'inline'
type ButtonSize = 'small' | 'medium' | 'large'

const variantStyles: Record<ButtonVariant, string> = {
  default: resolveClassNames(
    'bg-primary text-white shadow-button',
    'hover:bg-background disabled:bg-grey05 disabled:shadow-none',
  ),
  outline: resolveClassNames(
    'bg-transparent text-primary border border-primary border-opacity-20',
    'hover:bg-primary hover:bg-opacity-10 disabled:bg-transparent',
  ),
  inline: 'text-primary text-ps underline',
}

const sizeStyles: Record<ButtonSize, string> = {
  small: 'h-36 text-p rounded-[5px]',
  medium: 'h-46 text-18 rounded-8',
  large: 'h-64 text-pl rounded-8',
}

const Button: React.FC<ButtonProps> = ({
  variant = 'default',
  size = 'medium',
  selected,
  loading,
  children,
  className,
  ...props
}) => (
  <button
    className={resolveClassNames(
      'relative font-medium px-24 disabled:text-grey03',
      'focus:outline-none focus:ring-transparent',
      'focus-visible:outline-primary',
      variantStyles[variant],
      sizeStyles[size],
      selected && 'bg-foreground text-white hover:bg-foreground',
      className,
    )}
    type="button"
    {...props}
  >
    <div className={resolveClassNames(loading && 'invisible')}>{children}</div>
    {loading && <Spinner variant="button" />}
  </button>
)

export default Button
