import { defineRouteMiddleware } from '@astrojs/starlight/route-data';

const ARCHIVED_VERSION_PATH = /^\/(?:0\.7|0\.8)(?:\/|$)/;

export const onRequest = defineRouteMiddleware((context) => {
  const { pathname } = context.url;

  if (pathname === '/') {
    context.locals.starlightRoute.head.push({
      tag: 'script',
      attrs: {
        type: 'application/ld+json',
      },
      content: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'WebSite',
        name: 'Valgo',
        alternateName: ['Valgo Go validation library', 'Valgo validation library'],
        url: 'https://valgo.build/',
      }),
    });
  }

  if (!ARCHIVED_VERSION_PATH.test(pathname)) return;

  const { head } = context.locals.starlightRoute;
  const currentPathname =
    pathname.replace(/^\/(?:0\.7|0\.8)(?=\/|$)/, '').replace('/validators/or-operator/', '/validators/or-operators/') ||
    '/';
  const currentUrl = new URL(currentPathname, context.site ?? context.url.origin);

  context.locals.starlightRoute.head = [
    ...head.filter(
      ({ tag, attrs }) =>
        !(tag === 'meta' && attrs?.name === 'robots') &&
        !(tag === 'link' && attrs?.rel === 'canonical')
    ),
    {
      tag: 'meta',
      attrs: {
        name: 'robots',
        content: 'noindex, follow',
      },
    },
    {
      tag: 'link',
      attrs: {
        rel: 'canonical',
        href: currentUrl.href,
      },
    },
  ];
});
