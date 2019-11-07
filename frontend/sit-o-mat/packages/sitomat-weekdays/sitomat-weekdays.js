import { LitElement, html, css } from 'lit-element';

import '@material/mwc-checkbox';


export class SitomatWeekdays extends LitElement {
    static get properties() {
        return {
            weekdays: { type: Object }
        };
    }

    constructor() {
        super();
        this.weekdays = {
            Montag: false,
            Dienstag: false,
            Mittwoch: false,
            Donnerstag: false,
            Freitag: false
        };
    }


    render() {
        return html`

<div id="header">MEINE ANWESENHEIT</div>

<div id="container" @change=${e => this.onChange(e)}>
        <div class="column">
            <div>Montag</div>
            <mwc-checkbox id="Montag" .checked=${this.weekdays.Montag}></mwc-checkbox>
        </div>
        <div class="column">
            <div>Dienstag</div>
            <mwc-checkbox id="Dienstag" .checked=${this.weekdays.Dienstag}></mwc-checkbox>
        </div>
        <div class="column">
            <div>Mittwoch</div>
            <mwc-checkbox id="Mittwoch" .checked=${this.weekdays.Mittwoch}></mwc-checkbox>
        </div>
        <div class="column">
            <div>Donnerstag</div>
            <mwc-checkbox id="Donnerstag" .checked=${this.weekdays.Donnerstag}></mwc-checkbox>
        </div>
        <div class="column">
            <div>Freitag</div>
            <mwc-checkbox id="Freitag" .checked=${this.weekdays.Freitag}></mwc-checkbox>
        </div>
</div>

        `;
    }

    onChange(e) {

        const root = this.shadowRoot;
        const checkboxValue = function (day) {
            return root.getElementById(day).checked;
        }

        const event = new CustomEvent('sitomat-change-weeklyrequests', {
            detail: {
                Montag: checkboxValue('Montag'),
                Dienstag: checkboxValue('Dienstag'),
                Mittwoch: checkboxValue('Mittwoch'),
                Donnerstag: checkboxValue('Donnerstag'),
                Freitag: checkboxValue('Freitag'),
            }
        });

        this.dispatchEvent(event);
    }

    static get styles() {
        return [
            css`

            :host {
                display: block;
                padding: 8px 16px 32px 16px;


                background-color: white;
                border-radius: 2px;
                font-family: monospace;

            }

            #header {
                padding: 8px 16px;
                font-family: monospace;
                text-align: center;

            }

            #container {
                display: flex;
                flex-direction: row;
                justify-content: space-around;
                flex-wrap: wrap;
            
            }

            .column {
                display: flex;
                flex-direction: column;
                flex-grow: 1;
                align-items: center;

            }

            .column div {
                width: 100%;
                text-align: center;
                padding: 8px 4px;
                border-bottom: 1px solid black;
            }
            `,
        ];
    }
}

window.customElements.define('sitomat-weekdays', SitomatWeekdays);
