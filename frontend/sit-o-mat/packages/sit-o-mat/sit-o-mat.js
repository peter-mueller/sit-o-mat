import { LitElement, html, css } from 'lit-element';

import '@material/mwc-top-app-bar-fixed';
import '@material/mwc-drawer';
import '@material/mwc-icon-button';
import '@material/mwc-button';

import '../sitomat-login/sitomat-login'
import '../sitomat-workplace/sitomat-workplace'

export class SitOMat extends LitElement {
    static get properties() {
        return {
        };
    }

    constructor() {
        super();
    }

    render() {
        return html`


      <header id="topbar">

        <div id="title" id="title">Sit-o-Mat</div>
        <sitomat-login id="login"></sitomat-login>
        <mwc-icon-button icon="account_circle" title="Log in" slot="actionItems"
        @click=${e => this.shadowRoot.getElementById('login').open()}></mwc-icon-button>

    </header>


      <div id="page">
                <sitomat-workplace></sitomat-workplace >
        </div>
    `
    }

    static get styles() {
        return [
            css`
        :host {

          --mdc-theme-primary: #880e4f;


        }
        #title {
          font-family: Pacifico, cursive;
        }
        #page {
          max-width: 768px;
          margin: 0 auto;
          padding: 24px 16px 0 16px;
        }


        @media(max-device-width: 768px) {
          padding: 16px 4px 0 4px;
        }


        #topbar {

          color: white;
          display: flex;
          flex-direction: row;
          height: 64px;
          padding: 0 16px;
          align-items: center;
          background-color: var(--mdc-theme-primary, blue);

        }

        #topbar #title {
          flex-grow: 1;
          font-size: 24px;
        }

        #drawer {
          display: block;
        }

        #menu {
          display: flex;
          flex-direction: column;
          padding-top: 24px;
        }

        #menu mwc-button {
          padding: 4px 8px;
        }
      `,


        ];
    }
}


window.customElements.define('sit-o-mat', SitOMat);
