describe('common-angular', () => {
  beforeEach(() => cy.visit('/iframe.html?id=cardcomponent--primary'))
  it('should render the component', () => {
    cy.get('coffeewithegg-card').should('exist')
  })
})
