import { defineConfig } from 'astro/config';
import sitemap from '@astrojs/sitemap';
import starlight from '@astrojs/starlight';
import starlightVersions from 'starlight-versions';
import tailwindcss from '@tailwindcss/vite';
import { docsSidebar } from './src/data/sidebar.mjs';
import { isArchivedVersionUrl, starlightSeo } from './src/plugins/starlight-seo.mjs';

export default defineConfig({
  site: 'https://valgo.build',
  redirects: {
    '/0.8': '/0.8/getting-started/',
    '/0.7': '/0.7/getting-started/',
  },
  integrations: [
    sitemap({
      filter: (page) => !isArchivedVersionUrl(page),
    }),
    starlight({
      title: 'Valgo',
      description: 'Valgo is a type-safe Go validation library for validating values without struct tags.',
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
      head: [
        {
          tag: 'meta',
          attrs: {
            property: 'og:image',
            content: 'https://valgo.build/og-image.png',
          },
        },
        {
          tag: 'meta',
          attrs: {
            property: 'og:image:alt',
            content: 'Valgo logo on a white background',
          },
        },
        {
          tag: 'meta',
          attrs: {
            name: 'twitter:image',
            content: 'https://valgo.build/og-image.png',
          },
        },
      ],
      plugins: [
        starlightSeo(),
        starlightVersions({
          // When you run the site, the plugin will archive the current docs state
          // under the first configured version slug.
          versions: [
            { slug: '0.8', label: 'v0.8' },
            { slug: '0.7', label: 'v0.7' },
          ],
          current: { label: 'v0.8.1 Latest' },
        }),
      ],
      sidebar: docsSidebar,
    }),
  ],

  vite: {
    plugins: [tailwindcss()],
  },
});
