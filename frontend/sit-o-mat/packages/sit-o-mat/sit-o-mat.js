import { LitElement, html, css } from 'lit-element';

import '@material/mwc-top-app-bar-fixed';
import '@material/mwc-drawer';
import '@material/mwc-icon-button';
import '@material/mwc-button';
import '@material/mwc-snackbar';


import '../sitomat-login/sitomat-login'
import '../sitomat-password/sitomat-password'
import '../sitomat-workplace/sitomat-workplace'
import '../sitomat-weekdays/sitomat-weekdays'
import * as notification from '../notification/notification';
import * as userAPI from '../api/user'
import * as workplaceAPI from '../api/workplace'


export class SitOMat extends LitElement {
  static get properties() {
    return {
      user: { type: Object },
      workplaces: { type: Array }
    };
  }

  constructor() {
    super();

    this.user = {
      WeeklyRequests: {},
    };
    this.workplaces = [];

    window.addEventListener('sitomat-notify', this.onNotify.bind(this));

    this._loadWorkplaces();
    setInterval(this._loadWorkplaces,5 * 60 * 60 * 1000) //Every 5h
  }

  _loadWorkplaces() {
    workplaceAPI.getWorkplaces()
      .then(workplaces => {
        this.workplaces = workplaces.sort((a, b) => {
          if (a.Location > b.Location) { return -1 }
          if (a.Location < b.Location) { return 1 }
          if (a.Name > b.Name) { return 1 }
          if (a.Name < b.Name) { return -1 }
          return 0
        });
      }).catch(err => notification.error(err))
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



  _groupByLocation(workplaces) {
    return workplaces.reduce(function (r, a) {
      r[a.Location] = r[a.Location] || [];
      r[a.Location].push(a);
      return r;
    }, {});
  }

  renderWorkplace(w) {
    return html`<sitomat-workplace .workplace=${w}></sitomat-workplace>`
  }

  renderWorkplaces(workplaces) {
    var locationMap = this._groupByLocation(workplaces)
    return Object.entries(locationMap).map(([key, workplaces]) => {
      return html`
              <div class="location">${key}</div>
              ${workplaces.map(this.renderWorkplace)}
          `;
    })
  }

  _myWorkplaceString(workplaces) {
    const myworkplace = this.workplaces.find(w => w.CurrentOwner == this.user.Name);
    if (!myworkplace) {
      return "kein Raum";
    }
    return myworkplace.Location + " " + myworkplace.Name;
  }

  render() {
    return html`
    <header id="topbar">

        <div id="title">Sit-o-Mat</div>
        
          <div id="spacer">
          ${this.user.Name ? html`
            <span>heute für dich:</span> 
            ${this._myWorkplaceString(this.workplaces)}
          `: null}
          </div>
        <sitomat-login id="login" @sitomat-login=${this.onLogin}></sitomat-login>
        <sitomat-password id="password"></sitomat-password>
        <span>${this.user.Name}</span>

      ${!this.user.Name
        ? html`
            <mwc-icon-button 
              icon="account_circle" title="Log in" slot="actionItems"
              @click=${e => this.shadowRoot.getElementById('login').open()}>
            </mwc-icon-button>`
        : html`
          <mwc-icon-button 
            icon="vpn_key" title="Change Password" slot="actionItems"
            @click=${e => this.shadowRoot.getElementById('password').open()}>
          </mwc-icon-button>
          <mwc-icon-button 
              icon="logout" title="Logout" slot="actionItems"
              @click=${e => this.logout()}>
          </mwc-icon-button>`
      }

        
        
    </header>

    <div id="page">

      ${this.user.Name ? html`
        <sitomat-weekdays 
            .weekdays=${this.user.WeeklyRequests} 
            @sitomat-change-weeklyrequests=${e => this.onChangeWeeklyRequests(e)}>
        </sitomat-weekdays>` : null
      }


        ${this.renderWorkplaces(this.workplaces)}
        
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
          --mdc-theme-secondary: black;
    }
    #page {
      max-width: 768px;
      margin: 0 auto;
      padding: 24px 16px 16px 16px;
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

      position: sticky;
      top: 0;
      z-index: 1;

    }

    #topbar #title {
      font-size: 24px;
      font-family: Pacifico, cursive;

    }

    #topbar #spacer {
      flex-grow: 1;
      padding: 16px;
      font-family: cursive;
    }

    #topbar #spacer span {
      font-family: Pacifico, cursive;
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

    sitomat-weekdays {
      margin-bottom: 32px;
    }

    .location {
      padding: 16px 16px 4px 16px;
      font-size: 18px;
      font-weight: lighter;
    }
    `,


    ];
  }
}


window.customElements.define('sit-o-mat', SitOMat);
