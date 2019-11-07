import { LitElement, html, css } from 'lit-element';

import '@material/mwc-dialog';
import '@material/mwc-button';
import '@material/mwc-textfield';

export class SitomatLogin extends LitElement {
    static get properties() {
        return {
            username: { type: String },
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

                fetch()

                this.close()
            }}>
                Login
            </mwc-button>
            <mwc-button slot="secondaryAction" dialogAction="close">
                Cancel
            </mwc-button>
        </mwc-dialog>
        `;
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