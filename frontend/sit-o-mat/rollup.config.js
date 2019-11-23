//import { createDefaultConfig } from '@open-wc/building-rollup';
//export default createDefaultConfig({ input: './index.html' });

// if you need to support IE11 use "modern-and-legacy-config" instead.
import { createCompatibilityConfig } from '@open-wc/building-rollup';
export default createCompatibilityConfig({ input: './index.html' });
