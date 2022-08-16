import { render } from '@testing-library/react'

import Index from './app'

describe('App', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<Index />)

    expect(baseElement).toBeTruthy()
  })

  it('should have a greeting as the title', () => {
    const { getByText } = render(<Index />)

    expect(getByText(/Welcome sunnystream/gi)).toBeTruthy()
  })
})
