//import { createDefaultConfig } from '@open-wc/building-rollup';
//export default createDefaultConfig({ input: './index.html' });

// if you need to support IE11 use "modern-and-legacy-config" instead.
import { createCompatibilityConfig } from '@open-wc/building-rollup';
const cpy = require('rollup-plugin-cpy');

const config =  createCompatibilityConfig({ input: './index.html' });

module.exports = [
  {
    ...config[0],
    plugins: [
      ...config[0].plugins,
      cpy({
        files: ['manifest.json', 'favicon.ico', 'sit-o-mat.png'],
        dest: 'dist',
        options: {
          parents: true,
        },
      }),
    ],
  },
  config[1],
];
