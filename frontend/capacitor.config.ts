import type { CapacitorConfig } from '@capacitor/cli';

const config: CapacitorConfig = {
  appId: 'com.microstars.app',
  appName: 'Consensus',
  webDir: 'dist',
  server: {
    androidScheme: 'https'
  }
};

export default config;
