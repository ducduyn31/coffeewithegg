describe('dashboard', () => {
  beforeEach(() =>
    cy.visit('/iframe.html?id=projectoverviewcomponent--primary'),
  )
  it('should render the component', () => {
    cy.get('coffeewithegg-project-overview').should('exist')
  })
})
