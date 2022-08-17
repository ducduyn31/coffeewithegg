import React from 'react'
import Icon from '../../../icon'
import styled from '@emotion/styled'

type SpinnerType = 'default' | 'button'

export interface SpinnerProps {
  variant?: SpinnerType
}

const variantStyles: Record<SpinnerType, string> = {
  default: 'relative h-screen w-screen bg-bg01',
  button: 'w-full',
}

const SpinnerContainer = styled.div``

const Spinner: React.FC<SpinnerProps> = ({ variant = 'default', ...props }) => (
  <SpinnerContainer className={variantStyles[variant]} {...props}>
    <Icon
      name="spinner"
      fill="transparent"
      className="absolute animate-spin inset-0 m-auto"
    />
  </SpinnerContainer>
)

export default Spinner
