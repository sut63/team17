import { PluginEnvironment } from '../types';
export default function createPlugin({ logger, database, }: PluginEnvironment): Promise<import("express").Router>;
