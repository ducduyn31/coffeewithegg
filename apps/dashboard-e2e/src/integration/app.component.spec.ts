describe('dashboard', () => {
  beforeEach(() => cy.visit('/iframe.html?id=appcomponent--primary'));
  it('should render the component', () => {
    cy.get('coffeewithegg-root').should('exist');
  });
});