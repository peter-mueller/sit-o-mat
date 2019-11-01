import { LitElement, html, css } from 'lit-element';
import { classMap } from 'lit-html/directives/class-map.js';

import '@material/mwc-top-app-bar-fixed';
import '@material/mwc-drawer';
import '@material/mwc-icon-button';
import '@material/mwc-button';

import '../../page-main/page-main.js';
import '../../page-one/page-one.js';
import { templateAbout } from './templateAbout.js';

export class SitOMat extends LitElement {
  static get properties() {
    return {
      title: { type: String },
      page: { type: String },
      drawerShowInitial: { type: Boolean },
      drawerType: { type: Boolean },
      drawerWasOpen: { type: Boolean }
    };
  }

  constructor() {
    super();
    this.page = 'main';

    const mobile = window.innerWidth <= 760;
    this.drawerType = mobile ? 'modal' : 'dismissible';
    this.drawerShowInitial = !mobile;

    window.addEventListener('resize', e => {
      const drawer = this.shadowRoot.getElementById('drawer');
      const mobile = window.innerWidth <= 760;
      this.drawerType = mobile ? 'modal' : 'dismissible';
      drawer.open = !mobile;
    })
  }

  _renderPage() {
    switch (this.page) {
      case 'main':
        return html`
          <page-main></page-main>
        `;
      case 'pageOne':
        return html`
          <page-one></page-one>
        `;
      case 'about':
        return templateAbout;
      default:
        return html`
          <p>Page not found try going to <a href="#main">Main</a></p>
        `;
    }
  }

  __clickPageLink(ev) {
    ev.preventDefault();
    this.page = ev.target.hash.substring(1);
  }

  __addActiveIf(page) {
    return classMap({ active: this.page === page });
  }

  _toggleDrawer(e) {
    const drawer = this.shadowRoot.getElementById('drawer');
    drawer.open = !drawer.open;
  }

  render() {
    return html`


<mwc-drawer hasHeader .type=${this.drawerType} .open=${this.drawerShowInitial} id="drawer">
    <span slot="title">Sit-o-Mat</span>
    <span slot="subtitle">KM83</span>
    
    <nav id="menu">
      <mwc-button>Arbeitsplätze</mwc-button>
      <mwc-button>Nutzer</mwc-button>
    </nav>
    
    <div slot="appContent">
      <header id="topbar">
        <mwc-icon-button icon="menu"
        @click=${this._toggleDrawer}></mwc-icon-button>

        <div id="title" id="title">Sit-o-Mat</div>

        <mwc-icon-button icon="account_circle" title="Log in" slot="actionItems"></mwc-icon-button>

    </header>


      <div id="page">
                ${this._renderPage()}
        </div>
    </div>
</mwc-drawer>
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
