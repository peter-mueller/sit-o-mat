import { html, css, LitElement } from 'lit-element';

import '../../sitomat-workplace/sitomat-workplace.js';


export class PageMain extends LitElement {
  static get styles() {
    return css`
      :host {
      }

      .group {
        font-size: 24px;
        margin: 8px 16px;
        font-weight: lighter;
      }
    `;
  }

  static get properties() {
    return {
      title: { type: String },
      logo: { type: Function },
    };
  }

  constructor() {
    super();
  }

  render() {
    return html`

    <div class="group"> Raum 3.12</div>
      <sitomat-workplace .workplace=${{
        Name: "BluePanda",
        Location: "Raum 3.12",
        Ranking: 1,
        CurrentOwner: "p.mueller"
      }}></sitomat-workplace >
      <sitomat-workplace .workplace=${{
        Name: "RedWahu",
        Location: "Raum 3.12",
        Ranking: 1,
        CurrentOwner: "bob.bob"
      }}></sitomat-workplace >
    `;
  }
}
