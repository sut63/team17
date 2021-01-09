import { PluginEnvironment } from '../types';
export default function createPlugin({ logger, database, config, }: PluginEnvironment): Promise<import("express").Router>;
