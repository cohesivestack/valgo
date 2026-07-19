const ARCHIVED_VERSION_PATH = /^\/(?:0\.7|0\.8)(?:\/|$)/;

/**
 * Adds route-level SEO metadata that depends on the final Starlight route.
 */
export function starlightSeo() {
  return {
    name: 'valgo-starlight-seo',
    hooks: {
      'config:setup'({ addRouteMiddleware }) {
        addRouteMiddleware({
          entrypoint: new URL('../route-data/seo.mjs', import.meta.url).pathname,
          order: 'post',
        });
      },
    },
  };
}

export function isArchivedVersionUrl(url) {
  return ARCHIVED_VERSION_PATH.test(new URL(url).pathname);
}
