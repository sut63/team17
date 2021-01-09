import type { PluginEnvironment } from '../types';
export default function createPlugin({ logger }: PluginEnvironment): Promise<import("express").Router>;
