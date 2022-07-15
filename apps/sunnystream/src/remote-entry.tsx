import React from 'react';
import ReactDOM from 'react-dom';
import App from './app/app';

class MfeSunnyStream extends HTMLElement {
  connectedCallback() {
    ReactDOM.render(<App />, this);
  }
}

customElements.define('app-sunnystream', MfeSunnyStream);
