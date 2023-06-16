import { createRoot } from "react-dom/client";
import Index from "./app";

class MfeSunnyStream extends HTMLElement {
  connectedCallback() {
    const root = createRoot(this)
    root.render(<Index />)
  }
}

customElements.define('app-sunnystream', MfeSunnyStream)
