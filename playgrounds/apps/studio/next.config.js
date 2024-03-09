/** @type {import('next').NextConfig} */
module.exports = {
  transpilePackages: ["@missingstudio/ui"],
  eslint: {
    ignoreDuringBuilds: true,
  },
};
