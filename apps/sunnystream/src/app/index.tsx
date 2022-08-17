import styled from '@emotion/styled'
import { Button } from "@coffeewithegg/common-react";

const StyledApp = styled.div`
  position: fixed;
  width: 100%;
  height: 100%;
`

export function App() {
  return (
    <StyledApp className="bg-bg01 flex items-center">
      <Button variant="outline" className="ml-[200px]">Create a room now</Button>
    </StyledApp>
  )
}

export default App
