const { build } = require("esbuild");

const b = () =>
  build({
    bundle: true,
    entryPoints: [
      "node_modules/preline/preline.js",
      "node_modules/htmx.org/dist/htmx.js",
    ],
    platform: "node",
    outdir: "public/bundle",
    loader: { ".node": "file" },
    minify: true,
  });

Promise.all([b()]);
