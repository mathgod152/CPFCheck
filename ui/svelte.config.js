import adapter from '@sveltejs/adapter-static';
import sveltePreprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    preprocess: sveltePreprocess(),
    kit: {
        adapter: adapter({
            pages: '../server/public',
            assets: '../server/public',
            fallback: "index.html",
        }),
    },
};

export default config;
