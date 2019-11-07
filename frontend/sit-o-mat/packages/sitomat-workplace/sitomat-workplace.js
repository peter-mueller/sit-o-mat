import { LitElement, html, css } from 'lit-element';


export class SitomatWorkplace extends LitElement {
    static get properties() {
        return {
            workplace: { type: Object }
        };
    }

    constructor() {
        super();
        this.workplace = {};
    }


    render() {
        return html`
            <div id="main">
                <div id="title">${this.workplace.Name}</div>
                <div id="subtitle">${this.workplace.Location} - Rank ${this.workplace.Ranking}</div>
            </div>
            <div id="owner">${this.workplace.CurrentOwner}</div>
        `;
    }

    static get styles() {
        return [
            css`
        :host {
            display: flex;
            flex-direction: row;
            align-items: center;
            background-color: white;
            border-radius: 2px;
            padding: 8px 16px;
        }

        #main {
            flex-grow: 1;
            display: flex;
            flex-direction: column;
            align-items: start;
        }

        #subtitle {
            font-size: 12px;
            color: var(--mdc-dialog-content-ink-color, rgba(0, 0, 0, 0.6))
        }

        #owner {

        }
        
      `,


        ];
    }
}

window.customElements.define('sitomat-workplace', SitomatWorkplace);
