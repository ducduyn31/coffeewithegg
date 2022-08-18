import React from 'react'
import ReactDOM from 'react-dom'
import Index from './app'

class MfeSunnyStream extends HTMLElement {
  connectedCallback() {
    ReactDOM.render(<Index />, this)
  }
}

customElements.define('app-sunnystream', MfeSunnyStream)
