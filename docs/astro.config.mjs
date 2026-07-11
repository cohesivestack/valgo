import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import starlightVersions from 'starlight-versions';
import tailwindcss from '@tailwindcss/vite';
import { docsSidebar } from './src/data/sidebar.mjs';

export default defineConfig({
  site: 'https://valgo.build',
  redirects: {
    '/0.7': '/0.7/getting-started/',
  },
  integrations: [
    starlight({
      title: 'Valgo',
      description: 'Type-safe, expressive, and extensible validator library for Go.',
      logo: {
        light: './public/valgo-logo.svg',
        dark: './public/valgo-logo-dark.svg',
        alt: 'Valgo',
        replacesTitle: true,
      },
      favicon: '/favicon.png',
      social: [
        {
          icon: 'github',
          label: 'GitHub',
          href: 'https://github.com/cohesivestack/valgo',
        },
      ],
      customCss: [
        "./src/styles/global.css",
      ],
      plugins: [
        starlightVersions({
          // When you run the site, the plugin will archive the current docs state
          // under the first configured version slug.
          versions: [{ slug: '0.7', label: 'v0.7' }],
          current: { label: 'v0.8 Latest' },
        }),
      ],
      sidebar: docsSidebar,
    }),
  ],

  vite: {
    plugins: [tailwindcss()],
  },
});
