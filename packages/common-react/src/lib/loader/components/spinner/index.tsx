import React from 'react'
import Icon from '../../../icon'
import styled from '@emotion/styled'

type SpinnerType = 'default'

export interface SpinnerProps {
  variant?: SpinnerType
}

const SpinnerContainer = styled.div`
  position: relative;
  height: 100vh;
  width: 100vw;
`

const Spinner: React.FC<SpinnerProps> = ({ variant = 'default', ...props }) => (
  <SpinnerContainer {...props}>
    <Icon
      name="spinner"
      fill="transparent"
      className="animate-spin absolute inset-0 m-0"
    />
  </SpinnerContainer>
)

export default Spinner
