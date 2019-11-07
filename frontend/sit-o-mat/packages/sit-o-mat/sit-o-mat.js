import { LitElement, html, css } from 'lit-element';

import '@material/mwc-top-app-bar-fixed';
import '@material/mwc-drawer';
import '@material/mwc-icon-button';
import '@material/mwc-button';
import '@material/mwc-snackbar';


import '../sitomat-login/sitomat-login'
import '../sitomat-workplace/sitomat-workplace'
import '../sitomat-weekdays/sitomat-weekdays'
import * as notification from '../notification/notification';
import * as userAPI from '../api/user'


export class SitOMat extends LitElement {
  static get properties() {
    return {
      user: { type: Object }
    };
  }

  constructor() {
    super();

    this.user = {
      WeeklyRequests: {},
    };

    window.addEventListener('sitomat-notify', this.onNotify.bind(this));
  }

  onNotify(e) {
    const message = e.detail.message;
    const level = e.detail.level;
    const snackbarInfo = this.shadowRoot.getElementById('snackbarInfo');
    const snackbarWarning = this.shadowRoot.getElementById('snackbarWarning');
    const snackbarError = this.shadowRoot.getElementById('snackbarError');

    switch (level) {
      case notification.LEVEL_INFO:
        snackbarInfo.labelText = message;
        snackbarInfo.open();
      case notification.LEVEL_WARNING:
        snackbarWarning.labelText = message;
        snackbarWarning.open();
      case notification.LEVEL_ERROR:
        snackbarError.labelText = message;
        snackbarError.open();
    }
  }

  onLogin(e) {
    const user = e.detail;
    this.user = user;
  }
  onChangeWeeklyRequests(e) {
    const requests = e.detail;
    userAPI.patchWeeklyRequests(requests)
      .then(u => {
        this.user.WeeklyRequests = u.WeeklyRequests
      }).catch(err => {
        notification.error(err);
      });
  }

  logout() {
    userAPI.authentication.clear();
    this.user = {
      WeeklyRequests: {},
    };
  }

  render() {
    return html`
    <header id="topbar">

        <div id="title" id="title">Sit-o-Mat</div>
        <sitomat-login id="login" @sitomat-login=${this.onLogin}></sitomat-login>
        <span>${this.user.Name}</span>

      ${!this.user.Name
        ? html`
            <mwc-icon-button 
              icon="account_circle" title="Log in" slot="actionItems"
              @click=${e => this.shadowRoot.getElementById('login').open()}>
            </mwc-icon-button>`
        : html`
          <mwc-icon-button 
              icon="logout" title="Logout" slot="actionItems"
              @click=${e => this.logout()}>
          </mwc-icon-button>`
      }

        
        
    </header>

    <div id="page">
        <sitomat-weekdays 
            .weekdays=${this.user.WeeklyRequests} 
            @sitomat-change-weeklyrequests=${e => this.onChangeWeeklyRequests(e)}>
        </sitomat-weekdays>
        <sitomat-workplace></sitomat-workplace >
    </div>

    <mwc-snackbar leading id="snackbarInfo"></mwc-snackbar>
    <mwc-snackbar leading id="snackbarWarning"></mwc-snackbar>
    <mwc-snackbar leading id="snackbarError"></mwc-snackbar>

    `}

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
