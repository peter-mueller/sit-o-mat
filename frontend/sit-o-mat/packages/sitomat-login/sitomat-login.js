import { LitElement, html, css } from 'lit-element';

import '@material/mwc-dialog';
import '@material/mwc-button';
import '@material/mwc-textfield';

import * as userAPI from '../api/user';
import * as notification from '../notification/notification'

export class SitomatLogin extends LitElement {
    static get properties() {
        return {
            username: { type: String },
            password: { type: String },
        }
    }

    constructor() {
        super()

        const auth = userAPI.authentication;
        if (auth.username && auth.password) {
            this.login(auth.username, auth.password);
        }
    }


    open() {
        this.shadowRoot.getElementById('dialog').open = true;
    }
    close() {
        this.shadowRoot.getElementById('dialog').open = false;
    }

    render() {
        return html`
        <mwc-dialog title="Login" id="dialog">
            <mwc-textfield outlined
                id="username"
                required
                label="username"
                dialogInitialFocus>
            </mwc-textfield>
            <mwc-textfield outlined
            required
                type="password"
                id="password"
                label="password"
                dialogInitialFocus>
            </mwc-textfield>

            <mwc-button slot="primaryAction"
            @click=${e => {
                const username = this.shadowRoot.getElementById('username');
                const password = this.shadowRoot.getElementById('password');
                if (!username.reportValidity()) {
                    return;
                }
                if (!password.reportValidity()) {
                    return;
                }

                return this.login(username.value, password.value).then(
                    ok => {
                        if (ok) {
                            this.close();
                        }
                    }
                )
            }}>
                Login
            </mwc-button>
            <mwc-button slot="secondaryAction" dialogAction="close">
                Cancel
            </mwc-button>
        </mwc-dialog>
        `;
    }

    login(username, password) {
        const auth = userAPI.authentication;

        return userAPI.getUser(username, password)
            .then(user => {
                auth.username = username;
                auth.password = password;

                const event = new CustomEvent('sitomat-login', { detail: user });
                this.dispatchEvent(event);

                return true;
            }).catch(err => {
                if (err.status == 401 || err.status == 404) {
                    notification.warning("Benutzername/Passwort Kombination nicht bekannt, bitte erneut versuchen");

                    auth.clear();

                    return;
                }
                throw err;
            });
    }

    static get styles() {
        return [css`
            mwc-textfield {
                display: block;
                padding: 8px 0;
            }
        `]
    }
}

window.customElements.define('sitomat-login', SitomatLogin)
