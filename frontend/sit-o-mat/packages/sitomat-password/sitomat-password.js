import { LitElement, html, css } from 'lit-element';

import '@material/mwc-dialog';
import '@material/mwc-button';
import '@material/mwc-textfield';

import * as userAPI from '../api/user';
import * as notification from '../notification/notification'

export class SitomatPassword extends LitElement {
    static get properties() {
        return {
            password: { type: String },
        }
    }

    constructor() {
        super()
    }


    open() {
        this.shadowRoot.getElementById('dialog').open = true;
    }
    close() {
        this.shadowRoot.getElementById('dialog').open = false;
    }

    render() {
        return html`
        <mwc-dialog title="Change Password" id="dialog">
            <mwc-textfield outlined
            required
                type="password"
                id="password"
                label="password"
                dialogInitialFocus>
            </mwc-textfield>

            <mwc-button slot="primaryAction"
            @click=${e => {
                const password = this.shadowRoot.getElementById('password');
                if (!password.reportValidity()) {
                    return;
                }

                this.changePassword(password.value).then(
                    ok => {
                        if (ok) {
                            this.close();
                        }
                    }
                )
            }}>
                Change
            </mwc-button>
            <mwc-button slot="secondaryAction" dialogAction="close">
                Cancel
            </mwc-button>
        </mwc-dialog>
        `;
    }

  changePassword(password) {
        const auth = userAPI.authentication;

        return userAPI.changePassword(password)
            .then(user => {
                auth.password = password;

                return true;
            }).catch(err => {
            if (err.status == 401 || err.status == 404) {
              notification.warning("Passwort konnte nicht geändert werden.");

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

window.customElements.define('sitomat-password', SitomatPassword)
