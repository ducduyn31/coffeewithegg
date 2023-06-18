import { Component } from '@angular/core'
import { AuthService } from '@auth0/auth0-angular'

@Component({
  selector: 'coffeewithegg-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent {
  constructor(private auth: AuthService) {}

  login() {
    const currentUrl = window.location.href
    this.auth.loginWithRedirect({
      authorizationParams: {
        redirect_uri: `${currentUrl}`,
      },
    })
  }
}
