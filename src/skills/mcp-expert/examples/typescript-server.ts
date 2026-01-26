/**
 * Full-featured MCP server example with tools and resources.
 * Run with: npx tsx server.ts
 */
import { McpServer } from '@modelcontextprotocol/sdk/server/mcp.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import * as z from 'zod';
import * as fs from 'fs/promises';

const server = new McpServer({
    name: 'my-server',
    version: '1.0.0'
});


// ============================================================
// TOOLS
// ============================================================

server.registerTool(
    'read-file',
    {
        title: 'Read File',
        description: 'Read contents of a file at the given absolute path',
        inputSchema: {
            path: z.string().describe('Absolute path to the file')
        },
        outputSchema: {
            content: z.string(),
            size: z.number()
        }
    },
    async ({ path: filePath }) => {
        try {
            const content = await fs.readFile(filePath, 'utf-8');
            const output = { content, size: content.length };
            return {
                content: [{ type: 'text', text: JSON.stringify(output, null, 2) }],
                structuredContent: output
            };
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : String(error);
            return {
                content: [{ type: 'text', text: `Error: ${errorMessage}` }],
                isError: true
            };
        }
    }
);

server.registerTool(
    'list-directory',
    {
        title: 'List Directory',
        description: 'List contents of a directory',
        inputSchema: {
            path: z.string().describe('Absolute path to the directory')
        },
        outputSchema: {
            entries: z.array(z.object({
                name: z.string(),
                isDirectory: z.boolean()
            }))
        }
    },
    async ({ path: dirPath }) => {
        try {
            const entries = await fs.readdir(dirPath, { withFileTypes: true });
            const output = {
                entries: entries.map(e => ({
                    name: e.name,
                    isDirectory: e.isDirectory()
                }))
            };
            return {
                content: [{ type: 'text', text: JSON.stringify(output, null, 2) }],
                structuredContent: output
            };
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : String(error);
            return {
                content: [{ type: 'text', text: `Error: ${errorMessage}` }],
                isError: true
            };
        }
    }
);


// ============================================================
// RESOURCES
// ============================================================

server.registerResource(
    'system-info',
    {
        uri: 'system://info',
        name: 'System Information',
        description: 'Current system information',
        mimeType: 'application/json'
    },
    async () => ({
        contents: [{
            uri: 'system://info',
            mimeType: 'application/json',
            text: JSON.stringify({
                platform: process.platform,
                nodeVersion: process.version,
                cwd: process.cwd(),
                env: process.env.NODE_ENV || 'development'
            }, null, 2)
        }]
    })
);


// ============================================================
// MAIN
// ============================================================

const transport = new StdioServerTransport();
console.error('Starting MCP server...');
await server.connect(transport);
