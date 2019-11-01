import { html, css, LitElement } from 'lit-element';

import '../../sitomat-workplace/sitomat-workplace.js';

export class PageOne extends LitElement {
  static get styles() {
    return css`
      :host {
        --page-one-text-color: #000;

        display: block;
        padding: 25px;
        color: var(--page-one-text-color);
      }
    `;
  }

  static get properties() {
    return {
      title: { type: String },
      counter: { type: Number },
    };
  }

  constructor() {
    super();
    this.title = 'Hey there';
    this.counter = 5;
  }

  __increment() {
    this.counter += 1;
  }

  render() {
    return html`
      <h2>${this.title} Nr. ${this.counter}!</h2>
      <sitomat-workplace workplace=${
      {
        Name: "BluePanda",
        Location: "Raum 3.12",
        Ranking: 1,
        CurrentOwner: "p.mueller"
      }
      }></sitomat-workplace >
    <button @click=${ this.__increment}> increment</button >
      `;
  }
}
