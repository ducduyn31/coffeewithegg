import React, { Suspense } from 'react'
import { History } from 'history'
import { Loader } from '@coffeewithegg/common-react'

interface ProviderWrapperProps {
  children?: React.ReactNode
  history?: History
}

const ProviderWrapper: React.FC<ProviderWrapperProps> = ({ children }) => (
  <Suspense fallback={<Loader.Spinner />}>{children}</Suspense>
)

export default ProviderWrapper
